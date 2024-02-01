package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_product_of_array_except_self(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		productExceptSelf,
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
[24,12,8,6]

[-1,1,0,-3,3]
[0,0,9,0,0]

`
