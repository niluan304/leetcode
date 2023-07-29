package contains_duplicate_iii

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	nums      []int
	indexDiff int
	valueDiff int
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				nums:      []int{1, 2, 3, 1},
				indexDiff: 3,
				valueDiff: 0,
			},
			Except: true,
		},
		{
			Input: Input{
				nums:      []int{1, 5, 9, 1, 5, 9},
				indexDiff: 2,
				valueDiff: 3,
			},
			Except: false,
		},
	}
}

type Func func(nums []int, indexDiff int, valueDiff int) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.nums,
			in.indexDiff,
			in.valueDiff,
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
