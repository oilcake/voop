package library

import (
	"math/rand"
)

var (
	SupportedTypes = [...]string{".mp4", ".mpg", ".mov", ".avi", ".wmv", ".mkv"}
)

type Read struct {
	RightNow int
	Size     int
}

// Navigation
func (r *Read) Now() int {
	return r.RightNow
}

func (r *Read) Random() {
	r.RightNow = rand.Intn(r.Size - 1)
}

func (r *Read) Next() {
	r.RightNow = (r.RightNow + 1) % r.Size
}

func (r *Read) Previous() {
	m := r.Size - r.RightNow
	m = m % r.Size
	r.RightNow = r.Size - m - 1
}

func (r *Read) Default() {
	r.RightNow = 0
}
