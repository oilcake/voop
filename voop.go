package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math/rand"
	_ "net/http/pprof"
	"time"
	"voop/link"
	"voop/player"

	"gocv.io/x/gocv"
)

func main() {
	// read video folder
	var folder = flag.String("folder", "./", "path to video files")
	flag.Parse()
	fmt.Println(*folder)
	file, err := ChooseRandomFile(folder)
	if err != nil {
		log.Fatal("error while opening file", err)
	}
	fmt.Println()

	// initialize transport
	t, err := link.NewTransport()
	if err != nil || t == nil {
		log.Fatal("can't start transport", err)
	}
	// open video
	media, err := player.NewMedia(file)
	defer media.Close()
	// make window
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

	// play video in cycle forever
	for {
		ph := media.Position(t)
		fmt.Printf("\rCurrent beat is %.9f and phase is %.9f", t.St.Beat, ph)
		img := media.Frame(ph)
		window.IMShow(img)
		v := window.WaitKey(1)
		if v >= 0 {
			break
		}
		time.Sleep(time.Millisecond * 40)
	}
}

func ChooseRandomFile(path *string) (string, error) {

	files, err := ioutil.ReadDir(*path)
	if err != nil {
		return "", err
	}
	log.Println("files total", len(files))
	rand.Seed(time.Now().UnixNano())
	file := *path + "/" + files[rand.Intn(len(files))].Name()
	fmt.Printf("Playing file %v\n", file)
	return file, nil
}
