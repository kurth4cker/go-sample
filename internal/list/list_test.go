// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package list

import (
	"slices"
	"testing"
)

func TestList_All(t *testing.T) {
	want := []int{1, 2, 3}
	lst := newList(want...)

	got := slices.Collect(lst.All())

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func TestList_Push(t *testing.T) {
	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)

	var got []int
	for node := list.head; node != nil; node = node.next {
		got = append(got, node.val)
	}
	want := []int{3, 2, 1}

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}

func newList[T any](elems ...T) List[T] {
	var lst List[T]
	for idx := len(elems) - 1; idx >= 0; idx-- {
		elem := elems[idx]
		lst.Push(elem)
	}
	return lst
}
