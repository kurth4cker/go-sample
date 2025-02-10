// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import (
	"testing"
)

func TestGreet(t *testing.T) {
	assertGotWant := func(t testing.TB, got, want string) {
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}

	t.Run("greet to 'world'", func(t *testing.T) {
		want := "hello world"
		got := Greet("world")
		assertGotWant(t, got, want)
	})

	t.Run("greet to anyone else", func(t *testing.T) {
		want := "hello kurth4cker"
		got := Greet("kurth4cker")
		assertGotWant(t, got, want)
	})

	t.Run("empty string is not special", func(t *testing.T) {
		want := "hello "
		got := Greet("")
		assertGotWant(t, got, want)
	})
}
