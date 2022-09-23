package clip

import (
	"fmt"
	"math"
	"math/rand"
	"time"
	"voop/sync"
)

func (m *Media) PalindromemordnilaP(t *sync.Transport) {
	m.palindrome = !m.palindrome
	switch m.forward {
	case true:
		m.dirPld = 1
	case false:
		m.dirPld = 0
	}
	switch m.palindrome {
	case true:
		m.Multiple = m.Multiple * 2.0
	case false:
		m.Multiple = m.Multiple * 0.5
	}
	m.Grooverize(t)
}

// this function calculates positive remainder from division to 1
func (m *Media) Wrap(x, y float64) float64 {
	return math.Mod(math.Mod(x, y)+y, y)
}

func (m *Media) calcFrame() (frame float64) {
	m.antiphase = -m.phase
	switch m.forward {
	case true:
		m.shiftedPhase = m.phase - m.offset
		break
	case false:
		m.antiphase += m.offset
		m.shiftedPhase = m.antiphase
		break
	}
	// check palindromicity
	switch m.palindrome {
	case true:
		shift = (m.offset - m.timepoint) * 0.5
		m.shiftedPhase = m.Wrap(m.shiftedPhase+m.timepoint+shift, 1)
		m.shiftedPhase = math.Abs(m.dirPld - math.Abs(m.shiftedPhase*2.0-1))
	case false:
		m.shiftedPhase = m.Wrap(m.shiftedPhase+m.timepoint, 1)
	}
	frame = m.Framecount * m.shiftedPhase
	fmt.Printf("\rCurrent frame %06d, phase %.2f, offset %.2f, shiftedPhase %.2f, dirPld %.2f",
		int(frame), m.phase, m.offset, m.shiftedPhase, m.dirPld)
	return
}

func (m *Media) Swap() {
	m.forward = !m.forward
	m.timepoint = m.shiftedPhase
	m.offset = m.phase
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

func (m *Media) UpdateRate(rate float64, t *sync.Transport) {
	m.Multiple *= rate
	m.Grooverize(t)
}
