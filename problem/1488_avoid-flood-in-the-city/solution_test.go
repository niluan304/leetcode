package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_avoid_flood_in_the_city(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		avoidFlood,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4]
[-1,-1,-1,-1]

[1,2,0,0,2,1]
[-1,-1,2,1,-1,-1]

[1,2,0,1,2]
[]

`
