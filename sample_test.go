// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import (
	"fmt"
	"testing"
)

func TestGreet(t *testing.T) {
	format := "hello %s"
	names := []string{"world", "kurth4cker"}

	for _, name := range names {
		want := fmt.Sprintf(format, name)
		got := Greet(name)

		if want != got {
			t.Errorf("got = %q, want %q", got, want)
		}
	}
}
