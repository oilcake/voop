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
		log.Fatal("cannot establish connection", err)
		return nil
	}
	return
}

func Link(conn net.Conn, st chan<- Status) {
	// init status
	watch := &Status{
		Peers: 0,
		Bpm:   120.0,
		Start: 0,
		Beat:  0.0,
		D:     true,
	}
	// watch what's in Link
	for {
		st <- *watch
		oldTempo := watch.Bpm
		response, err := Ping(conn, "status")
		if err != nil {
			log.Fatal("no response from Carabiner", err)
		}
		err = Parse(response, watch)
		if err != nil {
			log.Fatal("Parsing error", err)
		}
		newTempo := watch.Bpm
		if oldTempo != newTempo {
			watch.D = true
		}
		watch.D = false
	}
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
