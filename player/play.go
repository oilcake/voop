package player

import (
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	"time"
	"voop/clip"
	"voop/library"

	"gocv.io/x/gocv"
)

// type Unit[T any] T

// Play functions
type Reader interface {
	Now() int
	Random()
	Next()
	Previous()
	What(index int) interface{}
}

func PlayLibrary(p *Player, r Reader) {
	for {
		// choose folder
		element := r.What(r.Now())
		path, ok := element.(*string)
		if !ok {
			log.Fatal("type conversion failed")
		}
		// preload set from folder
		set, err := library.NewSet(path, p.Transport)
		if err != nil {
			log.Fatal("cannot preload folder", err)
		}
		// (don't forget to close everything)
		defer library.CloseSet(set)

		action := PlaySet(p, set)
		fmt.Println(action)
		switch action {
		case "rnd":
			r.Random()
		case "next":
			r.Next()
		case "prev":
			r.Previous()
		case "stop":
			return
		}
	}
}

func PlaySet(p *Player, r Reader) (action string) {

	for {
		// play media
		element := r.What(r.Now())
		media, ok := element.(*clip.Media)
		if !ok {
			log.Fatal("type conversion failed")
		}
		action = PlayMedia(media, p) // until any keyboard action
		fmt.Println(action)
		switch action {
		case "rnd":
			r.Random()
		case "next":
			r.Next()
		case "prev":
			r.Previous()
		case "stop":
			return
		case "nextChapter":
			return "next"
		case "prevChapter":
			return "prev"
		case "randomChapter":
			return "rnd"
		}
	}

}

func PlayMedia(media *clip.Media, p *Player) (action string) {
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
		switch v {
		// quit
		case 27:
			action = "stop"
			break play
		// fullscreen
		case getKey('f'):
			p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
		case getKey('g'):
			p.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowNormal)
		// ratio
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
		// clip navigation
		case getKey('/'):
			action = "rnd"
			break play
		case getKey('.'):
			action = "next"
			break play
		case getKey(','):
			action = "prev"
			break play
		// folder navigation
		case getKey('§'):
			action = "randomChapter"
			break play
		case getKey(']'):
			action = "nextChapter"
			break play
		case getKey('['):
			action = "prevChapter"
			break play
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
