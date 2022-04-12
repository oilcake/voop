package sync

import "log"

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"

	Measure  = 4 // THIS IS A STAB!!!!
	Division = 4 // THIS IS A STAB!!!!
)

type Transport struct {
	St            chan Status
	TimeSignature *TimeSignature
}

type Status struct {
	Peers int
	Bpm   float32
	Start int64
	Beat  float64
	D     bool // tempo is change flag
}
type TimeSignature struct {
	Measure  uint8
	Division uint8
}

func NewTransport() (*Transport, error) {
	st := make(chan Status)
	return &Transport{
		St:            st,
		TimeSignature: &TimeSignature{4, 4},
	}, nil
}

func (t *Transport) BeatDur() (duration float64) {
	st := <-t.St
	oneBeatDuration := 60.0 / float64(st.Bpm)
	log.Printf("one beat is %v milliseconds\n", oneBeatDuration)
	return oneBeatDuration
}
