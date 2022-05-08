package sync

import (
	"bufio"
	"fmt"
	"log"
	"net"

	"olympos.io/encoding/edn"
)

type Link struct {
	message chan string
	st      *Status
}

func NewConnection() (conn net.Conn) {
	// open socket
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatalf("cannot establish connection\n %v", err)
		return nil
	}
	return
}

func NewLink(st chan Status) {
	var (
		response *string
		err      error
		watch    Status
		oldTempo float32
		newTempo float32
	)
	watch.Bpm = 0.0
	// start Ableton Link watcher
	conn := NewConnection()
	go func() {
		for {
			oldTempo = watch.Bpm
			response, err = Ping(conn, "status")
			if err != nil {
				log.Fatal("no response from Carabiner", err)
			}
			err = Parse(response, &watch)
			if err != nil {
				log.Fatal("Parsing error", err)
			}
			newTempo = watch.Bpm
			if oldTempo != newTempo {
				watch.D = true
				oldTempo = newTempo
			}
			st <- watch
			watch.D = false
		}
	}()

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
