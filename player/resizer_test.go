package player

import (
	"image"
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

func TestCenter(t *testing.T) {
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

func TestResize(t *testing.T) {

	type testCase struct {
		name          string
		inTo          imgRect
		inFrom        clip.ImgShape
		expectedShape image.Point
	}

	display := imgRect{16, 16}

	testCases := []testCase{
		{name: "square", inTo: display, inFrom: clip.ImgShape{W: 8, H: 8}, expectedShape: display.AsImagePoint()},
		{name: "longer X", inTo: display, inFrom: clip.ImgShape{W: 13, H: 8}, expectedShape: display.AsImagePoint()},
		{name: "longer Y", inTo: display, inFrom: clip.ImgShape{W: 13, H: 15}, expectedShape: display.AsImagePoint()},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			resiza := NewResizer(int(tc.inTo.X), int(tc.inTo.Y))
			resiza.ResizeFrom(tc.inFrom)
			actual := resiza.outIntShape
			assert.Equal(t, tc.expectedShape, actual, "should be equal")
		})
	}
}

func TestGetWidthFromHeight(t *testing.T) {
	type testCase struct {
		name          string
		inHeight      float64
		inAspect      float64
		expectedWidth float64
	}

	testCases := []testCase{
		{
			name:          "width from height 01",
			inHeight:      8,
			inAspect:      2,
			expectedWidth: 16,
		},
		{
			name:          "width from height 02",
			inHeight:      9,
			inAspect:      1.4444444444444444,
			expectedWidth: 13,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			actual := getWidthfromHeight(tc.inHeight, tc.inAspect)
			assert.Equal(t, tc.expectedWidth, actual, "should be equal")
		})
	}
}

func TestGetHeightFromWidth(t *testing.T) {
	type testCase struct {
		name           string
		inWidth        float64
		inAspect       float64
		expectedHeight float64
	}

	testCases := []testCase{
		{
			name:           "height from width 01",
			inWidth:        16,
			inAspect:       2,
			expectedHeight: 8,
		},
		{
			name:           "height from width 02",
			inWidth:        13,
			inAspect:       1.4444444444444444,
			expectedHeight: 9,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			tc := tc
			t.Parallel()
			actual := getHeightFromWidth(tc.inWidth, tc.inAspect)
			assert.Equal(t, tc.expectedHeight, actual, "should be equal")
		})
	}
}
