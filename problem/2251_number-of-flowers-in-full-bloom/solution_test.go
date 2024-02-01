package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_number_of_flowers_in_full_bloom(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		fullBloomFlowers,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,6],[3,7],[9,12],[4,13]]
[2,3,7,11]
[1,2,2,2]

[[1,10],[3,3]]
[3,3,2]
[2,2,1]

`
