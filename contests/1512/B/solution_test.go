package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_findRectangle(t *testing.T) {
	r := require.New(t)

	testCases := []struct {
		name string

		firstX, firstY, secondX, secondY int
		size                             int

		expThirdX, expThirdY, expFourthX, expFourthY int
	}{
		{
			name: "4x4: 1",

			firstX:  2,
			firstY:  0,
			secondX: 0,
			secondY: 2,
			size:    4,

			expThirdX:  0,
			expThirdY:  0,
			expFourthX: 2,
			expFourthY: 2,
		},
		{
			name: "2x2: 1",

			firstX:  0,
			firstY:  0,
			secondX: 1,
			secondY: 1,
			size:    2,

			expThirdX:  1,
			expThirdY:  0,
			expFourthX: 0,
			expFourthY: 1,
		},
		{
			name: "2x2: 2",

			firstX:  1,
			firstY:  0,
			secondX: 1,
			secondY: 1,
			size:    2,

			expThirdX:  0,
			expThirdY:  0,
			expFourthX: 0,
			expFourthY: 1,
		},
		{
			name: "3x3: 1",

			firstX:  0,
			firstY:  0,
			secondX: 2,
			secondY: 0,
			size:    3,

			expThirdX:  0,
			expThirdY:  1,
			expFourthX: 2,
			expFourthY: 1,
		},
		{
			name: "5x5: 1",

			firstX:  2,
			firstY:  1,
			secondX: 1,
			secondY: 3,
			size:    5,

			expThirdX:  1,
			expThirdY:  1,
			expFourthX: 2,
			expFourthY: 3,
		},
		{
			name: "4x4: 2",

			firstX:  0,
			firstY:  2,
			secondX: 0,
			secondY: 3,
			size:    4,

			expThirdX:  1,
			expThirdY:  2,
			expFourthX: 1,
			expFourthY: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			thirdX, thirdY, fourthX, fourthY := findRectangle(tc.firstX, tc.firstY, tc.secondX, tc.secondY, tc.size)

			r.Equal(tc.expThirdX, thirdX)
			r.Equal(tc.expThirdY, thirdY)
			r.Equal(tc.expFourthX, fourthX)
			r.Equal(tc.expFourthY, fourthY)
		})
	}
}

func Test_printMatrix(t *testing.T) {
	printMatrix(0, 2, 0, 3, 1, 2, 1, 3, 4)
}
