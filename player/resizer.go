package player

import (
	"image"

	"gocv.io/x/gocv"
)

var (
	targetSize image.Point
)

type Resizer struct {
	size image.Point
}

func NewResizer(width, height int) *Resizer {
	size := image.Point{width, height}
	return &Resizer{size: size}
}

func (r *Resizer) DumbResize(frame *gocv.Mat, clipAspect float64) {

	targetSize.X = r.size.X
	targetSize.Y = int(float64(targetSize.X) / clipAspect)
	gocv.Resize(*frame, frame, targetSize, 0.0, 0.0, gocv.InterpolationDefault)
}
