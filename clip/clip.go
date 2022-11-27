package clip

import (
	"errors"
	"fmt"
	"log"

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
	Shape      *ImgShape
	// creepy loop stuff
	transport    *sync.Transport
	LoopLen      float64 // Media pattern's length
	multiple     float64
	RateX        chan float64
	forward      bool
	palindrome   bool
	plndrmTrigga chan struct{}
	phase        float64
	offset       float64
	shiftedPhase float64
	antiphase    float64
	timepoint    float64
	hardSync     bool
	pldShift     float64
}

func NewMedia(filename string, t *sync.Transport) (m *Media, err error) {
	fmt.Println()
	log.SetFlags(log.Lshortfile)
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
	msDur := framecount / fps * 1000
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
		LoopLen:    0.0,
		Shape:      shape,
		multiple:   1.0,
		transport:  t,
		RateX:      make(chan float64),
		// creepy loop stuff
		forward:      true,
		palindrome:   false,
		plndrmTrigga: make(chan struct{}),
		offset:       0,
		phase:        0,
		shiftedPhase: 0,
		hardSync:     false,
	}
	media.Grooverize()
	return media, nil
}

func (m *Media) Frame() *gocv.Mat {
	// find number of frame and rewind
	m.V.Set(gocv.VideoCapturePosFrames, m.calcFrame())
	// read frame and place it into frame object
	m.V.Read(m.F)
	if m.F.Empty() {
		log.Fatal("Unable to read VideoCaptureFile")
	}
	return m.F
}

func (m *Media) calcFrame() (frame float64) {
	frame = m.Framecount * m.Position()
	return
}

func (m *Media) Close() {
	m.V.Close()
	m.F.Close()
}
