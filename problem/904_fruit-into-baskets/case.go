package fruit_into_baskets

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	fruits []int
}

type Output int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{fruits: []int{1, 2, 1}},
			Except: 3,
		},
		{
			Input:  Input{fruits: []int{0, 1, 2, 2}},
			Except: 3,
		},
		{
			Input:  Input{fruits: []int{1, 2, 3, 2, 2}},
			Except: 4,
		},
		{
			Input:  Input{fruits: []int{3, 3, 3, 1, 2, 1, 1, 2, 3, 3, 4}},
			Except: 5,
		},
	}
}

type Func func(fruits []int) int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.fruits,
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
