package clip

import (
	"fmt"
	"log"
	"math"
	"math/rand"
	"time"
)

func (m *Media) LoopPhase() float64 {
	st := <-m.transport.Status

	if st.D {
		log.SetFlags(log.Lshortfile)
		log.Println("Tempo is now", st.Bpm)
		m.Grooverize()
	}
	phase := math.Mod(st.Beat, m.LoopLen) / m.LoopLen
	return phase
}

func (m *Media) PalindromemordnilaP() {
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
	m.Grooverize()
}

// this function calculates positive remainder from division to 1
func (m *Media) Wrap(x, y float64) float64 {
	return math.Mod(math.Mod(x, y)+y, y)
}

func (m *Media) calcFrame() (frame float64) {
	m.phase = m.LoopPhase()
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
		m.shiftedPhase = m.Wrap(m.shiftedPhase+m.timepoint+shift, 1)
		m.shiftedPhase = math.Abs(m.dirPld - math.Abs(m.shiftedPhase*2.0-1))
	case false:
		m.shiftedPhase = m.Wrap(m.shiftedPhase+m.timepoint, 1)
	}
	frame = m.Framecount * m.shiftedPhase
	fmt.Printf("\rCurrent frame %06d, phase %.2f, offset %.2f, shiftedPhase %.2f, dirPld %.2f",
		int(frame), m.phase, m.offset, m.shiftedPhase, m.dirPld)
	// looks like I really have to do it twice, otherwise we have jumping frame
	m.phase = m.LoopPhase()
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

func (m *Media) DefaultRate() {
	m.Multiple = 1
	m.MultRate(1)
}

func (m *Media) MultRate(rate float64) {
	t := m.shiftedPhase
	m.Multiple *= rate
	m.Grooverize()
	m.phase = m.LoopPhase()
	m.offset = m.phase
	m.timepoint = t
}
