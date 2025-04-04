// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package stack

type Interface[T any] interface {
	Pop() (T, bool)
	Push(v T)
}
