package clip

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"sync"
	"time"
)

var (
	r  float64
	mu sync.Mutex
)

func (m *Media) calcFrame() (frame float64) {
	mu.Lock()
	select {
	case r = <-m.RateX:
		m.multRate(r)
	default:
		m.phase = m.LoopPhase()
	}
	m.antiphase = -m.phase
	switch m.forward {
	case true:
		m.shiftedPhase = m.phase - m.offset
		break
	case false:
		m.shiftedPhase = m.antiphase + m.offset
		break
	}
	// check palindromicity
	switch m.palindrome {
	case true:
		shift = (m.offset - m.timepoint) * 0.5
		m.shiftedPhase = Wrap(m.shiftedPhase+m.timepoint, 1)
		m.shiftedPhase = math.Abs(m.dirPld - math.Abs(m.shiftedPhase*2.0-1))
	case false:
		m.shiftedPhase = Wrap(m.shiftedPhase+m.timepoint, 1)
	}
	frame = m.Framecount * m.shiftedPhase
	fmt.Printf("\rCurrent frame %06d, phase %.2f, offset %.2f, shiftedPhase %.2f, dirPld %.2f",
		int(frame), m.phase, m.offset, m.shiftedPhase, m.dirPld)
	mu.Unlock()
	return
}

func (m *Media) GetLoopPhase(loopLen float64) (phase float64) {
	st := <-m.transport.Status

	if st.D {
		log.SetFlags(log.Lshortfile)
		log.Println("Tempo is now", st.Bpm)
		m.Grooverize()
	}
	phase = math.Mod(st.Beat, loopLen) / loopLen
	return
}

func (m *Media) LoopPhase() (phase float64) {
	phase = m.GetLoopPhase(m.LoopLen)
	return
}

func (m *Media) PalindromemordnilaP() {
	m.palindrome = !m.palindrome
	switch m.palindrome {
	case true:
		m.RateX <- 2.0
	case false:
		m.RateX <- 0.5
	}
}

func (m *Media) Swap() {
	m.forward = !m.forward
	if !m.hardSync {
		m.timepoint = m.shiftedPhase
		m.offset = m.phase
	}
}

func (m *Media) Zero() {
	m.offset = m.phase
	m.timepoint = 0
}

func (m *Media) ReSync() {
	m.offset = 0
	m.timepoint = 0
}
func (m *Media) Jump() {
	rand.Seed(time.Now().UnixNano())
	m.offset = rand.Float64()
}

func (m *Media) DefaultRate() {
	m.multiple = 1
	m.RateX <- 1
}

func (m *Media) multRate(rate float64) {
	switch m.hardSync {
	case false:
		t := m.shiftedPhase
		m.multiple *= rate
		m.Grooverize()
		m.phase = m.LoopPhase()
		m.offset = m.phase
		m.timepoint = t
		break
	case true:
		m.multiple *= rate
		m.Grooverize()
		m.ReSync()
	}
}

func (m *Media) HardSyncToggle() {
	m.hardSync = !m.hardSync
	if m.hardSync {
		m.ReSync()
	}
}

// this function calculates positive remainder from division
func Wrap(x, y float64) float64 {
	return math.Mod(math.Mod(x, y)+y, y)
}
