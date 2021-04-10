package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func Test_whatIndexOfSpy(t *testing.T) {
	r := require.New(t)

	testCases := []struct {
		name string

		arr []int

		expResp int
	}{
		{
			name: "second",

			arr: []int{11, 13, 11, 11},

			expResp: 2,
		},
		{
			name: "first",

			arr: []int{1, 4, 4, 4, 4},

			expResp: 1,
		},
		{
			name: "fifth",

			arr: []int{3, 3, 3, 3, 10, 3, 3, 3, 3, 3},

			expResp: 5,
		},
		{
			name: "third",

			arr: []int{20, 20, 10},

			expResp: 3,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			actual := whatIndexOfSpy(tc.arr)

			r.Equal(tc.expResp, actual)
		})
	}
}
