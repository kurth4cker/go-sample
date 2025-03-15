// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Sample Go repository for trying out any module related tech
//
// Deprecated: Use [greet] instead.
package sample

import (
	"io"

	"codeberg.org/kurth4cker/go-sample/greet"
)

// Fhelloln formats and prints a greeting string to given writer.
// See [fmt.Fprintln].
//
// Deprecated: Use [greet.Fhelloln] instead.
func Fhelloln(w io.Writer, user string) {
	greet.Fhelloln(w, user)
}

// Helloln writes formatted greeting and a newline to [os.Stdout].
//
// Deprecated: Use [greet.Helloln] instead.
func Helloln(user string) {
	greet.Helloln(user)
}

// Return a formatted greeting string to given user.
// User may be any string including empty string.
//
// Deprecated: Use [greet.Shello] instead.
func Shello(user string) string {
	return greet.Shello(user)
}
