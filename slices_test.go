package generics

import (
	"reflect"
	"strings"
	"testing"
)

func TestFilter(t *testing.T) {
	type testCase[T any] struct {
		name string
		s    []T
		f    func(T) bool
		want []T
	}

	case1 := testCase[string]{
		"website",
		[]string{"http://foo.com", "http://bar.com", "https://go.dev"},
		func(s string) bool {
			return strings.HasPrefix(s, "https://")
		},
		[]string{"https://go.dev"},
	}

	case2 := testCase[int]{
		"number",
		[]int{1, 2, 3, 4},
		func(i int) bool {
			return i%2 == 0
		},
		[]int{2, 4},
	}

	t.Run(case1.name, func(t *testing.T) {
		if got := Filter(case1.s, case1.f); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Filter() = %v, want %v", got, case1.want)
		}
	})

	t.Run(case2.name, func(t *testing.T) {
		if got := Filter(case2.s, case2.f); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Filter() = %v, want %v", got, case2.want)
		}
	})
}
