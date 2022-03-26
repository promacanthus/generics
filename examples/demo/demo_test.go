package demo

import (
	"fmt"
	"reflect"
	"testing"

	"golang.org/x/exp/constraints"
)

func Test_gmin(t *testing.T) {
	type args[T constraints.Ordered] struct {
		x T
		y T
	}
	type cases[T constraints.Ordered] struct {
		name string
		args args[int]
		want T
	}
	tests := []cases[int]{
		// TODO: Add test cases.
		{
			name: "test",
			args: args[int]{1, 2},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := min(tt.args.x, tt.args.y); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("gmin() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestScale(t *testing.T) {
	p := &Point{1, 2, 3}
	r := Scale(*p, 2)
	fmt.Println(r.String())
}

func TestScales(t *testing.T) {
	p := &Point{1, 2, 3}
	r := Scales(*p, 2)
	fmt.Println(r)
}
