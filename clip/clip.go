package clip

import (
	"errors"
	"image"
	"log"
	"math"

	"voop/sync"

	"gocv.io/x/gocv"
)

const (
	clipWidth = 1000.0
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
	P          float64   // Media pattern's length
	Shape      *ImgShape
	Multiple   float64
}

func NewMedia(filename string, t *sync.Transport) (m *Media, err error) {
	// open file
	clip, err := gocv.VideoCaptureFile(filename)
	if !clip.IsOpened() {
		return nil, errors.New("Error opening video stream or file")
	}
	if err != nil {
		return nil, err
	}

	// fill video properties
	framecount := clip.Get(gocv.VideoCaptureFrameCount)
	fps := clip.Get(gocv.VideoCaptureFPS)
	msDur := framecount / fps
	width := clip.Get(gocv.VideoCaptureFrameWidth)
	height := clip.Get(gocv.VideoCaptureFrameHeight)
	shape := &ImgShape{
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
		Multiple:   1.0,
	}
	media.Pattern(t)
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
	scaledSize := image.Point{clipWidth, int(math.Round(clipWidth / m.Shape.AspRt))}
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

func (m *Media) Squarize(t *sync.Transport, b float64) (length float64) {
	// finding a "square" pattern - bars count to fit media duration in musical time
	sqLog := math.Log2(b)
	// and return the needed power to make it square
	return math.Pow(2, math.Round(sqLog))

}

func (m *Media) Pattern(t *sync.Transport) {
	b := m.BarsTotal(t.BeatDur(), t.TimeSignature.Measure)
	if b > 4.0 {
		m.P = b
	} else {
		m.P = m.Squarize(t, b)
	}
	m.P = m.P * float64(t.TimeSignature.Measure) * m.Multiple
	log.Println("pattern", m.P)

}

func (m *Media) Position(t *sync.Transport) float64 {
	st := <-t.Status

	if st.D {
		log.Println("Tempo is now", (<-t.Status).Bpm)
		m.Pattern(t)
	}
	phase := math.Mod(st.Beat, m.P) / m.P
	return phase
}

func (m *Media) Close() {
	m.V.Close()
	m.F.Close()
}
