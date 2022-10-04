package player

import (
	"voop/clip"
	"voop/sync"
)

type Player struct {
	*sync.Clock
	*Window
	Resizer
	Media  *clip.Media
	HotKey chan int
}

func (p *Player) PlayMedia(media chan *clip.Media) {
	p.HotKey = make(chan int)
	var (
		k          int
		m, garbage *clip.Media
	)
	// play it in cycle forever
play:
	for range p.Clock.Trigger {
		select {
		case m = <-media:
			// ok, we've got a media
			garbage = p.Media
			p.Media = m
			garbage.Close()
		default:
			// pass
		}
		// retrieve frame
		img := p.Media.Frame()
		// resize
		p.Resizer.DumbResize(img, p.Media.Shape.AspRt)
		// and display it
		p.Window.DisplayFrame(img)
		k = p.Window.WaitKey(19)
		switch {
		case k == 27:
			break play
		case k != 27 && k != -1:
			p.HotKey <- k
		}
	}
}
