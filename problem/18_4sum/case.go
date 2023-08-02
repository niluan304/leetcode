package four_sum

import (
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums   []int
	target int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums:   []int{1, 0, -1, 0, -2, 2},
				target: 0,
			},
			Except: Output{{-1, 0, 0, 1}, {-2, -1, 1, 2}, {-2, 0, 0, 2}},
		},
		{
			Input: Input{
				nums:   []int{2, 2, 2, 2, 2},
				target: 8,
			},
			Except: Output{{2, 2, 2, 2}},
		},
	}
}

type Func func(nums []int, target int) [][]int

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

func Sort(x Output) (x1 Output) {
	sort.Slice(x, func(i, j int) bool {
		if len(x[i]) == 0 {
			return true
		}
		if len(x[j]) == 0 {
			return false
		}
		if x[i][0] != x[j][0] {
			return x[i][0] < x[j][0]
		}
		if x[i][1] != x[j][1] {
			return x[i][1] < x[j][1]
		}
		return x[i][2] < x[j][2]
	})

	for i, v := range x {
		if len(v) == 0 {
			continue
		}
		x1 = append(x1, x[i])
	}

	return x1
}

func Equal(x, y Output) bool {
	x1, y1 := Sort(x), Sort(y)

	if len(x1) != len(y1) {
		return false
	}

	for i, v := range x1 {
		if len(v) != len(y1[i]) {
			return false
		}
		for j, v1 := range v {
			if v1 != y1[i][j] {
				return false
			}
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
