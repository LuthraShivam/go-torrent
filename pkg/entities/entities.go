package entities

import "net"

type SHAHash [20]byte

type Peer struct {
	IP   net.IP
	Port uint16
}

// Static port that Client listens on for responses from Tracker
const ListenPort uint16 = 6885

// Structure that defines a torrent object - each torrent object represents a torrent being downloaded
// I will use one struct to define both single file and multi file torrents
// Single file torrents will contain just one entry under the file section
