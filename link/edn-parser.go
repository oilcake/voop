package link

import (
	"olympos.io/encoding/edn"
)

type Status struct {
	Peers int
	Bpm   float32
	Start int64
	Beat  float64
}

func Parse(message string) (Status, error) {
	var status Status
	err := edn.Unmarshal([]byte(message), &status)
	return status, err
}
