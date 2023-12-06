package bencode

import (
	"fmt"
	"os"

	bencode "github.com/jackpal/bencode-go"
)

type decodedTorrentData interface {
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
func (bto BencodeTorrentSingleFile) Unmarshal(torrentPath string) error {
	file, err := os.Open(torrentPath)
	if err != nil {
		return err
	}
	defer file.Close()
	bto = BencodeTorrentSingleFile{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}

	// fmt.Println("SingleFileTorrent Parsing: ///")
	// fmt.Println(bto.Announce)
	// fmt.Println(bto.CreationDate)
	// fmt.Println(bto.Info.Length)
	return nil
}

func (bto BencodeTorrentMultiFile) Unmarshal(torrentPath string) error {
	file, err := os.Open(torrentPath)
	if err != nil {
		return err
	}
	defer file.Close()
	bto = BencodeTorrentMultiFile{}
	err = bencode.Unmarshal(file, &bto)
	if err != nil {
		fmt.Println(err.Error())
		return err
	}
	return nil
}

// ////////// Package functions

/*
Decode function is responsible for
*/
func Decode(torrentPath string) (decodedTorrentData, error) {
	decodedData := BencodeTorrentSingleFile{}
	err := decodedData.Unmarshal(torrentPath)
	if err != nil {
		fmt.Println(err.Error())
		return BencodeTorrentSingleFile{}, err
	}

	// if you encounter multi file torrent file
	if decodedData.Info.Length == 0 {
		fmt.Println("Encountered a multi file torrent. Parsing appropriately")
		decodedData := BencodeTorrentMultiFile{}
		err := decodedData.Unmarshal(torrentPath)
		if err != nil {
			fmt.Println(err.Error())
			return BencodeTorrentMultiFile{}, nil
		}
		fmt.Println(decodedData.Info.Pieces)
		return &decodedData, nil
	}
	fmt.Println(decodedData.Info.Pieces)
	return &decodedData, nil
}
