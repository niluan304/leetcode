package largest_number

import (
	"testing"

	"leetcode/tests"
)

var cases = func() []tests.Case[[]int, string] {
	return []tests.Case[[]int, string]{
		{Input: []int{10, 2}, Except: "210"},
		{Input: []int{10, 2, 2}, Except: "2210"},
		{Input: []int{10, 2, 12, 1}, Except: "212110"},
		{Input: []int{10001, 2, 12, 1}, Except: "212110001"},
		{Input: []int{3, 30, 34, 5, 9}, Except: "9534330"},
		{Input: []int{3, 30, 34, 5, 9, 92}, Except: "992534330"},
		{Input: []int{3, 30, 32, 5, 9, 92}, Except: "992533230"},
		{Input: []int{3, 30, 32, 5, 9, 92, 0}, Except: "9925332300"},
		{Input: []int{34323, 3432}, Except: "343234323"},
		{Input: []int{34323, 3432, 34}, Except: "34343234323"},
		{Input: []int{34323, 3432, 343}, Except: "343343234323"},
		{Input: []int{0, 0}, Except: "0"},
		{Input: []int{0, 0, 0}, Except: "0"},
		{Input: []int{111311, 1113}, Except: "1113111311"},
	}
}

var funcs = tests.NewFunc[[]int, string]()

func AddCases(c func() []tests.Case[[]int, string]) {
	_cases := cases()
	cases = func() []tests.Case[[]int, string] {
		return append(_cases, c()...)
	}
}

func AddFunc(f ...func([]int) string) {
	funcs = append(funcs, tests.NewFunc(f...)...)
}

func Unit(t *testing.T) {
	tests.Unit(t, cases, funcs...)
}

var checkResult = true

func Bench(b *testing.B) {
	checkResult = false
	tests.Bench(b, cases, funcs...)
}
