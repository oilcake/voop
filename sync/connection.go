package sync

import (
	"bufio"
	"fmt"
	"log"
	"net"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"
)

type connection struct {
	tcp *net.Conn
}

func NewConnection() *connection {
	// open socket
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalf("cannot establish connection\nDo you have Carabiner running?\n %v", err)
		return nil
	}
	return &connection{
		tcp: &conn,
	}
}

// implementation of connector
// Get info from Carabiner
func (c connection) getStatus() (response string) {
	fmt.Fprintf(*c.tcp, "%s", "status")                       // send message
	response, err := bufio.NewReader(*c.tcp).ReadString('\n') // listen response
	if err != nil {
		log.Fatal("connection lost", err)
	}
	return response[7:]
}
