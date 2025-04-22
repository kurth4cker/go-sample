// SPDX-License-Identifier: MPL-2.0
// SPDX-FileCopyrightText: 2025 kurth4cker <kurth4cker@gmail.com>

package slices

import "slices"

func UnorderedEqual[S ~[]E, E comparable](s1, s2 S) bool {
	if len(s1) != len(s2) {
		return false
	}

	tmp := slices.Clone(s1)
	for _, elem := range s2 {
		idx := slices.Index(tmp, elem)
		if idx == -1 {
			return false
		}
		tmp = Remove(tmp, idx)
	}
	return true
}

func Remove[S ~[]E, E any](s S, idx int) S {
	if len(s) <= idx {
		return s
	}
	slice := slices.Clone(s[:idx])
	return append(slice, s[idx+1:]...)
}
