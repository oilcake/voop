package main

import (
	// "flag"
	"fmt"
	// "log"
	_ "net/http/pprof"
	"time"

	// "voop/clip"
	"voop/config"
	// "voop/library"
	// "voop/player"
	"voop/sync"
	// "voop/vj"
)

// Navigation
// Do i really need that????
type Reader interface {
	Now() int
	Random()
	Next()
	Previous()
	What(index int) interface{}
}

type Connector interface {
	Connect() chan interface{}
}

const (
	fps = 26
	gap = 1 * time.Second / fps // time between two frames
)

func main() {
	// read config and get an actions map
	k := config.ReadConfig()
	fmt.Println(k)
	// initialize clock
	clock := sync.NewClock(gap)

	clock.Start()
	for range clock.Trigger {
		select {
		case <-clock.Trigger:
			fmt.Println("tick")
		}
	}
	/*
		// initialize transport
		t, err := sync.NewTransport(clock)
		if err != nil || t == nil {
			log.Fatal("can't start transport", err)
		}
		t.Start()
		// player instance
		p := player.Player{Transport: t}

		// read video folder
		var folder = flag.String("folder", "./", "path to your videos")
		flag.Parse()
		fmt.Println(*folder)

		// preload a bunch of files
		lib, err := library.NewLibrary(folder, t)
		if err != nil {
			log.Fatal("cannot preload library", err)
		}

		// call something from library
		// 		// choose folder
		element := lib.What(lib.Now())
		path, ok := element.(*string)
		if !ok {
			log.Fatal("type cast failed")
		}

		set, err := library.NewSet(path, p.Transport)
		if err != nil {
			log.Fatal("cannot preload folder", err)
		}

		// (don't forget to close everything)
		defer library.CloseSet(set)

		// get media
		element = set.What(set.Now())
		media, ok := element.(*clip.Media)
		if !ok {
			log.Fatal("type conversion failed")
		}

		// send it to player
		p.Media = media

		// initialize VJ
		vj := vj.NewVJ(&p, lib, *k)

		// initialize a display
		window := player.NewWindow("Voop")

		// engage rendering
		go p.Play(window.Frame)

		go vj.Gig(window.Feedback)

		// and play it forever in main loop
		for range window.Frame {
			log.Println("window")
			window.Output()
		}

		// (don't forget to close everything!!!)
		// stopPlayer <- struct{}{}
		log.Println("stop window")
		window.Close()

		// player.PlayLibrary(&p, lib)

		// Bye
		log.Println("ciao")
	*/
}
