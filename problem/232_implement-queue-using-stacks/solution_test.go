package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_Constructor(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		// bruteForce,
		Constructor,
		Constructor2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["MyQueue", "push", "push", "peek", "pop", "empty"]
[[], [1], [2], [], [], []]
[null, null, null, 1, 1, false]

["MyQueue","push","push","push","push","pop","push","pop","pop","pop","pop"]
[[],[1],[2],[3],[4],[],[5],[],[],[],[]]
[null,null,null,null,null,1,null,2,3,4,5]


`
