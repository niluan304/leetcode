package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, rides [][]int) int64{
		maxTaxiEarnings,
		maxTaxiEarnings2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
5
[[2,5,4],[1,5,1]]
7

20
[[1,6,1],[3,10,2],[10,12,3],[11,12,2],[12,15,2],[13,18,1]]
20

10000
[[1,10000,1],[100,1000,100]]
10000

100
[[1,100,999],[1,50,500],[50,100,500]]
1099

5
[[2,5,4],[2,5,1]]
7

5
[[2,5,1],[2,5,4]]
7
`
