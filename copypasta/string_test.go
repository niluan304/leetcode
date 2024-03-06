package copypasta

import (
	"reflect"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func TestKMP(t *testing.T) {
	type args struct {
		str    []rune
		patten []rune
	}
	cases := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "",
			args: args{
				str:    []rune("bacbababadababacambabacaddababacasdsd"),
				patten: []rune("ababaca"),
			},
			want: []int{10, 26},
		},
	}

	fs := []func(text []rune, patten []rune) (pos []int){
		KMP,
		KMP2,
	}
	for _, f := range fs {
		name := tests.FuncName(f)
		t.Run(name, func(t *testing.T) {
			for _, c := range cases {
				t.Run(c.name, func(t *testing.T) {
					got := f(c.args.str, c.args.patten)
					if !reflect.DeepEqual(got, c.want) {
						t.Errorf("%v() = %v, want %v", name, got, c.want)
					}
				})
			}
		})
	}
}
