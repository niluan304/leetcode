package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_xor_of_two_numbers_in_an_array(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findMaximumXOR,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[3,10,5,25,2,8]
28

[14,70,53,83,49,91,36,80,92,51,66,70]
127

`
