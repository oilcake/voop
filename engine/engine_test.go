package engine

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

type mockedTransport struct {
	St   *Status
	stub *Status
}

func (m *mockedTransport) Report() *Status       { return m.St }
func (m *mockedTransport) BpmIsUpdated() bool    { return true }
func (m *mockedTransport) Sync()                 { m.St = m.stub }
func (m *mockedTransport) fillStub(stub *Status) { m.stub = stub }

func TestDurationOfOneBeatInMs(t *testing.T) {
	mT := &mockedTransport{
		St: &Status{Bpm: 120.000000, Beat: 110716.064800},
	}
	e := NewEngine(mT)

	var want float64 = 500.0
	got := e.DurationOfOneBeatInMs()
	assert.Equal(t, want, got, "Should be equal")
}

func TestForceBpmUpdate(t *testing.T) {
	mT := &mockedTransport{
		St: &Status{Bpm: 120.000000, Beat: 110716.064800},
	}
	e := NewEngine(mT)
	go e.ForceBpmUpdate()
	got := <-e.BpmUpdate
	assert.Equal(t, e.Status.Bpm, got, "should be equal")
}
