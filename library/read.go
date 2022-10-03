package library

import (
	"math/rand"
	"time"
)

var (
	tmp int
)

type read struct {
	rightNow int
	size     int
}

// Navigation
func (r *read) now() int {
	return r.rightNow
}

func (r *read) random() {
	rand.Seed(time.Now().UnixNano())
	tmp = r.rightNow
	r.rightNow = rand.Intn(r.size)
	if tmp == r.rightNow {
		r.next()
	}
}

func (r *read) next() {
	r.rightNow = (r.rightNow + 1) % r.size
}

func (r *read) previous() {
	m := r.size - r.rightNow
	m = m % r.size
	r.rightNow = r.size - m - 1
}

func (r *read) Default() {
	r.rightNow = 0
}
