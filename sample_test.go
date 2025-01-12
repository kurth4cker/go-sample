// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import "testing"

func TestGreet(t *testing.T) {
	want := "hello world"
	got := Greet("world")

	if want != got {
		t.Errorf("Greet(\"world\") = %q, want %q", got, want)
	}
}
