package main

import (
	"fmt"
	"log"
	"os"

	bencode "github.com/LuthraShivam/go-torrent/pkg/bencode"
)

func main() {
	/*
		There are some torrent files that have more content that just a single line.
		For now, we are only parsing the fist line of each torrent file, as that "may"
		contain all the info we need.
		Execution structure: ./<executable> <torrent1> <torrent2> ...
	*/
	if len(os.Args) < 2 {
		log.Fatal("Torrent File not passed. Exiting\n")
	}

	for _, torrentFile := range os.Args[1:] {
		torrentData := bencode.Decode(torrentFile)
		fmt.Println(torrentData.Files)
	}
}
