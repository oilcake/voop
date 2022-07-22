package link

import "voop/engine"

// Provider is just any Link interface to get basic transport info
type Provider interface {
	Sync()
	Provide() *engine.Status
}

type Link struct {
	p         Provider
	tempoRary float64 // placeholder for current tempo value to check if it changes
}

func NewLink(p Provider) *Link {
	return &Link{
		p:         p,
		tempoRary: p.Provide().Bpm,
	}
}

// implementation of Transport
func (l *Link) Report() *engine.Status {
	l.Sync()
	return l.p.Provide()
}

// Update Provider's link information
func (l *Link) Sync() { l.p.Sync() }

// check if Tempo has changed
func (l *Link) BpmIsUpdated() bool {
	if l.tempoRary != l.p.Provide().Bpm {
		l.tempoRary = l.p.Provide().Bpm
		return true
	}
	return false
}
