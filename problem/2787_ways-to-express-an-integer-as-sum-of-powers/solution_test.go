package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_numberOfWays(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, x int) int{
		//bruteForce,
		numberOfWays,
		numberOfWays2,
		numberOfWays3,
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
10
2
1

4
1
2

2
1
1

5
1
3


`

func Test_executionTime(t *testing.T) {
	fs := []func(n int, x int) int{
		//bruteForce,
		numberOfWays,
		numberOfWays2,
		numberOfWays3,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		t.Run(tests.FuncName(f), func(t *testing.T) {
			for n := 1; n <= 300; n++ {
				for x := 1; x <= 5; x++ {
					f(n, x)
				}
			}
		})
	}
}
