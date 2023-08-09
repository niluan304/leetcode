package find_peak_element

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums []int
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums: []int{1, 2, 3, 1},
			},
			Except: []int{2},
		},
		{
			Input: Input{
				nums: []int{1, 2, 1, 3, 5, 6, 4},
			},
			Except: []int{1, 5},
		},
	}
}

type Func func(nums []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return []int{f(in.nums)}
	}
}

func AddCases(c func() []tests.Case[Input, Output]) {
	_cases := cases()
	cases = func() []tests.Case[Input, Output] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...Func) {
	funcs = append(funcs, tests.NewFuncWithAdaptor(adaptor, f...)...)
}

func Equal(x, y Output) bool {
	if len(x) < len(y) {
		x, y = y, x
	}

	if len(y) != 1 {
		return false
	}

	var set = map[int]struct{}{}
	for _, v := range x {
		set[v] = struct{}{}
	}
	for _, v := range y {
		_, ok := set[v]
		if !ok {
			return false
		}
	}

	return true
}

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}

func Bench(b *testing.B) {
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  Equal,
	})
}
