package torrent

import (
	"crypto/rand"
	"errors"
	"fmt"
	"log"

	"github.com/LuthraShivam/go-torrent/pkg/bencode"
	"github.com/LuthraShivam/go-torrent/pkg/entities"
)

type File struct {
	Path   []string
	Length int
}

type Torrent struct {
	Announce    string
	InfoHash    entities.SHAHash
	PieceLength int
	Files       []File // represents the actual file
	Name        string
	PieceHashes [][20]byte
	Peers       []entities.Peer
	ID          [20]byte
}

// ////////// Function Definitions

func SplitPieceHashes(pieces string) ([][20]byte, error) {
	hashLength := 20
	buf := []byte(pieces)
	if len(buf)%hashLength != 0 {
		err := fmt.Errorf("Received malformed pieces of length: %d", len(buf))
		return nil, err
	}
	numberHashes := len(buf) / hashLength
	hashes := make([][20]byte, numberHashes)
	for i := 0; i < numberHashes; i++ {
		copy(hashes[i][:], buf[i*hashLength:(i+1)*hashLength])
	}
	return hashes, nil
}

func ParseTorrentFiles(files []string) ([]bencode.DecodedTorrentData, error) {
	decodedInterfaces := make([]bencode.DecodedTorrentData, len(files))
	for i, torrentFile := range files {
		// fmt.Printf("Parsing the following torrent file: %s\n", torrentFile)
		decodedInterface, err := bencode.Decode(torrentFile)
		if err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}
		decodedInterfaces[i] = decodedInterface
	}
	return decodedInterfaces, nil
}

func BuildTorrent(decodedInterface bencode.DecodedTorrentData) (Torrent, error) {

	var peerID [20]byte
	// we're using a random 20 byte value as the peer_id
	_, err := rand.Read(peerID[:])
	if err != nil {
		return Torrent{}, err
	}

	// populate torrent ID
	var torrent Torrent
	torrent.ID = peerID

	// determining if we are working with a single file torrent or a multi-file torrent
	if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentSingleFile); ok {
		// fmt.Println("BuildTorrent | Single file torrent file encountered")
		hash, err := decodedData.Info.InfoHash()
		if err != nil {
			error := errors.New("BuildTorrent | single file torrent | error while computing infohash")
			return Torrent{}, error
		}
		torrent.Announce = decodedData.Announce
		torrent.InfoHash = hash
		torrent.PieceLength = decodedData.Info.PieceLength
		torrent.Name = decodedData.Info.Name
		pieceHashes, err := SplitPieceHashes(decodedData.Info.Pieces)
		if err != nil {
			return Torrent{}, err
		}
		torrent.PieceHashes = pieceHashes
		// populate contents of one file
		file := File{}
		file.Length = decodedData.Info.Length
		filePath := []string{}
		filePath = append(filePath, decodedData.Info.Name)
		file.Path = filePath

		// finally add everything into torrent files
		files := []File{}
		files = append(files, file)
		torrent.Files = files
	} else {
		// multifile torrent
		if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentMultiFile); ok {
			// fmt.Println("BuildTorrent | Multi file torrent file encountered")
			hash, err := decodedData.Info.InfoHash()
			if err != nil {
				error := errors.New("BuildTorrent | multi file torrent | error while computing infohash")
				return Torrent{}, error
			}
			torrent.Announce = decodedData.Announce
			torrent.InfoHash = hash
			torrent.PieceLength = decodedData.Info.PieceLength
			torrent.Name = decodedData.Info.Name
			pieceHashes, err := SplitPieceHashes(decodedData.Info.Pieces)
			if err != nil {
				return Torrent{}, err
			}
			torrent.PieceHashes = pieceHashes
			files := []File{}
			// translate bencode files to entities file
			for _, file := range decodedData.Info.Files {
				f := File{}
				f.Length = file.Length
				f.Path = file.Path
				files = append(files, f)
			}
			torrent.Files = files
		}
		// peers will be populated by a different function
	}

	return torrent, nil
}
