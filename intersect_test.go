package generics

import "testing"

func TestContains(t *testing.T) {
	type testCase[T comparable] struct {
		name   string
		args   []T
		target T
		want   bool
	}

	case1 := testCase[string]{"string", []string{"a", "b", "c"}, "b", true}
	case2 := testCase[int]{"int", []int{1, 2, 3}, 4, false}

	t.Run(case1.name, func(t *testing.T) {
		if got := Contains(case1.args, case1.target); got != case1.want {
			t.Errorf("Contains() = %v, want %v", got, case1)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Contains(case2.args, case2.target); got != case2.want {
			t.Errorf("Contains() = %v, want %v", got, case2)
		}
	})
}
