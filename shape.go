// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Cristopher James
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import "math"

// General shape interface.
type Shape interface {
	Area() float64
}

// Circle with only radius.
type Circle struct {
	Radius float64
}

// Returns area of given circle with math.Pi.
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

type Rectangle struct {
	Width  float64
	Height float64
}

// Return area of given rectangle.
func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

type Triangle struct {
	Base   float64
	Height float64
}

// Return area of given triangle.
func (t Triangle) Area() float64 {
	return t.Base * t.Height * 0.5
}
