// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Sample Go repository for trying out any module related tech
package sample

import (
	"fmt"
	"io"
)

// Return a formatted greeting to given user.
func Greet(user string) string {
	return "hello " + user
}

// Fhello formats given user names and prints them to given writer.
//
// Formatting is done by [Greet]. Every formatted value is written line by line.
// See [fmt.Fprintln].
func Fhello(w io.Writer, users ...string) error {
	for _, user := range users {
		_, err := fmt.Fprintln(w, Greet(user))
		if err != nil {
			return err
		}
	}
	return nil
}
