package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(tomatoSlices int, cheeseSlices int) []int{
		numOfBurgers,
		numOfBurgers2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
16
7
[1,6]

17
4
[]

4
17
[]

0
0
[0,0]

2
1
[0,1]

2385088
164763
[]

3962
1205
[776,429]

2510308
531466
[]
`
