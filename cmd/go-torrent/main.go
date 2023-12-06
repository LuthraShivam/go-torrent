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
		fmt.Println(torrentFile)

		decodedInterface, err := bencode.Decode(torrentFile)
		if err != nil {
			log.Fatalln(err.Error())
			return
		}
		if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentSingleFile); ok {
			fmt.Println(decodedData.Announce)
		} else {
			fmt.Println("type assertion failed")
		}
		//fmt.Println(decodedData)
		// fmt.Println(torrentData)
		// fmt.Println(torrentData.Files[0].Path)
	}
}
