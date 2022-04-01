package link

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"time"

	"olympos.io/encoding/edn"
)

func Watch(conn net.Conn, st *Status, d *bool) {
	for {
		oldTempo := st.Bpm
		response, err := Ping(conn, "status")
		if err != nil {
			log.Fatal("no response from Carabiner", err)
		}
		err = Parse(response, st)
		if err != nil {
			log.Fatal("Parsing error", err)
		}
		newTempo := st.Bpm
		if oldTempo != newTempo {
			*d = true
		}
		time.Sleep(time.Millisecond * 80)
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
