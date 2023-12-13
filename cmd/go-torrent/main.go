package main

import (
	"fmt"
	"log"
	"os"

	"github.com/LuthraShivam/go-torrent/pkg/torrent"
)

func main() {
	/*
		There are some torrent files that have more content that just a single line.
		For now, we are only parsing the fist line of each torrent file, as that "may"
		contain all the info we need.
		Execution structure: ./<executable> <torrent1> <torrent2> ...
	*/
	if len(os.Args) < 2 {
		log.Fatal("Torrent Files not passed. Exiting\n")
	}
	decodedInterfaces, _ := torrent.ParseTorrentFiles(os.Args[1:])
	for _, dI := range decodedInterfaces {
		fmt.Println(dI)
	}
	// for _, torrentFile := range os.Args[1:] {
	// 	fmt.Println(torrentFile)

	// 	decodedInterface, err := bencode.Decode(torrentFile)
	// 	if err != nil {
	// 		log.Fatalln(err.Error())
	// 		return
	// 	}
	// 	if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentSingleFile); ok {
	// 		fmt.Println("Single file torrent file encountered")
	// 		hash, _ := decodedData.Info.InfoHash()
	// 		fmt.Println(hash)

	// 	} else {
	// 		if decodedData, ok := decodedInterface.(*bencode.BencodeTorrentMultiFile); ok {
	// 			fmt.Println("Multi file torrent file encountered")
	// 			hash, _ := decodedData.Info.InfoHash()
	// 			fmt.Println(hash)
	// 		}
	// 	}

	// }
}
