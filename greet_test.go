// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	testCases := []struct {
		given string
		want  string
	}{
		{given: "world", want: "hello world"},
		{given: "kurth4cker", want: "hello kurth4cker"},
		{given: "", want: "hello "},
	}

	for _, tc := range testCases {
		got := Greet(tc.given)
		if got != tc.want {
			t.Errorf("got %q, want %q, given %q",
				got, tc.want, tc.given)
		}
	}
}

func TestFhelloln(t *testing.T) {
	users := []string{"world", "kurth4cker", "emacs"}
	for _, user := range users {
		var buffer bytes.Buffer

		err := Fhelloln(&buffer, user)
		assertNotError(t, err)

		got := buffer.String()
		want := Greet(user)
		assertStrings(t, got, want)
	}
}
