package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countOfPairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, x int, y int) []int{
		countOfPairs,
		countOfPairs2,
		countOfPairs3,
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
3
1
3
[6,0,0]

5
2
4
[10,8,2,0,0]

4
1
1
[6,4,2,0]

9
2
6
[18,22,16,10,6,0,0,0,0]


`
