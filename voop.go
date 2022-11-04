package main

import (
	"flag"
	"fmt"
	"image"
	"log"
	"time"
	"voop/clip"
	"voop/config"
	"voop/player"
	"voop/sync"
	"voop/vj"
)

const (
	DefaultBeatQuantity = 5 // THIS IS A STUB!!!!
	DefaultDivisor      = 4 // THIS IS A STUB!!!!
	clipWidth           = 300.0
	windowWidth         = 600
)

var (
	scaledSize image.Point
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

	// create video FX engine (that currently will just resize your videos)
	reszr := player.NewResizer(conf.Size.Width, conf.Size.Height, windowWidth, image.Point{16, 9})

	// make a player instance
	p := player.Player{Clock: clock, Window: window, Resizer: *reszr}

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
	log.Println("\n\nciao")
}
