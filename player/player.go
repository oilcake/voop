package player

import (
	"fmt"
	"log"
	"math"
	"voop/clip"
	"voop/sync"
)

type Player struct {
	*sync.Clock
	*sync.Transport
	*Window
	*clip.Media
	HotKey chan int
}

func (p *Player) LoopPhase() float64 {
	st := <-p.Transport.Status

	if st.D {
		log.Println("Tempo is now", (<-p.Transport.Status).Bpm)
		p.Media.Grooverize(p.Transport)
	}
	phase := math.Mod(st.Beat, p.Media.LoopLen) / p.Media.LoopLen
	return phase
}

func (p *Player) PlayMedia() {
	// go p.WatchTempo()
	p.HotKey = make(chan int)
	var k int
	// play it in cycle forever
play:
	for range p.Clock.Trigger {
		// calculate a playing phase
		ph := p.LoopPhase()
		fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", (<-p.Transport.Status).Beat, ph)
		// to retrieve specific frame
		img := p.Media.Frame(ph)
		// and display it
		p.Window.IMShow(img)
		k = p.Window.WaitKey(19)
		switch {
		case k == 27:
			break play
		case k != 27 && k != -1:
			p.HotKey <- k
		}
	}
}

// func (p *Player) WatchTempo() {
//     for range (<-p.Transport.Status).UpdatedTempo {
//         p.Media.Grooverize(p.Transport)
//     }
// }
