package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int, x int) []int{
		getSubarrayBeauty,
		getSubarrayBeauty2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,-1,-3,-2,3]
3
2
[-1,-2,-2]

[-1,-2,-3,-4,-5]
2
2
[-1,-2,-3,-4]

[-3,1,2,-3,0,-3]
2
1
[-3,0,-3,-3,-3]

[-2,-1,-3,-2,3]
3
2
[-2,-2,-2]

[-1,-2,-3,-4,-5,-6,-7]
3
2
[-2,-3,-4,-5,-6]

[-1,-2,-3,-7,-5,-3,-1]
3
2
[-2,-3,-5,-5,-3]


`
