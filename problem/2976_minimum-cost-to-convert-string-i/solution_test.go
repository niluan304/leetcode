package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimumCost(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(source string, target string, original []byte, changed []byte, cost []int) int64{
		//bruteForce,
		minimumCost,
		//minimumCost2,
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
"abcd"
"acbe"
["a","b","c","c","e","d"]
["b","c","b","e","b","e"]
[2,5,5,1,2,20]
28

"aaaa"
"bbbb"
["a","c"]
["c","b"]
[1,2]
12

"abcd"
"abce"
["a"]
["e"]
[10000]
-1


`
