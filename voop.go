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
	// read video folder
	var folder = flag.String("folder", "./", "path to video files")
	flag.Parse()
	fmt.Println(*folder)

	// initialize clock
	clock := sync.NewClock(40 * time.Millisecond)

	// initialize transport
	t, err := sync.NewTransport()
	if err != nil || t == nil {
		log.Fatal("can't start transport", err)
	}
	log.Println("Transport is created")

	// initialize a display
	window := gocv.NewWindow("Voop")
	defer window.Close()
	if window == nil {
		log.Fatal("Unable to create Window")
	}
	if !window.IsOpen() {
		log.Fatal("Window should have been open")
	}
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowNormal)
	window.SetWindowProperty(gocv.WindowPropertyAspectRatio, gocv.WindowKeepRatio)
	window.ResizeWindow(100, 100)

	// initialize random generator
	rand.Seed(time.Now().UnixNano())
	// preload a bunch of files
	b, err := player.OpenFolder(folder, t)
	if err != nil {
		log.Fatal("cannot preload folder", err)
	}

	// don't forget to close everything
	defer player.CloseFolder(b)

	// and play randoms from it forever
	for {
		media := b[rand.Intn(len(b)-1)]
		PlayMedia(media, t, window, clock) // until any key
	}
}

func PlayMedia(media *player.Media, t *sync.Transport, window *gocv.Window, clock *sync.Clock) {
	// who is it
	log.Println("playing file", media.Name)
	// and play it in cycle forever
play:
	for {
		select {
		case <-clock.Trigger:
			// calculate a playing phase
			ph := media.Position(t)
			fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", (<-t.Status).Beat, ph)
			// to retrieve specific frame
			img := media.Frame(ph)
			// and display it
			window.IMShow(img)
		}
		v := window.WaitKey(1)
		switch v {
		case getKey('q'):
			break play
		case getKey('f'):
			window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
		case getKey('g'):
			window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowNormal)
		}
	}
}

func ChooseRandomFile(path *string) (string, error) {

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return "", err
	}
	log.Println("files total", len(files))
	rand.Seed(time.Now().UnixNano())
	file := *path + "/" + files[rand.Intn(len(files)-1)].Name()
	log.Printf("Playing file %v\n", file)
	return file, nil
}

func getKey(r rune) int {
	return int(r)
}
