package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_serialize_and_deserialize_bst(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,1,3]
[2,1,3]

[]
[]

`
