package link_test

import (
	"testing"
	"voop/link"
)

func TestPing(t *testing.T) {
	got := link.Ping("status")[:7]
	want := "status "
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
