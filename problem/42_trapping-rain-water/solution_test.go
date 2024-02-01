package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_trapping_rain_water(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		trap,
		trap2,
		trap3,
		trap4,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[0,1,0,2,1,0,1,3,2,1,2,1]
6

[4,2,0,3,2,5]
9

`
