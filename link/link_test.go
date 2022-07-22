package link

import (
	"testing"
	"voop/engine"

	"github.com/stretchr/testify/assert"
)

type mockedProvider struct {
	St   *engine.Status
	stub *engine.Status
}

func (m *mockedProvider) Provide() *engine.Status      { return m.St }
func (m *mockedProvider) Sync()                        { m.St = m.stub }
func (m *mockedProvider) fillStub(stub *engine.Status) { m.stub = stub }

func TestCheckBpmIsUpdated(t *testing.T) {
	mp := &mockedProvider{
		St: &engine.Status{Bpm: 120.000000, Beat: 110716.064800},
	}
	link := NewLink(mp)
	// make sure it returns false on a fresh instance of link object
	assert.False(t, link.BpmIsUpdated(), "Bpm Is static")
	// Now update Bpm and check again
	mp.fillStub(
		&engine.Status{Bpm: 128.000000, Beat: 110716.064800},
	)
	link.Sync()
	assert.True(t, link.BpmIsUpdated(), "Bpm was updated, true expected")
}
