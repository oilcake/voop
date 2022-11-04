package clip

import (
	"log"
	"math"
)

const mediaIsTooLong = 8 // do not grooverize anything longer than this

// Note that value is rounded
func (m *Media) ConvertMsToBeats(ms float64) (beats float64) {
	beats = math.Round(ms / m.transport.DurationOfOneBeatInMs())
	return
}

func (m *Media) BarsTotal(duration float64, BeatQuantity uint8) (f float64) {
	beatsTotal := m.ConvertMsToBeats(m.Duration)
	log.SetFlags(log.Lshortfile)
	log.Println("beats total is", beatsTotal)
	bars := beatsTotal / float64(BeatQuantity)
	defer log.Println("bars total", bars)
	return float64(bars)
}

func (m *Media) Squarize(b float64) (length float64) {
	// finding a "square" pattern - bars count to fit media duration in musical time
	sqLog := math.Log2(b)
	// and return the needed power to make it square
	return math.Pow(2, math.Round(sqLog))

}

func (m *Media) findLoopLength() (loopLen float64) {
	log.SetFlags(log.Lshortfile)
	log.Printf("one beat is %v milliseconds\n", m.transport.DurationOfOneBeatInMs())
	b := m.BarsTotal(m.transport.DurationOfOneBeatInMs(), m.transport.TimeSignature.BeatQuantity)
	if b > mediaIsTooLong {
		loopLen = b
	} else {
		loopLen = m.Squarize(b)
	}
	loopLen = loopLen * float64(m.transport.TimeSignature.BeatQuantity) * m.multiple
	log.Println("pattern", loopLen)
	return
}

func (m *Media) Grooverize() {
	m.LoopLen = m.findLoopLength()
}
