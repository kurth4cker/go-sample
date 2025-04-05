// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package assert

import "testing"

func Equal[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func EqualFunc[T any](t testing.TB, got, want T, eq func(T, T) bool) {
	t.Helper()
	if !eq(got, want) {
		t.Errorf("got %+v, want %+v", got, want)
	}
}

func True(t testing.TB, got bool) {
	t.Helper()
	if !got {
		t.Errorf("got %v, want true", got)
	}
}

func False(t testing.TB, got bool) {
	t.Helper()
	if got {
		t.Errorf("got %v, want false", got)
	}
}
