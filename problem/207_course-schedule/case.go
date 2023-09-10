package course_schedule

import (
	"reflect"
	"testing"

	"leetcode/tests"
)

type Input struct {
	numCourses    int
	prerequisites [][]int
}

type Output bool

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input: Input{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
				},
			},
			Except: true,
		},
		{
			Input: Input{
				numCourses: 2,
				prerequisites: [][]int{
					{1, 0},
					{0, 1},
				},
			},
			Except: false,
		},
	}
}

type Func func(numCourses int, prerequisites [][]int) bool

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.numCourses,
			in.prerequisites,
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
