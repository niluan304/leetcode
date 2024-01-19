package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minCost(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, x int) int64{
		minCost,
		minCost2,
		minCost3,
		minCost4,
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
[20,1,15]
5
13

[1,2,3]
4
6

[20,1,15,1,2,3,24,123,4,2,23,23,3]
5
36

[1,2,3,312,34,123,23,34,1,1]
4
33

[1,2,3,4,3,1,23,4,1,2,3,4,1,2,3,1]
2
23

[2,1,3,4,5,6,1,7,8,2,1,8,4,3,2,1,3,2]
4
33

[2,1,3,4,5,6,1,7,8,2,1,8,4,3,2,1,3,2]
9
46

[2,1,3,4,5,6,1,7,8,2,1,8,4,3,2,1,3,2]
3
30

[2,1,3,4,5,6,1,7,8,2,1,8,4,3,2,1,3,2]
6
39

`
