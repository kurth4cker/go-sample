// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Sample Go repository for trying out any module related tech
package sample

import (
	"fmt"
	"io"
)

// Return a formatted greeting to given user.
func Sgreet(user string) string {
	return "hello " + user
}

// Fhelloln formats given name and prints it to given writer.
// Formatting is done by [Greet].
// See [fmt.Fprintln].
func Fhelloln(w io.Writer, user string) error {
	_, err := fmt.Fprint(w, Sgreet(user))
	if err != nil {
		return err
	}
	return nil
}
