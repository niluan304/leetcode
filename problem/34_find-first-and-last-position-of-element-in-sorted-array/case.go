package find_first_and_last_position_of_element_in_sorted_array

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums   []int
	target int
}

type Output []int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 8,
			},
			Except: []int{3, 4},
		},
		{
			Input: Input{
				nums:   []int{5, 7, 7, 8, 8, 10},
				target: 6,
			},
			Except: []int{-1, -1},
		},
		{
			Input: Input{
				nums:   []int{},
				target: 0,
			},
			Except: []int{-1, -1},
		},
	}
}

type Func func(nums []int, target int) []int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
			in.target,
		))
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

	return reflect.DeepEqual(x, y)
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