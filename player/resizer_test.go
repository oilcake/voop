package player

import (
	"testing"
	"voop/clip"

	"github.com/stretchr/testify/assert"
)

func TestAspectRatio(t *testing.T) {
	type testCase struct {
		name     string
		in       imgRect
		expected float64
	}
	testCases := []testCase{
		{
			name:     "aspect test 01",
			in:       imgRect{13, 8},
			expected: 1.625,
		},
		{
			name:     "aspect test 02",
			in:       imgRect{8, 8},
			expected: 1.0,
		},
		{
			name:     "aspect test 03",
			in:       imgRect{16, 8},
			expected: 2.0,
		},
		{
			name:     "aspect test 03",
			in:       imgRect{4, 8},
			expected: 0.5,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			actual := aspectRatio(tc.in)
			assert.Equal(t, tc.expected, actual, "should be equal")
		})
	}
}

func TestCenterPads(t *testing.T) {
	type testCase struct {
		name         string
		inTo         imgRect
		inFrom       clip.ImgShape
		expectedPads imgRect
	}

	testCases := []testCase{
		{name: "pads test 01", inTo: imgRect{16, 8}, inFrom: clip.ImgShape{W: 13, H: 8}, expectedPads: imgRect{1.5, 0}},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			resiza := NewResizer(int(tc.inTo.X), int(tc.inTo.Y))
			resiza.ResizeFrom(tc.inFrom)
			actual := resiza.center()
			assert.Equal(t, tc.expectedPads, actual, "should be equal")
		})
	}
}
