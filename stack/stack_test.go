// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package stack_test

import (
	"fmt"
	"slices"
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
	"codeberg.org/kurth4cker/go-sample/stack"
)

func TestArray_IsEmpty(t *testing.T) {
	t.Run("Fresh", func(t *testing.T) {
		ints := new(stack.Array[int])
		assert.True(t, ints.IsEmpty())
	})

	t.Run("PurePush", func(t *testing.T) {
		ints := new(stack.Array[int])
		for range 8 {
			ints.Push(0)
		}
		assert.False(t, ints.IsEmpty())
	})

	t.Run("PurePop", func(t *testing.T) {
		ints := new(stack.Array[int])
		for range 4 {
			ints.Pop()
		}
		assert.True(t, ints.IsEmpty())
	})
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

func TestArray_Push(t *testing.T) {
	ints := new(stack.Array[int])
	interface_PushOrder(t, ints, 1, 2, 3, 4)
}

func TestArray_Pop(t *testing.T) {
	t.Run("correct order", func(t *testing.T) {
		ints := new(stack.Array[int])
		interface_PopOrder(t, ints, 1, 2, 3, 4)
	})

	t.Run("pop from empty array", func(t *testing.T) {
		ints := new(stack.Array[int])
		interface_PopEmpty(t, ints)
	})
}

func interface_PopEmpty[T any](t testing.TB, st stack.Interface[T]) {
	_, ok := st.Pop()
	if ok {
		t.Error("got true but want false")
	}
}

func interface_PopOrder[T comparable](t testing.TB, st stack.Interface[T], values ...T) {
	interface_PushOrder(t, st, values...)
}

func interface_PushOrder[T comparable](t testing.TB, st stack.Interface[T], values ...T) {
	for _, value := range values {
		st.Push(value)
	}

	got := make([]T, 0, len(values))
	for {
		value, ok := st.Pop()
		if !ok {
			break
		}
		got = append(got, value)
	}
	want := slices.Clone(values)
	slices.Reverse(want)

	if !slices.Equal(got, want) {
		t.Errorf("got %v, want %v", got, want)
	}
}
