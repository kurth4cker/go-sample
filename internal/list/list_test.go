// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package list

import (
	"slices"
	"testing"
)

func TestList_All(t *testing.T) {
	list := List[int]{}
	list.Push(1)
	list.Push(2)
	list.Push(3)

	got := slices.Collect(list.All())
	want := []int{3, 2, 1}

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
