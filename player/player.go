package player

import (
	"fmt"
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

func (p *Player) PlayMedia() {
	p.HotKey = make(chan int)
	var k int
	// play it in cycle forever
play:
	for range p.Clock.Trigger {
		// calculate a playing phase
		ph := p.Media.Position(p.Transport)
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
