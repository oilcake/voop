package sync

import (
	"log"
)

const (
	BeatQuantity = 4 // THIS IS A STUB!!!!
	Divisor      = 4 // THIS IS A STUB!!!!
)

type Transport struct {
	Status        <-chan Status
	TimeSignature *TimeSignature
}

type Status struct {
	Peers int
	Bpm   float32
	Beat  float64
	D     bool // "tempo has been changed" flag
}

type TimeSignature struct {
	BeatQuantity uint8
	Divisor      uint8
}

type Engine interface {
	Dock() chan Status
}

func NewTransport(e Engine) (*Transport, error) {
	return &Transport{
		Status:        e.Dock(),
		TimeSignature: &TimeSignature{BeatQuantity, Divisor},
	}, nil
}

func (t *Transport) OneBeatDurationInMs() (duration float64) {
	duration = 60.0 / float64((<-t.Status).Bpm)
	log.SetFlags(log.Lshortfile)
	log.Printf("one beat is %v seconds\n", duration)
	return
}
