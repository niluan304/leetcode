package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_partition(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(num string) bool{
		//bruteForce,
		isAdditiveNumber,
		isAdditiveNumber2,
		isAdditiveNumber3,
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
"112358"
true

"199100199"
true

"123"
true

"124"
false

"11235812"
false

"1123581321"
true

"112233"
true

"1122"
false

"199111992"
true

"199001200"
false

`
