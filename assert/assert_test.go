// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert_test

import (
	"testing"

	"codeberg.org/kurth4cker/go-sample/assert"
)

func TestAssert(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		assert.Equal(t, 1, 1)
		assert.NotEqual(t, 1, 2)
	})

	t.Run("Strings", func(t *testing.T) {
		assert.Equal(t, "hello", "hello")
		assert.NotEqual(t, "hello", "bye")
	})
}
