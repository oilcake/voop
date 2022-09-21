package player

import (
	"log"

	"gocv.io/x/gocv"
)

type Window struct {
	*gocv.Window
}

func NewWindow(name string) *Window {
	window := gocv.NewWindow(name)
	if window == nil {
		log.Fatal("Unable to create Window")
	}
	if !window.IsOpen() {
		log.Fatal("Window should have been open")
	}
	window.SetWindowProperty(gocv.WindowPropertyAutosize, gocv.WindowNormal)
	window.SetWindowProperty(gocv.WindowPropertyAspectRatio, gocv.WindowKeepRatio)
	window.ResizeWindow(100, 100)
	return &Window{window}
}

func (w *Window) Fullscreen() {
	w.Window.SetWindowProperty(gocv.WindowPropertyFullscreen, gocv.WindowFullscreen)
}
