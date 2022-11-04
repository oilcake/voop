package player

import (
	"image"
	"image/color"
	"log"

	"gocv.io/x/gocv"
)

var (
	targetSize image.Point
)

type Resizer struct {
	size         image.Point
	aspects      image.Point
	window_width int
}

func NewResizer(width, height, window_width int, aspects image.Point) *Resizer {
	size := image.Point{width, height}
	return &Resizer{
		size:         size,
		aspects:      aspects,
		window_width: window_width,
	}
}

func (r *Resizer) DumbResize(frame *gocv.Mat, clipAspect float64) {

	targetSize.X = r.size.X
	targetSize.Y = int(float64(targetSize.X) / clipAspect)
	gocv.Resize(*frame, frame, targetSize, 0.0, 0.0, gocv.InterpolationDefault)
}

func (r *Resizer) ResizeAndPad(frame *gocv.Mat, dim image.Point) {
	// make correct borders:
	pads := r.center(frame, dim)
	log.Println("pads", pads.X, pads.Y)
	gocv.CopyMakeBorder(*frame, frame, pads.X, pads.X, pads.Y, pads.Y, gocv.BorderDefault, color.RGBA{0, 0, 0, 0})
	x_delta := pads.X * 2
	y_delta := pads.Y * 2
	// new_dim := (dim.X + x_delta, dim.Y + y_delta)
	// get size of the image fitted into window:
	dim = r.get_resized_dim(dim.X+x_delta, dim.Y+y_delta)
	gocv.Resize(*frame, frame, dim, 0, 0, gocv.InterpolationNearestNeighbor)
}

func (r *Resizer) get_resized_dim(w, h int) (dim image.Point) {
	width := float64(w)
	height := float64(h)
	aspect := width / height
	ratio := float64(r.window_width) / width
	width = float64(width) * ratio
	height = width / aspect
	dim = image.Point{int(width), int(height)}
	return
}

func (r *Resizer) center(frame *gocv.Mat, dim image.Point) (pads image.Point) {

	width, height := float64(dim.X), float64(dim.Y)
	aspect := width / height
	// define the axis that should be padded (0-x, 1-y):
	dominant := aspect_ratio(r.aspects) < aspect
	pads = r.pad(boolToInt(dominant), width, height)
	return pads
}

func boolToInt(b bool) (i int) {
	switch b {
	case true:
		i = 1
	case false:
		i = 0
	}
	return
}

func aspect_ratio(aspects image.Point) float64 {
	x, y := aspects.X, aspects.Y
	return float64(x) / float64(y)
}

func (r *Resizer) pad(dominant int, width float64, height float64) (pads image.Point) {
	var (
		top, left, difference float64
	)
	if dominant == 0 {
		// pad x:
		desired_width := get_width_from_height(height, aspect_ratio(r.aspects))
		difference := desired_width - width
		top = 0
		left = difference / 2
	}
	if dominant == 1 {
		// pad y:
		desired_height := get_height_from_width(width, aspect_ratio(r.aspects))
		difference = desired_height - height
		top = difference / 2
		left = 0
	}
	pads = image.Point{int(top), int(left)}
	return
}

func get_height_from_width(width float64, aspect_ratio float64) float64 {
	return width / aspect_ratio
}

func get_width_from_height(height float64, aspect_ratio float64) float64 {
	return height * aspect_ratio
}
