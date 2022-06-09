package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"time"
	"voop/player"
	"voop/sync"

	"gocv.io/x/gocv"
)

func main() {
	fmt.Println("Starting")

	// initialize clock
	clock := sync.NewClock(40 * time.Millisecond)

	// initialize transport
	t, err := sync.NewTransport()
	if err != nil || t == nil {
		log.Fatal("can't start transport", err)
	}
	log.Println("Transport is created")

	// initialize a display
	window := player.NewWindow("Voop")
	defer window.Close()

	p := player.Player{Clock: clock, Transport: t, Window: window}

	// read video folder
	var folder = flag.String("folder", "./", "path to video files")
	flag.Parse()
	fmt.Println(*folder)
	// preload a bunch of files
	set, err := player.NewSet(folder, t)
	if err != nil {
		log.Fatal("cannot preload folder", err)
	}

	// don't forget to close everything
	defer player.CloseSet(set)

	// and play randoms from it forever
	PlaySet(&p, set)
}

// Play functions
type Navigator interface {
	Now() *player.Media
	Random()
	Next()
	Previous()
}

func PlaySet(p *player.Player, n Navigator) {

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
		}
	}

}

func PlayMedia(media *player.Media, p *player.Player) (action string) {
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
		case getKey('-'):
			media.Multiple = media.Multiple * 2.0
			media.Pattern(p.Transport)
		case getKey('='):
			media.Multiple = media.Multiple / 2.0
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
