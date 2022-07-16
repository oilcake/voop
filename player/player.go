package player

import (
	"fmt"
	"log"
	"voop/clip"
	"voop/sync"
)

type Player struct {
	*sync.Transport
	*clip.Media
}

func (p *Player) Play(f Frame) {
	// play:
	for status := range p.Transport.Status {
		log.Println("Play")
		// select {
		// case <-stop:
		// 	return
		// default:
		// 	// pass
		// }
		// log.Println("stop player")
		// return
		// calculate a playing phase
		ph := p.Media.Position(p.Transport)
		fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", status.Beat, ph)
		// to retrieve specific frame
		img := p.Media.Frame(ph)
		f <- img
	}
	close(f)
}

// who is it
// log.Println("\nplaying file", media.Name)
// and play it in cycle forever
// switch v {
// quit
// case 27:
// action = "stop"
// break play
// fullscreen
// case getKey('f'):
// 	p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
// case getKey('g'):
// 	p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowNormal)
// // ratio
// case getKey('-'):
// 	media.Multiple = media.Multiple * 2.0
// 	media.Pattern(p.Transport)
// case getKey('_'):
// 	media.Multiple = media.Multiple * 1.5
// 	media.Pattern(p.Transport)
// case getKey('='):
// 	media.Multiple = media.Multiple / 2.0
// 	media.Pattern(p.Transport)
// case getKey('+'):
// 	media.Multiple = media.Multiple * 0.75
// 	media.Pattern(p.Transport)
// case getKey('0'):
// 	media.Multiple = 1.0
// 	media.Pattern(p.Transport)
// // clip navigation
// case getKey('/'):
// 	action = "rnd"
// 	break play
// case getKey('.'):
// 	action = "next"
// 	break play
// case getKey(','):
// 	action = "prev"
// 	break play
// // folder navigation
// case getKey('§'):
// 	action = "randomChapter"
// 	break play
// case getKey(']'):
// 	action = "nextChapter"
// 	break play
// case getKey('['):
// 	action = "prevChapter"
// 	break play
// 		}
// 	}
// 	return
// }
