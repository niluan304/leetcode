package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_construct_product_matrix(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		mostSavedTree,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
2
[[1,2]]
0

4
[[1,2],[2,3],[2,4]]
2

7
[[1,2],[1,5],[2,3],[2,4],[5,6],[5,7]]
2

15
[[1,2],[2,3],[3,4],[4,5],[4,6],[3,7],[2,8],[1,9],[9,10],[9,11],[10,12],[10,13],[11,14],[11,15]]
10

15
[[1,2],[2,3],[3,4],[4,5],[4,6],[5,7],[7,8],[1,9],[15,10],[9,11],[10,12],[10,13],[11,14],[14,15]]
11
`
