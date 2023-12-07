package tracker

import "net"

type Peer struct {
	IP   net.IP
	Port uint16
}

// Static port that Client listens on for responses from Tracker
const ListenPort uint16 = 6885
