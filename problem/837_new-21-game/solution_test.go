package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_new_21_game(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		new21Game,
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
1
10
1.00000

6
1
10
0.60000

21
17
10
0.73278

`
