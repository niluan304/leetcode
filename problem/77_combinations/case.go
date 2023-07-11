package combinations

import (
	"reflect"
	"sort"
	"testing"

	"leetcode/tests"
)

type Input struct {
	n int
	k int
}

type Output [][]int

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		{
			Input:  Input{n: 4, k: 2},
			Except: [][]int{{2, 4}, {3, 4}, {2, 3}, {1, 2}, {1, 3}, {1, 4}},
		},
		{
			Input:  Input{n: 1, k: 1},
			Except: [][]int{{1}},
		},
	}
}

func Equal(except Output, output Output) bool {
	if len(except) != len(output) {
		return false
	}

	sortOutput := func(v Output) {
		sort.Slice(v, func(i, j int) bool {
			if len(v[i]) != len(v[j]) {
				return len(v[i]) < len(v[j])
			}

			for k := 0; k < len(v[i]); k++ {
				if v[i][k] != v[j][k] {
					return v[i][k] < v[j][k]
				}
			}

			return false
		})
	}

	sortOutput(except)
	sortOutput(output)

	return reflect.DeepEqual(except, output)
}

type Func func(n int, k int) [][]int

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.n,
			in.k,
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

func Unit(t *testing.T) {
	tests.Unit(t, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, tests.Test[Input, Output]{
		Solution: funcs,
		Cases:    cases,
		IsEqual:  nil,
	})
}
