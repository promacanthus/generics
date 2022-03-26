package main

import "golang.org/x/exp/constraints"

// Max returns the maximal value from a slice.
func Max[T constraints.Ordered](s []T) T {
	var max T
	if len(s) == 0 {
		return max
	}

	max = s[0]
	for _, v := range s {
		if v > max {
			max = v
		}
	}
	return max
}

// Min returns the minimal value from a slice.
func Min[T constraints.Ordered](s []T) T {
	var min T
	if len(s) == 0 {
		return min
	}

	min = s[0]
	for _, v := range s {
		if v < min {
			min = v
		}
	}
	return min
}
