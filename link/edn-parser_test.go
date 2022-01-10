package link_test

import (
	"testing"
	"voop/link"
)

func TestParser(t *testing.T) {
	// want := link.Status{peers: 0, bpm: 120.000000, start: 24619121429, beat: 110716.064800}
	want := link.Status{Peers: 0, Bpm: 120.000000, Start: 24619121429, Beat: 110716.064800}
	got, _ := link.Parse("{ :peers 0 :bpm 120.000000 :start 24619121429 :beat 110716.064800 }")
	if got != want {
		t.Errorf("got %+v\n want %+v\n", got, want)
	}
}
