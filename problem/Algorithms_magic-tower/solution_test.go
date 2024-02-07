package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_magicTower(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		//bruteForce,
		magicTower,
		//magicTower2,
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
[100,100,100,-250,-60,-140,-50,-50,100,150]
1

[-200,-300,400,0]
-1

[-1,-1,10]
2

`
