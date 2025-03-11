// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package stack_test

import (
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
