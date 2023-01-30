package main

import (
	"flag"
	"fmt"
	"log"
	"time"
	"voop/clip"
	"voop/config"
	"voop/player"
	"voop/sync"
	"voop/vj"
)

const (
	// should be moved to project's entity
	// currently only for testing purposes
	DefaultBeatQuantity = 5
	DefaultDivisor      = 4
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

	fmt.Println()
	fmt.Println(conf.Display)
	fmt.Println()

	// and get an actions map from it
	k := config.CollectShortCuts(conf)
	fmt.Println(k)

	// initialize clock
	clock := sync.NewClock(40 * time.Millisecond)
	defer close(clock.Trigger)

	// establish connection with carabiner
	cnn := sync.NewConnection()
	crbnr := sync.NewCarabiner(cnn)

	// and start Link with it
	lnk := sync.NewLink(crbnr)

	// set time signature
	ts := &sync.TimeSignature{
		BeatQuantity: DefaultBeatQuantity,
		Divisor:      DefaultDivisor,
	}
	// initialize transport
	t := sync.NewTransport(lnk, ts)

	// initialize display
	window := player.NewWindow("Voop")
	defer window.Close()

	// instance of resizer object - needed to keep clip's aspect ratio
	reszr := player.NewResizer(conf.Size.Width, conf.Size.Height)

	// player instance
	p := player.Player{Clock: clock, Window: window, Resizer: reszr}

	// call VJ
	m := make(chan *clip.Media)
	vj := vj.VJ{Player: p, Config: conf, Shortcuts: *k, Transport: t, Media: m}

	// preload a bunch of files
	vj.OpenLibrary(folder)

	// listen for hotkeys
	go vj.WaitForAction()

	// will play until somebody hit Esc
	vj.Player.PlayMedia(vj.Media)

	// Bye
	log.SetFlags(log.Lshortfile)
	log.Println("\n\nciao")
}
