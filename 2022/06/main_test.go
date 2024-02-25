package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

var testCases = []struct {
	input string
	want  int
}{
	{
		input: "bvwbjplbgvbhsrlpgdmjqwftvncz",
		want:  5,
	},
	{
		input: "nppdvjthqldpwncqszvftbrmjlhg",
		want:  6,
	},
	{
		input: "mjqjpqmgbljsphdztnvjfqwrcgsmlb",
		want:  7,
	},
	{
		input: "nznrnfrfntjfmvfwmzdfjlvtqnbhcprsg",
		want:  10,
	},
	{
		input: "zcfzfwzzqfrljwzlrfnpqdbhtmscgvjw",
		want:  11,
	},
	{
		input: "aaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaaabcd",
		want:  60,
	},
}

func TestFind(t *testing.T) {
	for i := range testCases {
		tc := &testCases[i]
		t.Run(tc.input, func(t *testing.T) {
			got := find(tc.input)
			assert.Equal(t, tc.want, got)
		})
	}
}

func BenchmarkFind(b *testing.B) {
	for i := range testCases {
		tc := testCases[i]
		b.Run(fmt.Sprint(tc.want), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				find(tc.input)
			}
		})
	}
}
