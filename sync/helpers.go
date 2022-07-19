package sync

import (
	"log"
	"net"
)

func NewConnection() *net.Conn {
	// open socket
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalf("cannot establish connection\nDo you have Carabiner running?\n %v", err)
		return nil
	}
	return &conn
}
