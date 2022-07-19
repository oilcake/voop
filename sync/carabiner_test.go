package sync

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

const testMessage string = "{ :peers 0 :bpm 120.000000 :start 24619121429 :beat 110716.064800 }"

var sTest Status

func TestParser(t *testing.T) {
	crbnr := NewCarabiner()
	*crbnr.message = testMessage
	want := &Status{Peers: 0, Bpm: 120.000000, Start: 24619121429, Beat: 110716.064800}
	crbnr.parse()
	got := crbnr.st
	// if got != want {
	// 	t.Errorf("got %+v\n want %+v\n", got, want)
	// }
	assert.Equal(t, got, want, "should be equal")
}

func BenchmarkParser(b *testing.B) {
	crbnr := NewCarabiner()
	*crbnr.message = testMessage
	for i := 0; i < b.N; i++ {
		crbnr.parse()
	}
}
