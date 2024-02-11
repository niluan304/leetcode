package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_beautifulSubsets(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		//bruteForce,
		beautifulSubsets,
		beautifulSubsets2,
		beautifulSubsets3,
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
[2,4,6]
2
4

[1]
1
1

[2,4,6,10]
2
9

[2,4,6,10,20]
2
19

[1,1,1,1,1,1,1,1]
2
255

[2,4,6,2]
2
8

[2,4,6,2,4]
2
10

[2,4,6,2,4,6]
2
18

[10,4,5,7,2,1]
3
23

[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]
2
20735

[1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20]
3
24275


`
