package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maxNumberOfAlloys(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, k int, budget int, composition [][]int, stock []int, cost []int) int{
		//bruteForce,
		maxNumberOfAlloys,
		//maxNumberOfAlloys2,
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
2
15
[[1,1,1],[1,1,10]]
[0,0,0]
[1,2,3]
2

3
2
15
[[1,1,1],[1,1,10]]
[0,0,100]
[1,2,3]
5

2
3
10
[[2,1],[1,2],[1,1]]
[1,1]
[5,5]
2

1
1
100000000
[[1]]
[100000000]
[1]
200000000


`
