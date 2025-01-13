// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

// Sample Go repository for trying out any module related tech
package sample

import "fmt"

// Return a formatted greeting to given user.
func Greet(user string) string {
	return fmt.Sprintf("hello %s", user)
}
