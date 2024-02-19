package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_preorder(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(root *Node) []int{
		//bruteForce,
		preorder,
		//preorder2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,null,3,2,4,null,5,6]
[1,3,5,6,2,4]

[1,null,2,3,4,5,null,null,6,7,null,8,null,9,10,null,null,11,null,12,null,13,null,null,14]
[1,2,3,6,7,11,14,4,8,12,5,9,13,10]


`
