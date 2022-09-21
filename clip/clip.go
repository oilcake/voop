package clip

import (
	"errors"
	"fmt"
	"image"
	"log"
	"math"
	"math/rand"
	"time"

	"voop/sync"

	"gocv.io/x/gocv"
)

const (
	clipWidth = 400.0
)

type ImgShape struct {
	W     float64
	H     float64
	AspRt float64
}

type Media struct {
	Name         string
	V            *gocv.VideoCapture
	Duration     float64
	Framecount   float64
	F            *gocv.Mat //this is a current frame object
	LoopLen      float64   // Media pattern's length
	Shape        *ImgShape
	Multiple     float64
	forward      bool
	palindrome   bool
	phase        float64
	alteredPhase float64
	offset       float64
	shiftedPhase float64
	antiphase    float64
	timepoint    float64
}

func NewMedia(filename string, t *sync.Transport) (m *Media, err error) {
	fmt.Println()
	log.Println("opening ", filename)
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
		Name:         filename,
		V:            clip,
		Duration:     msDur,
		Framecount:   framecount,
		F:            &f,
		LoopLen:      0.0,
		Shape:        shape,
		Multiple:     1.0,
		forward:      true,
		palindrome:   false,
		offset:       0,
		phase:        0,
		alteredPhase: 0,
		shiftedPhase: 0,
	}
	media.Grooverize(t)
	return media, nil
}

func (m *Media) PalindromemordnilaP(t *sync.Transport) {
	switch {
	case !m.palindrome:
		m.Multiple = m.Multiple * 2.0
		m.palindrome = true
		m.Grooverize(t)
	case m.palindrome:
		m.Multiple = m.Multiple * 0.5
		m.palindrome = false
		m.Grooverize(t)
	}
}

func PositiveMod(m, n float64) float64 {
	return math.Mod(math.Mod(m, n)+n, n)
}

func (m *Media) calcFrame() (frame float64) {
	m.antiphase = -m.phase
	switch {
	case m.palindrome:
		m.phase = m.phase*2.0 - 1.0
		frame = m.Framecount * math.Abs(m.phase)
		break
	case m.forward:
		m.shiftedPhase = m.phase - m.offset + m.timepoint
		m.shiftedPhase = PositiveMod(m.shiftedPhase, 1.0)
		frame = m.Framecount * m.shiftedPhase
		break
	case !m.forward:
		m.antiphase += m.offset
		m.shiftedPhase = m.timepoint + m.antiphase
		m.shiftedPhase = PositiveMod(m.shiftedPhase, 1.0)
		frame = m.Framecount * m.shiftedPhase
		break
	}
	fmt.Printf("\rCurrent frame %06d, phase %.2f, offset %.2f, shiftedPhase %.2f, alteredPhase %.2f",
		int(frame), m.phase, m.offset, m.shiftedPhase, m.alteredPhase)
	return
}

func (m *Media) Swap() {
	m.forward = !m.forward
	m.timepoint = m.shiftedPhase
	m.offset = m.phase
}

func (m *Media) Zero() {
	m.offset = m.phase
	m.timepoint = 0
}

func (m *Media) Jump() {
	rand.Seed(time.Now().UnixNano())
	m.offset = rand.Float64()
}
func (m *Media) Frame(phase float64) gocv.Mat {
	m.phase = phase
	// find number of frame
	f := m.calcFrame()
	// rewind
	m.V.Set(gocv.VideoCapturePosFrames, f)
	// read it
	m.V.Read(m.F)
	if m.F.Empty() {
		log.Fatal("Unable to read VideoCaptureFile")
	}
	// resize
	scaledSize := image.Point{clipWidth, int(math.Round(clipWidth / m.Shape.AspRt))}
	m.F = Resize(m.F, scaledSize)
	return *m.F
}

func Resize(frame *gocv.Mat, size image.Point) *gocv.Mat {

	gocv.Resize(*frame, frame, size, 0.0, 0.0, gocv.InterpolationDefault)
	return frame
}

func (m *Media) BarsTotal(duration float64, Measure uint8) (f float64) {
	beatsTotal := int(Round((m.Duration / duration), float64(Measure)))
	log.Println("beats total is", beatsTotal)
	bars := beatsTotal / int(Measure)
	defer log.Println("bars total", bars)
	if bars < 1.0 {
		return 1.0
	}
	return float64(bars)
}

func Round(x, unit float64) float64 {
	return math.Round(x/unit) * unit
}

func (m *Media) Squarize(b float64) (length float64) {
	// finding a "square" pattern - bars count to fit media duration in musical time
	sqLog := math.Log2(b)
	// and return the needed power to make it square
	return math.Pow(2, math.Round(sqLog))

}

func (m *Media) Grooverize(t *sync.Transport) {
	b := m.BarsTotal(t.OneBeatDurationInMs(), t.TimeSignature.Measure)
	if b > 4.0 {
		m.LoopLen = b
	} else {
		m.LoopLen = m.Squarize(b)
	}
	m.LoopLen = m.LoopLen * float64(t.TimeSignature.Measure) * m.Multiple
	log.Println("pattern", m.LoopLen)
}

func (m *Media) UpdateRate() {

}
func (m *Media) Close() {
	m.V.Close()
	m.F.Close()
}
