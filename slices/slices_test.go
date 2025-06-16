// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package slices_test

import (
	stdslices "slices"
	"testing"

	"github.com/kurth4cker/go-sample/slices"
)

func TestRemove(t *testing.T) {
	s1 := []int{1, 2, 3, 4}

	got := slices.Remove(s1, 2)
	want := []int{1, 2, 4}
	if !stdslices.Equal(got, want) {
		t.Errorf("got slice %v, want %v, given %v", got, want, s1)
	}
}

func TestUnorderedEqual(t *testing.T) {
	cases := []struct {
		name   string
		s1, s2 []int
		want   bool
	}{
		{
			name: "equal slices",
			s1:   []int{1, 2, 3, 4, 2},
			s2:   []int{3, 2, 1, 2, 4},
			want: true,
		},
		{
			name: "unequal slices",
			s1:   []int{1, 2, 2, 2},
			s2:   []int{1, 2, 2, 3},
			want: false,
		},
		{
			name: "different lengths",
			s1:   []int{1, 1, 1},
			s2:   []int{1, 1},
			want: false,
		},
	}

	for _, c := range cases {
		t.Run(c.name, func(t *testing.T) {
			got := slices.UnorderedEqual(c.s1, c.s2)
			if got != c.want {
				t.Errorf(`got %v, but want %v,
    values:
        s1 = %v
        s2 = %v`, got, c.want, c.s1, c.s2)
			}
		})
	}
}
