package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_linked_list_cycle(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		hasCycle,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,2,0,-4]
1
true

[1,2]
0
true

[1]
-1
false

`
