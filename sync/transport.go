package sync

import "log"

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"

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
	Start int64
	Beat  float64
	D     bool // "tempo has been changed" flag
}

type TimeSignature struct {
	Measure  uint8
	Division uint8
}

func NewTransport() (*Transport, error) {
	st := make(chan Status)
	NewLink(st)
	return &Transport{
		Status:        st,
		TimeSignature: &TimeSignature{Measure, Division},
	}, nil
}

func (t *Transport) OneBeatDurationInMs() (duration float64) {
	duration = 60.0 / float64((<-t.Status).Bpm)
	log.Printf("one beat is %v seconds\n", duration)
	return
}
