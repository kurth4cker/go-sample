// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sort_test

import (
	"math/rand"
	"slices"
	"testing"

	"codeberg.org/kurth4cker/go-sample/sort"
)

func TestSortedMerge(t *testing.T) {
	cases := []struct {
		s1, s2 []int
	}{
		{
			s1: []int{1, 3},
			s2: []int{2},
		},
	}

	for _, c := range cases {
		want := slices.Concat(c.s1, c.s2)
		slices.Sort(want)
		got := sort.Merge(c.s1, c.s2)
		if !slices.Equal(got, want) {
			t.Errorf("got %+v, want %+v", got, want)
		}
	}
}

func BenchmarkSortedMerge1024(b *testing.B) {
	var s1, s2 []int
	for range 1024 {
		s1 = append(s1, rand.Int())
		s2 = append(s2, rand.Int())
	}
	b.ResetTimer()
	for range b.N {
		sort.Merge(s1, s2)
	}
}
