package player

import (
	"image"
	"image/color"
	"voop/clip"

	"gocv.io/x/gocv"
)

var (
	targetSize image.Point
)

type imgRect struct {
	X, Y float64
}

func (i *imgRect) AsImagePoint() image.Point {
	return image.Point{int(i.X), int(i.Y)}
}

type Resizer struct {
	from imgRect
	to   imgRect
}

func NewResizer(width, height int) *Resizer {
	return &Resizer{
		to: imgRect{float64(width), float64(height)},
	}
}

func (r *Resizer) ResizeAndPad(frame *gocv.Mat) {
	// make correct borders:
	pads := r.center()
	left, top := pads.AsImagePoint().X, pads.AsImagePoint().Y
	// yeah pretty odd order of axes
	gocv.CopyMakeBorder(*frame, frame, top, top, left, left, gocv.BorderConstant, color.RGBA{0, 0, 0, 0})
	x_delta := pads.X * 2
	y_delta := pads.Y * 2
	// get size of the image fitted into window:
	desiredShape := r.getResizedDim(r.from.X+x_delta, r.from.Y+y_delta)
	gocv.Resize(*frame, frame, desiredShape.AsImagePoint(), 0, 0, gocv.InterpolationNearestNeighbor)
}

func (r *Resizer) getResizedDim(width, height float64) (dim imgRect) {
	aspect := width / height
	ratio := r.to.X / width
	width = width * ratio
	height = getHeightFromWidth(width, aspect)
	dim = imgRect{width, height}
	return
}

func (r *Resizer) center() (pads imgRect) {
	// define the axis that should be padded (0-x, 1-y):
	width, height := r.from.X, r.from.Y
	var (
		top, left, difference float64
	)
	switch aspectRatio(r.to) > aspectRatio(r.from) {
	case true:
		// pad x:
		desired_width := getWidthfromHeight(height, aspectRatio(r.to))
		difference = desired_width - width
		top = 0
		left = difference / 2
	case false:
		// pad y:
		desired_height := getHeightFromWidth(width, aspectRatio(r.to))
		difference = desired_height - height
		top = difference / 2
		left = 0
	}
	pads = imgRect{left, top}
	return
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

func aspectRatio(shape imgRect) float64 {
	return shape.X / shape.Y
}

func (r *Resizer) ResizeFrom(sh clip.ImgShape) {
	r.from = imgRect{sh.W, sh.H}
}

func getHeightFromWidth(width float64, aspectRatio float64) float64 {
	return width / aspectRatio
}

func getWidthfromHeight(height float64, aspectRatio float64) float64 {
	return height * aspectRatio
}
