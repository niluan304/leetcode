package main

import (
	"testing"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(head *ListNode) *ListNode{
		insertGreatestCommonDivisors,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[18,6,10,3]
[18,6,6,2,10,1,3]

[7]
[7]


`
