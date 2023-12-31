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
	torrentStruct, err := torrent.BuildTorrent(decodedInterfaces[0])
	if err != nil {
		fmt.Println("Unable to create torrent file")
		return
	}
	torrentStruct.RequestPeersFromTracker()
	fmt.Println(torrentStruct.Peers)

}
