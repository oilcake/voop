package main

import (
	"flag"
	"fmt"
	"log"
	_ "net/http/pprof"
	"time"
	"voop/clip"
	"voop/config"
	"voop/player"
	"voop/sync"
	"voop/vj"
)

func main() {
	// read video folder
	var folder = flag.String("folder", "./samples", "path to your videos")
	var confFile = flag.String("config", "./config.yml", "configuration")
	flag.Parse()
	fmt.Println(*folder, *confFile)

	// read config
	conf := config.ReadConfig(*confFile)
	fmt.Println(conf)
	fmt.Println()
	fmt.Println(conf.Supported)

	// and get an actions map from it
	k := config.CollectShortCuts(conf)
	fmt.Println(k)
	// initialize clock
	clock := sync.NewClock(40 * time.Millisecond)
	defer close(clock.Trigger)

	// initialize transport
	t, err := sync.NewTransport()
	if err != nil || t == nil {
		log.Fatal("can't start transport", err)
	}

	// initialize a display
	window := player.NewWindow("Voop")
	defer window.Close()

	// make a player instance
	p := player.Player{Clock: clock, Window: window}

	// call VJ
	m := make(chan *clip.Media)
	vj := vj.VJ{Player: p, Config: conf, Shortcuts: *k, Transport: t, Media: m}
	// preload a bunch of files
	vj.OpenLibrary(folder)
	// listen for key presses
	go vj.WaitForAction()

	// and play it forever
	vj.Player.PlayMedia(vj.Media)

	// Bye
	log.SetFlags(log.Lshortfile)
	log.Println()
	log.Println("ciao")
}
