package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_split_array_into_fibonacci_sequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		splitIntoFibonacci,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
"1101111"
[11,0,11,11]

"112358130"
[]

"0123"
[]

"74912134825162255812723932620170946950766784234934"
[]

"121202436"
[]

"1212002436"
[]

"00123"
[]`
