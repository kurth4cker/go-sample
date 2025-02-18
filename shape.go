// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Cristopher James
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

// General shape interface.
//
// Deprecated: Will be internal.
type Shape interface {
	Area() float64
}

// Deprecated: Will be internal.
type Rectangle struct {
	Width  float64
	Height float64
}

// Return area of given rectangle.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}
