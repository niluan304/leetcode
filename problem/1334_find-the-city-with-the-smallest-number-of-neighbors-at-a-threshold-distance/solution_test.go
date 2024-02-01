package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_find_the_city_with_the_smallest_number_of_neighbors_at_a_threshold_distance(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findTheCity,
		findTheCity2,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
4
[[0,1,3],[1,2,1],[1,3,4],[2,3,1]]
4
3

5
[[0,1,2],[0,4,8],[1,2,3],[1,4,2],[2,3,1],[3,4,1]]
2
0

`
