package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_operations_on_tree(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["LockingTree", "lock", "unlock", "unlock", "lock", "upgrade", "lock"]
[[[-1, 0, 0, 1, 1, 2, 2]], [2, 2], [2, 3], [2, 2], [4, 5], [0, 1], [0, 1]]
[null, true, false, true, true, true, false]

`
