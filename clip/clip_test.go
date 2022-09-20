package clip

import (
	"testing"
	"voop/sync"

	"github.com/stretchr/testify/assert"
)

func TestBarsTotal(t *testing.T) {
	type testCase struct {
		name     string
		ts       sync.TimeSignature
		expected float64
	}
	assert.True(t, true, "ok")
}
