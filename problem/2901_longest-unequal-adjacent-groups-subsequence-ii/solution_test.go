package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_getWordsInLongestSubsequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(n int, words []string, groups []int) []string{
		//bruteForce,
		getWordsInLongestSubsequence,
		getWordsInLongestSubsequence2,
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
["bab","dab","cab"]
[1,2,2]
["bab","cab"]

4
["a","b","c","d"]
[1,2,3,4]
["a","b","c","d"]

3
["bdb","aaa","ada"]
[2,1,3]
["aaa","ada"]
`
