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
type Torrent struct {
	Announce    string
	InfoHash    SHAHash
	PieceLength int
	Length      int
	Name        string
	Pieces      string
	Peers       []Peer
	ID          [20]byte
}
