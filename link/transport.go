package link

import (
	"fmt"
	"log"
	"net"
)

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"

	Measure  = 4 // THIS IS A STAB!!!!
	Division = 4 // THIS IS A STAB!!!!
)

type Transport struct {
	St            *Status
	TimeSignature *TimeSignature
	D             *bool // tempo change watcher
}

type Status struct {
	Peers int
	Bpm   float32
	Start int64
	Beat  float64
}
type TimeSignature struct {
	Measure  uint8
	Division uint8
}

func NewTransport() (*Transport, error) {
	var st Status
	// open socket
	conn, err := net.Dial(protocol, address)
	if err != nil {
		log.Fatal("cannot establish connection", err)
	}

	tWatcher := true
	go Watch(conn, &st, &tWatcher)

	return &Transport{
		St:            &st,
		TimeSignature: &TimeSignature{4, 4},
		D:             &tWatcher,
	}, nil
}

func (t *Transport) BeatDur() (duration float64) {
	oneBeatDuration := 60.0 / float64(t.St.Bpm)
	fmt.Printf("one beat is %v milliseconds\n", oneBeatDuration)
	return oneBeatDuration
}
