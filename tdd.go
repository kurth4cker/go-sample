// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) * Fibonacci(n-2)
}
