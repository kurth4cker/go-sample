// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample_test

import (
	"bytes"
	"testing"

	"codeberg.org/kurth4cker/go-sample"
)

func TestFhelloln(t *testing.T) {
	cases := []struct {
		given string
		want  string
	}{
		{given: "world", want: "hello world\n"},
		{given: "kurth4cker", want: "hello kurth4cker\n"},
		{given: "", want: "hello \n"},
	}

	for _, c := range cases {
		buffer := new(bytes.Buffer)
		sample.Fhelloln(buffer, c.given)

		got := buffer.String()
		want := c.want
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
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
