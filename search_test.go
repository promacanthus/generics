package generic

import (
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func TestMax(t *testing.T) {
	type testCases[T constraints.Ordered] struct {
		name string
		args []T
		want T
	}

	case1 := testCases[int32]{"int32", []int32{1, 2, 3}, 3}
	case2 := testCases[float64]{"float64", []float64{1.2, 3.4, 5.6}, 5.6}

	t.Run(case1.name, func(t *testing.T) {
		if got := Max(case1.args); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Max() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Max(case2.args); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Max() = %v, want %v", got, case2)
		}
	})
}

func TestMin(t *testing.T) {
	type testCases[T constraints.Ordered] struct {
		name string
		args []T
		want T
	}

	case1 := testCases[int32]{"int32", []int32{1, 2, 3}, 1}
	case2 := testCases[float64]{"float64", []float64{1.2, 3.4, 5.6}, 1.2}

	t.Run(case1.name, func(t *testing.T) {
		if got := Min(case1.args); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Min() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Min(case2.args); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Min() = %v, want %v", got, case2)
		}
	})
}
