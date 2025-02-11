// SPDX-License-Identifier: MIT
// SPDX-FileCopyrightText: 2025 Cristopher James
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package sample

import "testing"

func TestArea(t *testing.T) {
	areaTests := []struct {
		name  string
		shape Shape
		want  float64
	}{
		{name: "Circle",
			shape: Circle{Radius: 10},
			want:  314.1592653589793},
		{name: "Rectangle",
			shape: Rectangle{Height: 12, Width: 6},
			want:  72},
		{name: "Rectangle",
			shape: Rectangle{Height: 10, Width: 10},
			want:  100},
		{name: "Triangle",
			shape: Triangle{Base: 12, Height: 6},
			want:  36},
	}

	for _, tc := range areaTests {
		t.Run(tc.name, func(t *testing.T) {
			got := tc.shape.Area()
			if got != tc.want {
				t.Errorf("%#v got %g, want %g",
					tc.shape, got, tc.want)
			}
		})
	}
}
