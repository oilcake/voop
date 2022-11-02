package sync

import (
	"log"
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

func NewTransport(e Engine, t *TimeSignature) *Transport {
	return &Transport{
		Status:        e.Dock(),
		TimeSignature: t,
	}
}

func (t *Transport) OneBeatDurationInMs() (duration float64) {
	duration = 60.0 / float64((<-t.Status).Bpm)
	log.SetFlags(log.Lshortfile)
	log.Printf("one beat is %v seconds\n", duration)
	return
}
