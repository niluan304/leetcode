package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_countOfPairs(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(array []string) []string{
		findLongestSubarray,
		findLongestSubarray2,
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
["A","1","B","C","D","2","3","4","E","5","F","G","6","7","H","I","J","K","L","M"]
["A","1","B","C","D","2","3","4","E","5","F","G","6","7"]

["A","A"]
[]



`
