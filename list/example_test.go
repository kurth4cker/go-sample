// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package list_test

import (
	"fmt"
	"slices"

	"github.com/kurth4cker/go-sample/list"
)

func Example() {
	lst := list.List[int]{}
	for i := range 4 {
		lst.Push(i)
	}

	numbers := slices.Collect(lst.All())
	fmt.Println(numbers)
	// Output: [3 2 1 0]
}
