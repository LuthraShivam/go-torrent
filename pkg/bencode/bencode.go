package bencode

import (
	"bytes"
	"crypto/sha1"
	"fmt"
	"os"

	entities "github.com/LuthraShivam/go-torrent/pkg/entities"
	bencode "github.com/jackpal/bencode-go"
)

type DecodedTorrentData interface {
	Unmarshal(string) error
}

// ////////// Decoded torrent data structures
type bencodeFile struct {
	Path   []string
	Length int
}
type bencodeInfoMultiFile struct {
	Pieces      string        `bencode:"pieces"`
	PieceLength int           `bencode:"piece length"`
	Name        string        `bencode:"name"`
	Files       []bencodeFile `bencode:"files"`
}

type bencodeInfoSingleFile struct {
	Pieces      string `bencode:"pieces"`
	PieceLength int    `bencode:"piece length"`
	Length      int    `bencode:"length"`
	Name        string `bencode:"name"`
}

type BencodeTorrentSingleFile struct {
	Announce     string                `bencode:"announce"`
	CreationDate int                   `bencode:"creation date"`
	Info         bencodeInfoSingleFile `bencode:"info"`
}
type BencodeTorrentMultiFile struct {
	Announce     string               `bencode:"announce"`
	CreationDate int                  `bencode:"creation date"`
	Info         bencodeInfoMultiFile `bencode:"info"`
}

// ////////// Interface functions for above structures

func (bto *BencodeTorrentSingleFile) Unmarshal(torrentPath string) error {
	file, err := os.Open(torrentPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

func (bto *BencodeTorrentMultiFile) Unmarshal(torrentPath string) error {
	file, err := os.Open(torrentPath)
	if err != nil {
		return err
	}
	defer file.Close()
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// ////////// Bencode Info hash functions

// To generate Hash, you need the info json segment of the torrent file, and return the hash result.
// The length of the hash will be 20 bytes

// for single file torrent files
func (bto *bencodeInfoSingleFile) InfoHash() (entities.SHAHash, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *bto)
	if err != nil {
		return entities.SHAHash{}, nil
	}
	hash := sha1.Sum(buf.Bytes())
	return hash, nil
}

// for multi-file torrent files
func (bto *bencodeInfoMultiFile) InfoHash() (entities.SHAHash, error) {
	var buf bytes.Buffer
	err := bencode.Marshal(&buf, *bto)
	if err != nil {
		return entities.SHAHash{}, nil
	}
	hash := sha1.Sum(buf.Bytes())
	return hash, nil
}

// ////////// Package functions

func Decode(torrentPath string) (DecodedTorrentData, error) {
	decodedData := BencodeTorrentSingleFile{}
	err := decodedData.Unmarshal(torrentPath)
	if err != nil {
		fmt.Println(err.Error())
		return nil, err
	}

	// if you encounter multi file torrent file
	if decodedData.Info.Length == 0 {
		fmt.Println("Encountered a multi file torrent. Parsing appropriately")
		decodedData := BencodeTorrentMultiFile{}
		err := decodedData.Unmarshal(torrentPath)
		if err != nil {
			fmt.Println(err.Error())
			return nil, err
		}
		return &decodedData, nil
	}
	return &decodedData, nil // &decodeData to match the interface
}
