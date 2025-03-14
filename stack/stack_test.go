// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package stack_test

import (
	"fmt"
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
	"codeberg.org/kurth4cker/go-sample/stack"
)

func TestArray(t *testing.T) {
	ints := new(stack.Array[int])
	assert.True(t, ints.IsEmpty())

	ints.Push(123)
	assert.False(t, ints.IsEmpty())

	ints.Push(456)
	value, _ := ints.Pop()
	assert.Equal(t, value, 456)

	value, _ = ints.Pop()
	assert.Equal(t, value, 123)
	assert.True(t, ints.IsEmpty())

	_, ok := ints.Pop()
	assert.False(t, ok)
}

func TestArray_Len(t *testing.T) {
	t.Run("Empty", func(t *testing.T) {
		ints := new(stack.Array[int])
		assert.Equal(t, ints.Len(), 0)
	})

	t.Run("Pushes", func(t *testing.T) {
		cases := []int{0, 3, 5, 10}

		for _, c := range cases {
			t.Run(fmt.Sprint(c), func(t *testing.T) {
				ints := new(stack.Array[int])
				for range c {
					ints.Push(8)
				}

				want := c
				got := ints.Len()

				assert.Equal(t, got, want)
			})
		}
	})

	t.Run("Pops", func(t *testing.T) {
		cases := []struct {
			pushes int
			pops   int
		}{
			{10, 5},
			{2, 1},
			{2, 2},
		}

		for _, c := range cases {
			name := fmt.Sprintf("%d-%d", c.pushes, c.pops)
			t.Run(name, func(t *testing.T) {
				ints := new(stack.Array[int])
				for range c.pushes {
					ints.Push(2)
				}
				for range c.pops {
					ints.Pop()
				}

				want := c.pushes - c.pops
				got := ints.Len()

				assert.Equal(t, got, want)
			})
		}
	})
}
