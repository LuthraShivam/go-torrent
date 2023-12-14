package torrent

import (
	"crypto/rand"
	"fmt"
	"log"

	"github.com/LuthraShivam/go-torrent/pkg/bencode"
	"github.com/LuthraShivam/go-torrent/pkg/entities"
)

// ////////// Function Definitions
func ParseTorrentFiles(files []string) ([]bencode.DecodedTorrentData, error) {
	decodedInterfaces := make([]bencode.DecodedTorrentData, len(files))
	for i, torrentFile := range files {
		fmt.Printf("Parsing the following torrent file: %s\n", torrentFile)
		decodedInterface, err := bencode.Decode(torrentFile)
		if err != nil {
			log.Fatalln(err.Error())
			return nil, err
		}
		decodedInterfaces[i] = decodedInterface
	}
	return decodedInterfaces, nil
}

func BuildTorrent(decodedInterface bencode.DecodedTorrentData) (entities.Torrent, error) {

	var peerID [20]byte
	_, err := rand.Read(peerID[:])
	if err != nil {
		return entities.Torrent{}, err
	}

	var torrent entities.Torrent
	torrent.ID = peerID

	// determining if we are working with a single file torrent or a multi-file torrent
	if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentSingleFile); ok {
		fmt.Println("Single file torrent file encountered")
		hash, _ := decodedData.Info.InfoHash()
		fmt.Println(hash)

	} else if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentMultiFile); ok {
		fmt.Println("Multi file torrent file encountered")
		hash, _ := decodedData.Info.InfoHash()
		fmt.Println(hash)
	} else {
		fmt.Println("Foo")
	}
	return entities.Torrent{}, nil
}

func RequestPeersFromTracker(decodedInterface bencode.DecodedTorrentData) {
	// determining if we are working with a single file torrent or a multi-file torrent
	if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentSingleFile); ok {
		fmt.Println("Single file torrent file encountered")
		hash, _ := decodedData.Info.InfoHash()
		fmt.Println(hash)

	} else if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentMultiFile); ok {
		fmt.Println("Multi file torrent file encountered")
		hash, _ := decodedData.Info.InfoHash()
		fmt.Println(hash)
	} else {
		fmt.Println("Foo")
	}
}
