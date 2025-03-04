package sample

import (
	"fmt"
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
)

func TestIsUgly(t *testing.T) {
	trueCases := []int{1, 2, 3, 4, 5, 6, 8, 9, 10, 12, 3, 2, 1}
	falseCases := []int{7, 11}

	for _, tc := range trueCases {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			got := IsUgly(tc)
			assert.True(t, got)
		})
	}

	for _, fc := range falseCases {
		t.Run(fmt.Sprint(fc), func(t *testing.T) {
			got := IsUgly(fc)
			assert.False(t, got)
		})
	}
}

func TestNthUglyNumber(t *testing.T) {
	cases := []struct {
		given, want int
	}{
		{10, 12},
		{1, 1},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := NthUglyNumber(c.given)
			assert.Equal(t, got, c.want)
		})
	}
}

func BenchmarkNthUglyNumber(b *testing.B) {
	cases := []int{1, 16, 256, 1352}
	for _, c := range cases {
		b.Run(fmt.Sprint(c), func(b *testing.B) {
			for range b.N {
				NthUglyNumber(c)
			}
		})
	}
}
