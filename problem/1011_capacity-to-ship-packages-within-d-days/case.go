package capacity_to_ship_packages_within_d_days

import (
	"testing"

	"leetcode/tests"
)

type Input struct {
	Weights []int
	Days    int
}

var cases = func() []tests.Case[Input, int] {
	return []tests.Case[Input, int]{
		{Input: Input{Weights: []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, Days: 5}, Except: 15},
		{Input: Input{Weights: []int{3, 2, 2, 4, 1, 4}, Days: 3}, Except: 6},
		{Input: Input{Weights: []int{1, 2, 3, 1, 1}, Days: 4}, Except: 3},
	}
}

type Func func(weights []int, days int) int

var funcs = tests.NewFuncWithAdaptor(func(f Func) func(in Input) (out int) {
	return func(in Input) (out int) {
		return f(in.Weights, in.Days)
	}
})

func Test(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
