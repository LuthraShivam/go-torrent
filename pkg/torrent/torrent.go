package torrent

import (
	"fmt"
	"log"

	"github.com/LuthraShivam/go-torrent/pkg/bencode"
)

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
