package player

import (
	"image"
	"image/color"
	"voop/clip"

	"gocv.io/x/gocv"
)

type imgRect struct {
	X, Y float64
}

func (i *imgRect) AsImagePoint() image.Point {
	return image.Point{int(i.X), int(i.Y)}
}

type Resizer struct {
	from        imgRect
	to          imgRect
	pads        image.Point
	outIntShape image.Point
}

func NewResizer(width, height int) *Resizer {
	return &Resizer{
		to: imgRect{float64(width), float64(height)},
	}
}

func (r *Resizer) ResizeAndPad(frame *gocv.Mat) {
	// pretty odd order of axes, remember that
	gocv.CopyMakeBorder(*frame, frame, r.pads.Y, r.pads.Y, r.pads.X, r.pads.X, gocv.BorderConstant, color.RGBA{0, 0, 0, 0})
	gocv.Resize(*frame, frame, r.outIntShape, 0, 0, gocv.InterpolationArea)
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

func aspectRatio(shape imgRect) float64 {
	return shape.X / shape.Y
}

func (r *Resizer) ResizeFrom(sh clip.ImgShape) {
	r.from = imgRect{sh.W, sh.H}
	// make correct borders:
	pads := r.center()
	r.pads = pads.AsImagePoint()
	r.outIntShape = r.to.AsImagePoint()
}

func getHeightFromWidth(width, aspectRatio float64) float64 {
	return width / aspectRatio
}

func getWidthfromHeight(height, aspectRatio float64) float64 {
	return height * aspectRatio
}
