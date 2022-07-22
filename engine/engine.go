package engine

import (
	"fmt"
	"math"
	"voop/clock"
)

const (
	BeatsInBar = 7 // THIS IS A STUB!!!!
	BeatSize   = 8 // THIS IS A STUB!!!!
)

type Boom struct {
	Bar  chan struct{}
	Beat chan struct{}
}

type TimeSignature struct {
	BeatsInBar float64
	BeatSize   float64
}

type Status struct {
	Bpm  float64
	Beat float64
}

type Transport interface {
	Report() *Status
	BpmIsUpdated() bool
	Sync()
}

type Engine struct {
	Transport     Transport
	Status        *Status
	TimeSignature TimeSignature
	BpmUpdate     chan float64
	VClock        clock.VClock
	Boom          *Boom
}

func NewEngine(t Transport) *Engine {
	return &Engine{
		Transport:     t,
		Status:        t.Report(),
		TimeSignature: TimeSignature{BeatsInBar, BeatSize},
		BpmUpdate:     make(chan float64),
		VClock:        *clock.NewVClock(),
		Boom: &Boom{
			Bar:  make(chan struct{}),
			Beat: make(chan struct{}),
		},
	}
}

func (e *Engine) Start() {
	e.VClock.Start()
	var (
		mP          float64
		phaseBefore float64
		phaseAfter  float64
	)
	mP = 1.0 * 7

	for {
		e.Transport.Sync()
		phaseAfter = math.Mod(float64(e.Status.Beat), mP) / mP
		if phaseAfter < phaseBefore {
			e.Boom.Bar <- struct{}{}
			// fmt.Println()
		}
		phaseBefore = phaseAfter

		fmt.Printf("\rCurrent bpm is %.9f and beat is %.9f and phase is %.9f", e.Status.Bpm, e.Status.Beat, phaseAfter)
	}
}

func (e *Engine) DurationOfOneBeatInMs() (duration float64) {
	duration = 60.0 * 1000.0 / e.Status.Bpm
	return
}

func (e *Engine) ForceBpmUpdate() {
	if e.Transport.BpmIsUpdated() {
		e.BpmUpdate <- e.Transport.Report().Bpm
	}
}
