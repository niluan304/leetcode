package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) bool{
		find132pattern,
		find132pattern2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4]
false

[3,1,4,2]
true

[-1,3,2,0]
true

[1,4,2]
true

[1,4]
false

[4]
false

[3,5,0,3,4]
true

[1,0,1,-4,-3]
false

[5,100,1,3,2]
true

[-2,1,1]
false`
