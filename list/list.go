// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Package list implements Singly Linked List.
package list

import "iter"

// Singly Linked List implementation.
// The zero value for a List is an empty List ready to use.
type List[T any] struct {
	head *node[T]
}

// Returns all elements in list.
func (l *List[T]) All() iter.Seq[T] {
	return func(yield func(T) bool) {
		for e := l.head; e != nil; e = e.next {
			if !yield(e.val) {
				return
			}
		}
	}
}

// Push given item to head of [List].
// When you iterate over list, pushed items will be reverse order.
// For example:
//
//	lst := List[int]{}
//	lst.Push(1)
//	lst.Push(2)
//	lst.Push(3)
//	elems := slices.Collect(lst.All())
//	# elems: [3 2 1]
func (l *List[T]) Push(v T) {
	elem := &node[T]{
		next: l.head,
		val:  v,
	}
	l.head = elem
}

// Node implements Singly Linked List node.
type node[T any] struct {
	next *node[T]
	val  T
}
