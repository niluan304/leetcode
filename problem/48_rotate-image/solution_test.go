package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_rotate_image(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		rotate,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,2,3],[4,5,6],[7,8,9]]
[[7,4,1],[8,5,2],[9,6,3]]

[[5,1,9,11],[2,4,8,10],[13,3,6,7],[15,14,12,16]]
[[15,13,2,5],[14,3,4,1],[12,6,8,9],[16,7,10,11]]

`
