package sync

import "log"

const (
	protocol = "tcp"
	address  = "127.0.0.1:17000"

	Measure  = 4 // THIS IS A STUB!!!!
	Division = 4 // THIS IS A STUB!!!!
)

type Transport struct {
	Status        chan Status
	TimeSignature *TimeSignature
	clock         *Clock
	stop          chan struct{}
}

type Status struct {
	Peers int
	Bpm   float32
	Start int64
	Beat  float64
	D     bool // "tempo have been changed" flag
}

type TimeSignature struct {
	Measure  uint8
	Division uint8
}

func NewTransport(c *Clock) (*Transport, error) {
	st := make(chan Status)
	return &Transport{
		Status:        st,
		TimeSignature: &TimeSignature{Measure, Division},
		clock:         c,
	}, nil
}

func (t *Transport) Start() {
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
			select {
			case <-t.stop:
				log.Println("stop Link")
				close(t.Status)
				t.clock.Stop()
				return
			default:
				// pass
			}

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
			t.Status <- watch
			watch.D = false
		}
	}()

}

func (t *Transport) Stop() {
	t.stop <- struct{}{}
}
func (t *Transport) BeatDur() (duration float64) {
	oneBeatDuration := 60.0 / float64((<-t.Status).Bpm)
	return oneBeatDuration
}
