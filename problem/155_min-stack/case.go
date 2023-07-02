package min_stack

import (
	"reflect"
	"runtime"
	"testing"
)

/**
* Your MinStack object will be instantiated and called as such:
* obj := Constructor();
* obj.Push(val);
* obj.Pop();
* param_3 := obj.Top();
* param_4 := obj.GetMin();
 */

type Stack interface {
	Push(val int)
	Pop()
	Top() int
	GetMin() int
}

var cases = []func(stack Stack){
	Case01,
	Case28,
}

var stackFunc []func() Stack

func once(caseFunc func() Stack) {
	for _, c := range cases {
		c(caseFunc())
	}
}

func AddCases(_cases ...func(stack Stack)) {
	cases = append(cases, _cases...)
}

func AddFunc(f ...func() Stack) {
	stackFunc = append(stackFunc, f...)
}

func Unit(t *testing.T) {
	for _, stack := range stackFunc {
		once(stack)
	}
}

func Bench(b *testing.B) {
	for _, stack := range stackFunc {
		name := runtime.FuncForPC(reflect.ValueOf(stack).Pointer()).Name()
		b.Run(name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				once(stack)
			}
		})
	}
}

func Case01(stack Stack) {
	stack.Push(-2)
	stack.Push(0)
	stack.Push(-3)

	if stack.GetMin() != -3 {
		panic(stack)
	}

	stack.Pop()
	if stack.Top() != 0 {
		panic(stack)
	}
	if stack.GetMin() != -2 {
		panic(stack)
	}

}

func Case28(stack Stack) {
	stack.Push(2147483646)
	stack.Push(2147483646)
	stack.Push(2147483647)

	if stack.Top() != 2147483647 {
		panic(stack)
	}
	stack.Pop()
	if stack.GetMin() != 2147483646 {
		panic(stack)
	}

	stack.Pop()
	if stack.GetMin() != 2147483646 {
		panic(stack)
	}

	stack.Pop()
	stack.Push(2147483647)
	if stack.Top() != 2147483647 {
		panic(stack)
	}
	if stack.GetMin() != 2147483647 {
		panic(stack)
	}

	stack.Push(-2147483648)
	if stack.Top() != -2147483648 {
		panic(stack)
	}
	if stack.GetMin() != -2147483648 {
		panic(stack)
	}

	stack.Pop()
	if stack.GetMin() != 2147483647 {
		panic(stack)
	}
}
