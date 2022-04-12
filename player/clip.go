package player

import (
	"errors"
	"image"
	"log"
	"math"

	"voop/sync"

	"gocv.io/x/gocv"
)

type ImgShape struct {
	W     float64
	H     float64
	AspRt float64
}

type Media struct {
	Name       string
	V          *gocv.VideoCapture
	Duration   float64
	Framecount float64
	F          *gocv.Mat //this is a current frame object
	P          float64
	Shape      ImgShape
}

func NewMedia(filename string) (m *Media, err error) {
	clip, err := gocv.VideoCaptureFile(filename)
	if !clip.IsOpened() {
		return nil, errors.New("Error opening video stream or file")
	}
	if err != nil {
		return nil, err
	}
	// get number of frames in video
	framecount := clip.Get(gocv.VideoCaptureFrameCount)
	fps := clip.Get(gocv.VideoCaptureFPS)
	msDur := framecount / fps
	log.Printf("duration in seconds is %v of type %T\n", msDur, msDur)
	width := clip.Get(gocv.VideoCaptureFrameWidth)
	height := clip.Get(gocv.VideoCaptureFrameHeight)
	shape := ImgShape{
		W:     width,
		H:     height,
		AspRt: width / height,
	}

	f := gocv.NewMat()
	media := &Media{
		Name:       filename,
		V:          clip,
		Duration:   msDur,
		Framecount: framecount,
		F:          &f,
		P:          0.0,
		Shape:      shape,
	}
	return media, nil
}

func (m *Media) Frame(phase float64) gocv.Mat {
	// find frame
	f := phase * m.Framecount
	// rewind
	m.V.Set(gocv.VideoCapturePosFrames, f)
	// read video frame
	m.V.Read(m.F)
	if m.F.Empty() {
		log.Fatal("Unable to read VideoCaptureFile")
	}
	// resize frame
	scaledSize := image.Point{1000, int(math.Round(1000.0 / m.Shape.AspRt))}
	gocv.Resize(*m.F, m.F, scaledSize, 0.0, 0.0, gocv.InterpolationDefault)
	return *m.F
}

func (m *Media) BarsTotal(BeatDuration float64, Measure uint8) (f float64) {
	beatsTotal := int(math.Round(m.Duration / BeatDuration))
	log.Println("beats total is", beatsTotal)
	bars := beatsTotal / int(Measure)
	defer log.Println("bars total", bars)
	if bars < 1.0 {
		return 1.0
	}
	return float64(bars)
}

func (m *Media) Pattern(t *sync.Transport) {
	log.Println("Tempo is now", (<-t.St).Bpm)
	// finding a "square" pattern - bars count to fit media duration in musical time
	sqLog := math.Log2(m.BarsTotal(t.BeatDur(), t.TimeSignature.Measure))
	// I am adding 2 to pow to round to a greater value just because it feels better
	length := math.Pow(2, math.Round(sqLog)+2)
	log.Println("pattern", length)
	m.P = length
}

func (m *Media) Position(t *sync.Transport) float64 {
	if (<-t.St).D {
		m.Pattern(t)
	}
	phase := math.Mod((<-t.St).Beat, m.P) / m.P
	return phase
}

func (m *Media) Close() {
	m.V.Close()
	m.F.Close()
}
