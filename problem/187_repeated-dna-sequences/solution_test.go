package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_repeated_dna_sequences(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findRepeatedDnaSequences,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"AAAAACCCCCAAAAACCCCCCAAAAAGGGTTT"
["AAAAACCCCC","CCCCCAAAAA"]

"AAAAAAAAAAAAA"
["AAAAAAAAAA"]

`
