package combination_sum_ii

import (
	"reflect"
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	candidates []int
	target     int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				candidates: []int{10, 1, 2, 7, 6, 1, 5},
				target:     8,
			},
			Except: Output{
				{1, 1, 6},
				{1, 2, 5},
				{1, 7},
				{2, 6},
			},
		},
		{
			Input: Input{
				candidates: []int{2, 5, 2, 1, 2},
				target:     5,
			},
			Except: Output{
				{1, 2, 2},
				{5},
			},
		},
	}
}

type Func func(candidates []int, target int) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.candidates,
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

func Sort(x Output) {
	for i := range x {
		sort.Ints(x[i])
	}

	sort.Slice(x, func(i, j int) bool {
		if len(x[i]) != len(x[j]) {
			return len(x[i]) < len(x[j])
		}

		for k := range x[i] {
			if x[i][k] != x[j][k] {
				return x[i][k] < x[j][k]
			}
		}

		return i < j
	})
}

func Equal(x, y Output) bool {
	if len(x) != len(y) {
		return false
	}

	Sort(x)
	Sort(y)

	for i := range x {
		if !reflect.DeepEqual(x[i], y[i]) {
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
