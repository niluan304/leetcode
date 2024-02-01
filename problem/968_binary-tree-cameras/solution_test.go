package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minCameraCover,
		minCameraCover2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,0,null,0,0]
1

[0,0,null,0,null,0,null,null,0]
2

[0,1,null,2,null,3,null,null,4]
2

[0,1,2,3,4]
2

[0,1,2,3,4,5,6,7,8]
3

`
