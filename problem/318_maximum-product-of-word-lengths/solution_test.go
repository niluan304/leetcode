package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_maximum_product_of_word_lengths(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		maxProduct,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["abcw","baz","foo","bar","xtfn","abcdef"]
16

["a","ab","abc","d","cd","bcd","abcd"]
4

["a","aa","aaa","aaaa"]
0

`
