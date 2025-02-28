// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package list_test

import (
	"slices"
	"testing"

	"codeberg.org/kurth4cker/go-sample/list"
)

func TestList(t *testing.T) {
	t.Run("All()", func(t *testing.T) {
		want := []int{1, 2, 3}
		lst := newList(want...)

		got := slices.Collect(lst.All())

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("Push()", func(t *testing.T) {
		list := list.List[int]{}
		list.Push(1)
		list.Push(2)
		list.Push(3)

		got := slices.Collect(list.All())
		want := []int{3, 2, 1}

		if !slices.Equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

func newList[T any](elems ...T) list.List[T] {
	var lst list.List[T]
	for idx := len(elems) - 1; idx >= 0; idx-- {
		elem := elems[idx]
		lst.Push(elem)
	}
	return lst
}
