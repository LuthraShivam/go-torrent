package torrent

import (
	"encoding/binary"
	"errors"
	"fmt"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"strings"
	"time"

	"github.com/LuthraShivam/go-torrent/pkg/entities"
	bencode "github.com/jackpal/bencode-go"
)

type bencodeTrackerResp struct {
	Interval int    `bencode:"interval"`
	Peers    string `bencode:"peers"`
}

type connRequestUDP struct {
	ProtocolID    uint64
	Action        uint32
	TransactionID uint32
}

type connResponseUDP struct {
	Action        uint32
	TransactionID uint32
	ConnectionID  uint64
}

type announceRequestUDP struct {
	// for IPv4 - https://www.bittorrent.org/beps/bep_0015.html
	ConnectionID  uint64
	Action        uint32
	TransactionID uint32
	InfoHash      entities.SHAHash
	PeerID        [20]byte
	Download      uint64
	Left          uint64
	Uploaded      uint64
	Event         uint32
	IPAddress     uint32
	Key           uint32
	NumWanted     uint32
	Port          uint16
}

// ////////// Function definitions
func (t *Torrent) RequestPeersFromTracker() {
	if strings.Contains(t.Announce, "http") {
		bytesLeft := 0
		for _, file := range t.Files {
			bytesLeft += file.Length
		}
		peers, err := requestPeersHTTP(*t, 0, 0, bytesLeft)
		if err != nil {
			fmt.Printf("RequestPeersFromTracker | %s\n", err.Error())
		}
		t.Peers = peers
	} else if strings.Contains(t.Announce, "udp") {
		fmt.Println("UDP trackers Not yet supported")
	} else {
		fmt.Println("Unsupported Protocol. Exiting")
	}
}

func UnmarshelPeers(peersBin []byte) ([]entities.Peer, error) {
	const peerSize = 6 // 4 for IP address, 2 for ports
	numPeers := len(peersBin) / peerSize
	if len(peersBin)%peerSize != 0 {
		err := fmt.Errorf("UnmarshelPeers | malformed peer information received")
		return nil, err
	}
	peers := make([]entities.Peer, numPeers)
	for i := 0; i < numPeers; i++ {
		offset := peerSize * i
		peers[i].IP = net.IP(peersBin[offset : offset+4])
		// https://stackoverflow.com/questions/38675266/go-convert-2-byte-array-into-a-uint16-value
		peers[i].Port = binary.BigEndian.Uint16(peersBin[offset+4 : offset+6])
	}
	return peers, nil
}

/*
Issues that can be faced while getting Peers:
 1. Tracker can not recognize the info hash provided - check why this is happening
 2. Unable to reach tracker
 3. Requested download is not authorized for use with this tracker - you are not meeting some certain criteria
    that the above mentioned tracker wants - occurs with private trackers, Tracker Rules Violations etc.
*/
func requestPeersHTTP(t Torrent, uploadedBytes int, downloadedBytes int, bytesLeft int) ([]entities.Peer, error) {
	base, err := url.Parse(t.Announce)
	if err != nil {
		return nil, err
	}
	params := url.Values{
		"info_hash":  []string{string(t.InfoHash[:])}, // byte array to string to string array ?
		"peer_id":    []string{string(t.ID[:])},
		"port":       []string{strconv.Itoa(int(entities.ListenPort))},
		"uploaded":   []string{strconv.Itoa(uploadedBytes)},
		"downloaded": []string{strconv.Itoa(downloadedBytes)},
		"left":       []string{strconv.Itoa(bytesLeft)},
		"compact":    []string{"1"},
	}
	base.RawQuery = params.Encode()
	// making http requests to tracker URL
	c := http.Client{Timeout: 10 * time.Second}
	resp, err := c.Get(base.String())
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	trackerResponse := bencodeTrackerResp{}
	err = bencode.Unmarshal(resp.Body, &trackerResponse)

	if err != nil {
		fmt.Printf("error val: %s\n", err.Error())
		return nil, errors.New(fmt.Sprintf("tracker | requestPeersHTTP | unable to unmarshal bencoded response, err: %s", err.Error()))
	}

	peers, err := UnmarshelPeers([]byte(trackerResponse.Peers))
	if err != nil {
		return nil, err
	}
	return peers, nil

}

func requestPeersUDP(t Torrent, uploadedBytes int, downloadedBytes int, bytesLeft int) {
	connRequest := connRequestUDP{
		Action: 0,
	}
	fmt.Println(connRequest)
}
