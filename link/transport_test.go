package link

import (
	"net"
	"testing"
)

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
	var st Status
	message := "{ :peers 0 :bpm 120.000000 :start 24619121429 :beat 110716.064800 }"
	want := Status{Peers: 0, Bpm: 120.000000, Start: 24619121429, Beat: 110716.064800}
	Parse(&message, &st)
	got := st
	if got != want {
		t.Errorf("got %+v\n want %+v\n", got, want)
	}
}
