package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_mergeKLists(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(lists []*ListNode) *ListNode{
		// bruteForce,
		mergeKLists,
		mergeKLists2,
		// leetcode,
		// endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,4,5],[1,3,4],[2,6]]
[1,1,2,3,4,4,5,6]

[]
[]

[[]]
[]


`
