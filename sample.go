// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import "fmt"

// Return a formatted greeting to given user.
func Greet(user string) string {
	return fmt.Sprintf("hello %s", user)
}
