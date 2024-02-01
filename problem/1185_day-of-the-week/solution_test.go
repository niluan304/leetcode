package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_main(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(day int, month int, year int) string{
		dayOfTheWeek,
		dayOfTheWeek2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
31
8
2019
"Saturday"

18
7
1999
"Sunday"

15
8
1993
"Sunday"

`
