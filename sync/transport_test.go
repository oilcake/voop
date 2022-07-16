package sync

import (
	"net"
	"testing"
)

var testMessage string = "{ :peers 0 :bpm 120.000000 :start 24619121429 :beat 110716.064800 }"
var sTest Status

func TestPing(t *testing.T) {
	conn, _ := net.Dial(protocol, address)
	response, _ := Ping(conn, "status")
	got := *response
	got = got[:8]
	want := "{ :peers"
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

func TestParser(t *testing.T) {
	want := Status{Peers: 0, Bpm: 120.000000, Start: 24619121429, Beat: 110716.064800}
	Parse(&testMessage, &sTest)
	got := sTest
	if got != want {
		t.Errorf("got %+v\n want %+v\n", got, want)
	}
}

func BenchmarkParser(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Parse(&testMessage, &sTest)
	}
}
