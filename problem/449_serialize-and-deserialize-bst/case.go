package serialize_and_deserialize_bst

import (
	"reflect"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type TreeNode = structs.TreeNode

type CodecDriver interface {
	serialize(root *TreeNode) string
	deserialize(data string) *TreeNode
}

func ToPtr[T any](x T) *T {
	return &x
}

type Input struct {
	root *structs.TreeNode
}

type Output *structs.TreeNode

var cases = func() []tests.Case[Input, Output] {
	return []tests.Case[Input, Output]{
		// TODO add question case
	}
}

type Func func(root *structs.TreeNode) *structs.TreeNode

var funcs = tests.NewFuncWithAdaptor(adaptor)

func adaptor(f Func) func(in Input) Output {
	return func(in Input) Output {
		return Output(f(
			in.root,
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
