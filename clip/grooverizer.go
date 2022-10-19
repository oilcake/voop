package clip

import (
	"log"
	"math"
)

func (m *Media) BarsTotal(duration float64, BeatQuantity uint8) (f float64) {
	beatsTotal := int(Round((m.Duration / duration), float64(BeatQuantity)))
	log.SetFlags(log.Lshortfile)
	log.Println("beats total is", beatsTotal)
	bars := beatsTotal / int(BeatQuantity)
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

func (m *Media) findLoopLength() (loopLen float64) {
	log.SetFlags(log.Lshortfile)
	log.Println("one beat is ", m.transport.OneBeatDurationInMs())
	b := m.BarsTotal(m.transport.OneBeatDurationInMs(), m.transport.TimeSignature.BeatQuantity)
	if b > 4.0 {
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
