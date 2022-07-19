package sync

const (
	Measure  = 4 // THIS IS A STUB!!!!
	Division = 4 // THIS IS A STUB!!!!
)

// Linker is just any Link interface to get basic transport info
type Linker interface {
	Link(chan Status, TempoWatcher)
}

type TempoWatcher chan struct{} // "tempo have been changed" signal

type Transport struct {
	Status        chan Status
	TimeSignature *TimeSignature
	clock         *Clock
	stop          chan struct{}
	TempoWatcher  TempoWatcher
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

func NewTransport(c *Clock) (*Transport, error) {
	st := make(chan Status)
	return &Transport{
		Status:        st,
		TimeSignature: &TimeSignature{Measure, Division},
		clock:         c,
	}, nil
}

// Note that this is a blocking operation
func (t *Transport) Sync(l Linker) {
	l.Link(t.Status, t.TempoWatcher)
}

func (t *Transport) BeatDur() (duration float64) {
	oneBeatDuration := 60.0 / float64((<-t.Status).Bpm)
	return oneBeatDuration
}

func (t *Transport) Stop() {
	t.stop <- struct{}{}
}
