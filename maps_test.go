package generics

import (
	"reflect"
	"testing"
)

func TestKeys(t *testing.T) {
	type testCase[K comparable, V any] struct {
		name string
		args map[K]V
		want []K
	}

	case1 := testCase[string, bool]{
		"string key",
		map[string]bool{
			"potato":  true,
			"cabbage": true,
			"carrot":  true},
		[]string{"potato", "cabbage", "carrot"},
	}
	case2 := testCase[int, string]{
		"int key",
		map[int]string{
			1: "strawberry",
			2: "raspberry",
			3: "blueberry",
		},
		[]int{1, 2, 3},
	}

	t.Run(case1.name, func(t *testing.T) {
		if got := Keys(case1.args); !reflect.DeepEqual(got, case1.want) {
			t.Errorf("Keys() = %v, want %v", got, case1.want)
		}
	})
	t.Run(case2.name, func(t *testing.T) {
		if got := Keys(case2.args); !reflect.DeepEqual(got, case2.want) {
			t.Errorf("Keys() = %v, want %v", got, case2.want)
		}
	})
}
