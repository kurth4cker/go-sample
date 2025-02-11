// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import "testing"

func BenchmarkFibonacci10(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fibonacci(10)
	}
}
