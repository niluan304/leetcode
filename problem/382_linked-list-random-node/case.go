package linked_list_random_node

import (
	"encoding/json"
	"fmt"
	"testing"

	"leetcode/structs"
	"leetcode/tests"
)

type ListNode = structs.ListNode

type Randomizer interface {
	GetRandom() int
}

type Func func(head *ListNode) Randomizer

func unit(f Func) {
	const num, n = 100, 100

	var nodes = make([]int, num)
	for i := range nodes {
		nodes[i] = i
	}

	var count = make(map[int]int, num)
	r := f(structs.NewListNode(nodes))
	for i := 0; i < n; i++ {
		node := r.GetRandom()
		count[node]++
	}

	data, err := json.MarshalIndent(count, "", "")
	data, err = json.Marshal(count)
	if err != nil {
		panic(err)
	}
	fmt.Printf("len:%d,\t data: %+v\n", len(count), string(data))
}

var funcs []Func

func AddFuncs(f ...Func) {
	funcs = append(funcs, f...)
}

func Unit(t *testing.T) {
	for _, f := range funcs {
		t.Run(tests.FuncName(f), func(t *testing.T) {
			unit(f)
		})
	}
}

func Bench(b *testing.B) {
	for _, f := range funcs {
		b.Run(tests.FuncName(f), func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				unit(f)
			}
		})
	}
}
