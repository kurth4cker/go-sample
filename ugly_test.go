package sample_test

import (
	"fmt"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
	"codeberg.org/kurth4cker/go-sample/assert"
)

func TestIsUgly(t *testing.T) {
	trueCases := []int{12, 1, 8, 5832}
	falseCases := []int{7, 11, 22, 33, 21}

	for _, tc := range trueCases {
		t.Run(fmt.Sprint(tc), func(t *testing.T) {
			got := sample.IsUgly(tc)
			assert.True(t, got)
		})
	}

	for _, fc := range falseCases {
		t.Run(fmt.Sprint(fc), func(t *testing.T) {
			got := sample.IsUgly(fc)
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
		{7, 8},
		{150, 5832},
	}

	for _, c := range cases {
		t.Run(fmt.Sprint(c.given), func(t *testing.T) {
			got := sample.NthUglyNumber(c.given)
			assert.Equal(t, got, c.want)
		})
	}
}

func BenchmarkNthUglyNumber(b *testing.B) {
	cases := []int{1, 16, 256, 1352}
	for _, c := range cases {
		b.Run(fmt.Sprint(c), func(b *testing.B) {
			for range b.N {
				sample.NthUglyNumber(c)
			}
		})
	}
}
