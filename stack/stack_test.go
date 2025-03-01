// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package stack_test

import (
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
	"codeberg.org/kurth4cker/go-sample/stack"
)

func TestArray(t *testing.T) {
	t.Run("Int", func(t *testing.T) {
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
	})

	t.Run("String", func(t *testing.T) {
		strings := new(stack.Array[string])
		assert.True(t, strings.IsEmpty())

		strings.Push("hello")
		assert.False(t, strings.IsEmpty())

		strings.Push("bye")
		value, _ := strings.Pop()
		assert.Equal(t, value, "bye")

		value, _ = strings.Pop()
		assert.Equal(t, value, "hello")
		assert.True(t, strings.IsEmpty())
	})
}
