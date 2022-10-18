package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testMessage string = "{ :peers 0 :bpm 120.000000 :start 24619121429 :beat 110716.064800 }"

// tcp connection stub
type mockedConnection struct{}

func (m mockedConnection) getStatus() (response string) {
	return testMessage
}

func TestParser(t *testing.T) {
	mC := mockedConnection{}
	crbnr := NewCarabiner(mC)
	crbnr.Sync()
	want := &Status{Bpm: 120.000000, Beat: 110716.064800}
	got := crbnr.St
	assert.Equal(t, want, got, "should be equal")
}

func BenchmarkParser(b *testing.B) {
	mC := mockedConnection{}
	crbnr := NewCarabiner(mC)
	for i := 0; i < b.N; i++ {
		crbnr.Sync()
	}
}
