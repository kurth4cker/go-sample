package sample_test

import (
	"fmt"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
	"codeberg.org/kurth4cker/go-sample/assert"
)

func TestNthUglyNumber(t *testing.T) {
	cases := []struct {
		given, want int
	}{
		{10, 12},
		{1, 1},
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
