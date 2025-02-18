// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package internal

import "iter"

// Singly Linked List implementation.
type List[T any] struct {
	head *node[T]
}

// Returns all elements in reverse push order.
// For example if you push 1, 2, 3 respectively; when you call [slices.Collect]
// on All(), you will get [3 2 1].
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
func (l *List[T]) Push(v T) {
	elem := &node[T]{
		next: l.head,
		val:  v,
	}
	l.head = elem
}

type node[T any] struct {
	next *node[T]
	val  T
}
