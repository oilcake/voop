package sync

import (
	"log"
)

const (
	Measure  = 4 // THIS IS A STUB!!!!
	Division = 4 // THIS IS A STUB!!!!
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
	Measure  uint8
	Division uint8
}

type Engine interface {
	Dock() chan Status
}

func NewTransport(e Engine) (*Transport, error) {
	return &Transport{
		Status:        e.Dock(),
		TimeSignature: &TimeSignature{Measure, Division},
	}, nil
}

func (t *Transport) OneBeatDurationInMs() (duration float64) {
	duration = 60.0 / float64((<-t.Status).Bpm)
	log.SetFlags(log.Lshortfile)
	log.Printf("one beat is %v seconds\n", duration)
	return
}
