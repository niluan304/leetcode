package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		dailyTemperatures,
		dailyTemperatures2,
		dailyTemperatures3,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[73,74,75,71,69,72,76,73]
[1,1,4,2,1,1,0,0]

[30,40,50,60]
[1,1,1,0]

[30,60,90]
[1,1,0]

[89,62,70,58,47,47,46,76,100,70]
[8,1,5,4,3,2,1,1,0,0]
`
