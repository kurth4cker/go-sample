// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample_test

import (
	"bytes"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
)

func TestFhelloln(t *testing.T) {
	user := "world"
	buffer := new(bytes.Buffer)
	sample.Fhelloln(buffer, user)

	got := buffer.String()
	want := sample.Shello(user) + "\n"

	if got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}

func TestShello(t *testing.T) {
	testCases := []struct {
		given string
		want  string
	}{
		{given: "world", want: "hello world"},
		{given: "kurth4cker", want: "hello kurth4cker"},
		{given: "", want: "hello "},
	}

	for _, tc := range testCases {
		got := sample.Shello(tc.given)
		if got != tc.want {
			t.Errorf("got %q, want %q, given %q",
				got, tc.want, tc.given)
		}
	}
}
