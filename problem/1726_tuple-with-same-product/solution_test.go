package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_tuple_with_same_product(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		tupleSameProduct,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[2,3,4,6]
8

[1,2,4,5,10]
16

`
