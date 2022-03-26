package demo

import (
	"golang.org/x/exp/constraints"
)

func min[T constraints.Ordered](x, y T) T {
	if x < y {
		return x
	}
	return y
}

type Point []int32

func (p *Point) String() string {
	return "point"
}

func Scale[S ~[]E, E constraints.Integer](s S, c E) S {
	res := make(S, len(s))
	for i, v := range s {
		res[i] = v * c
	}
	return res
}

func Scales[T constraints.Integer](s []T, c T) []T {
	res := make([]T, len(s))
	for i, v := range s {
		res[i] = v * c
	}

	return res
}
