package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_add_two_numbers_ii(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(l1 *ListNode, l2 *ListNode) *ListNode{
		addTwoNumbers,
		addTwoNumbers2,
		addTwoNumbers3,
		addTwoNumbers4,
		addTwoNumbers5,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[7,2,4,3]
[5,6,4]
[7,8,0,7]

[2,4,3]
[5,6,4]
[8,0,7]

[0]
[0]
[0]

`
