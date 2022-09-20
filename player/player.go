package player

import (
	"fmt"
	"log"
	"math"
	"voop/clip"
	"voop/sync"
)

var (
	err             error
	oldMedia, media *clip.Media
	mNext           chan *clip.Media
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
	p.HotKey = make(chan int)
	mNext = make(chan *clip.Media)
	var k int
	// play it in cycle forever
play:
	for range p.Clock.Trigger {
		select {
		case m := <-mNext:
			p.Media = m
		default:
			// pass
		}
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

func (p *Player) SwitchMedia(path string) {
	go func() {
		media, err = clip.NewMedia(path, p.Transport)
		if err != nil {
			log.Fatal("error while opening media", err)
		}
		oldMedia = p.Media
		mNext <- media
		oldMedia.Close()
	}()

}
