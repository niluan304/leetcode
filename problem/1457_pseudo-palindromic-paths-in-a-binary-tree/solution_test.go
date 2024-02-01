package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		pseudoPalindromicPaths,
		pseudoPalindromicPaths2,
		pseudoPalindromicPaths3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,1,3,1,null,1]
2

[2,1,1,1,3,null,null,null,null,null,1]
1

[9]
1

[2,3,1,3,1,1,1]
3

[]
0
`
