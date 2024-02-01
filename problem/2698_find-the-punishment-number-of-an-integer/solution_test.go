package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_the_punishment_number_of_an_integer(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		punishmentNumber,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
10
182

37
1478

`
