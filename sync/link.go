package sync

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"olympos.io/encoding/edn"
)

func NewConnection() (conn net.Conn) {
	// open socket
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalf("cannot establish connection\nDo you have Carabiner running?\n %v", err)
		return nil
	}
	return
}

func Parse(message *string, st *Status) error {
	err := edn.Unmarshal([]byte(*message), st)
	return err
}

func Ping(conn net.Conn, message string) (*string, error) {
	fmt.Fprintf(conn, "%s", message)                        // send message
	response, err := bufio.NewReader(conn).ReadString('\n') // listen response
	if err != nil {
		return nil, err
	}
	response = response[7:]
	return &response, nil
}
