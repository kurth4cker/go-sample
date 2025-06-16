// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert_test

import (
	"testing"

	"github.com/kurth4cker/go-sample/assert"
)

func TestEqual(t *testing.T) {
	t.Run("Integers", func(t *testing.T) {
		assert.Equal(t, 1, 1)
	})

	t.Run("Strings", func(t *testing.T) {
		assert.Equal(t, "hello", "hello")
	})
}

func TestEqualFunc(t *testing.T) {
	eqTrueFunc := func(_, _ any) bool {
		return true
	}

	assert.EqualFunc(t, nil, nil, eqTrueFunc)
}

func TestTrue(t *testing.T) {
	assert.True(t, true);
}

func TestFalse(t *testing.T) {
	assert.False(t, false)
}
