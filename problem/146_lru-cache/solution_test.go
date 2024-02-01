package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_lru_cache(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		Constructor,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["LRUCache", "put", "put", "get", "put", "get", "put", "get", "get", "get"]
[[2], [1, 1], [2, 2], [1], [3, 3], [2], [4, 4], [1], [3], [4]]
[null, null, null, 1, null, -1, null, -1, 3, 4]

`
