package move_zeroes

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
			Input:  Input{nums: []int{0, 1, 0, 3, 12}},
			Except: []int{1, 3, 12, 0, 0},
		},
		{
			Input:  Input{nums: []int{0}},
			Except: []int{0},
		},
	}
}

type Func func(nums []int)

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {

	return func(in Input) Output {
		var nums = make([]int, len(in.nums))
		copy(nums, in.nums)

		f(nums)
		return Output(nums)
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

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
