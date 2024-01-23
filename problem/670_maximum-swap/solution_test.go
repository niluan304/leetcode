package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximumSwap(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(num int) int{
		maximumSwap,
		maximumSwap2,
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
2736
7236

9973
9973

3274893279
9274893273

832469392394349
932469392394348

9912373
9972313

`
