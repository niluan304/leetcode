package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_score_triangulation_of_polygon(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		minScoreTriangulation,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3]
6

[3,7,4,5]
144

[1,3,1,4,1,5]
13

`
