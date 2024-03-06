package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_linked_list_random_node(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
		Constructor2,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["Solution", "getRandom", "getRandom", "getRandom", "getRandom", "getRandom"]
[[[1, 2, 3]], [], [], [], [], []]
[null, 1, 3, 2, 2, 3]

`
