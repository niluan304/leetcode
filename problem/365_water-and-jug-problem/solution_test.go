package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_canMeasureWater(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(jug1 int, jug2 int, target int) bool{
		//bruteForce,
		canMeasureWater,
		//canMeasureWater2,
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
5
4
true

2
6
5
false

1
2
3
true

676766
364766
216376
true

`
