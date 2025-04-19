// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert

import "testing"

// Assert equality with == operator
//
// If they are not equal call tb.Errorf.
func Equal[T comparable](tb testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// Same as [Equal] but test with given function.
func EqualFunc[T any](t testing.TB, got, want T, eq func(T, T) bool) {
	t.Helper()
	if !eq(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

// If given value is not true, fail.
func True(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

// Fail if value is false.
func False(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
