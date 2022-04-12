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

	// start Ableton Link watcher
	go sync.Link(sync.NewConnection(), t.St)

	// initialize random generator
	rand.Seed(time.Now().UnixNano())
	// preload a bunch of files
	b, err := player.OpenFolder(folder)
	if err != nil {
		log.Fatal("cannot preload folder", err)
	}

	// don't forget to close everything
	defer player.CloseFolder(b)

	// and play randoms from it forever
	for {
		media := b[rand.Intn(len(b)-1)]
		PlayMedia(media, t, window) // until any key
	}
}

func PlayMedia(media *player.Media, t *sync.Transport, window *gocv.Window) {
	// who is it
	log.Println("playing file", media.Name)
	// perform init calculations on it
	if media.P == 0.0 {
		media.Pattern(t)
	}
	// in cycle forever
	for {
		// calculate a playing phase
		ph := media.Position(t)
		fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", (<-t.St).Beat, ph)
		// to retrieve specific frame
		img := media.Frame(ph)
		// and display it
		window.IMShow(img)
		v := window.WaitKey(1)
		if v >= 0 {
			break
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
