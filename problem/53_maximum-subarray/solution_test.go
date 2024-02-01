package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxSubArray,
		maxSubArray2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[-2,1,-3,4,-1,2,1,-5,4]
6

[1]
1

[5,4,-1,7,8]
23

[-3,-2,-1]
-1

[-1,-2,-3]
-1

[-2,-3,-1]
-1

`
