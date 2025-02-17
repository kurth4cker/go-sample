// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Sample Go repository for trying out any module related tech
package sample

import (
	"fmt"
	"io"
	"os"
)

// Fhello formats given name and prints it to given writer.
// Formatting is done by [Shello].
// See [fmt.Fprint].
func Fhello(w io.Writer, user string) {
	fmt.Fprint(w, Shello(user))
}

// Fhelloln formats and prints a greeting string to given writer.
// See [fmt.Fprintln].
func Fhelloln(w io.Writer, user string) {
	fmt.Fprintln(w, Shello(user))
}

// Helloln writes formatted greeting and a newline to [os.Stdout].
func Helloln(user string) {
	Fhelloln(os.Stdout, user)
}

// Return a formatted greeting string to given user.
// User may be any string including empty string.
func Shello(user string) string {
	return "hello " + user
}
