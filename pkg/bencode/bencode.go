package bencode

import (
	"log"

	gotorrentparser "github.com/j-muller/go-torrent-parser"
)

func Decode(torrentPath string) gotorrentparser.Torrent {
	torrent, err := gotorrentparser.ParseFromFile(torrentPath)
	if err != nil {
		log.Fatal(err)
	}
	return *torrent
	// fmt.Println(torrent)
}
