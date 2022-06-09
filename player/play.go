package player

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"

	"gocv.io/x/gocv"
)

// Play functions
type Navigator interface {
	Now() *Media
	Random()
	Next()
	Previous()
}

func PlaySet(p *Player, n Navigator) {

	for {
		// play media
		media := n.Now()
		action := PlayMedia(media, p) // until any keyboard action
		fmt.Println(action)
		switch action {
		case "rnd":
			n.Random()
		case "next":
			n.Next()
		case "prev":
			n.Previous()
		case "stop":
			return
		}
	}

}

func PlayMedia(media *Media, p *Player) (action string) {
	// who is it
	log.Println("\nplaying file", media.Name)
	// and play it in cycle forever
play:
	for {
		select {
		case <-p.Clock.Trigger:
			// calculate a playing phase
			ph := media.Position(p.Transport)
			fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", (<-p.Transport.Status).Beat, ph)
			// to retrieve specific frame
			img := media.Frame(ph)
			// and display it
			p.Window.IMShow(img)
		}
		v := p.Window.WaitKey(1)
		log.Println("key pressed ", v)
		switch v {
		case 27:
			action = "stop"
			break play
		case getKey('-'):
			media.Multiple = media.Multiple * 2.0
			media.Pattern(p.Transport)
		case getKey('_'):
			media.Multiple = media.Multiple * 1.5
			media.Pattern(p.Transport)
		case getKey('='):
			media.Multiple = media.Multiple / 2.0
			media.Pattern(p.Transport)
		case getKey('+'):
			media.Multiple = media.Multiple / 1.5
			media.Pattern(p.Transport)
		case getKey('0'):
			media.Multiple = 1.0
			media.Pattern(p.Transport)
		case getKey('/'):
			action = "rnd"
			break play
		case getKey('.'):
			action = "next"
			break play
		case getKey(','):
			action = "prev"
			break play
		case getKey('f'):
			p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
		case getKey('g'):
			p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowNormal)
		}
	}
	return
}

func ChooseRandomFile(path *string) (string, error) {

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return "", err
	}
	log.Println("files total", len(files))
	rand.Seed(time.Now().UnixNano())
	file := *path + "/" + files[rand.Intn(len(files)-1)].Name()
	fmt.Println()
	log.Printf("Playing file %v\n", file)
	return file, nil
}

func getKey(r rune) int {
	return int(r)
}
