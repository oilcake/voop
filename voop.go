package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"time"
	"voop/player"
	"voop/sync"
)

func main() {
	// initialize clock
	clock := sync.NewClock(40 * time.Millisecond)

	// initialize transport
	t, err := sync.NewTransport()
	if err != nil || t == nil {
		log.Fatal("can't start transport", err)
	}

	// initialize a display
	window := player.NewWindow("Voop")
	defer window.Close()

	// make a player instance
	p := player.Player{Clock: clock, Transport: t, Window: window}

	// read video folder
	var folder = flag.String("folder", "./", "path to your videos")
	flag.Parse()
	fmt.Println(*folder)

	// preload a bunch of files
	set, err := player.NewSet(folder, t)
	if err != nil {
		log.Fatal("cannot preload folder", err)
	}
	// (don't forget to close everything)
	defer player.CloseSet(set)

	// and play it forever
	player.PlaySet(&p, set)

	// Bye
	log.Println("closing files")
}
