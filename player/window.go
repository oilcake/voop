package player

import (
	"log"

	"gocv.io/x/gocv"
)

type Frame chan gocv.Mat

type Window struct {
	*gocv.Window
	Frame
	Feedback chan int
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
	return &Window{window, make(Frame, 1), make(chan int)}
}

func (w *Window) Output() {
	// and display it
	w.Window.IMShow(<-w.Frame)
	w.Feedback <- w.Window.WaitKey(1)
}
