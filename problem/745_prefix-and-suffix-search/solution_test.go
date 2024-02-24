package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_Constructor(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		//bruteForce,
		Constructor,
		Constructor2,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunClass(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["WordFilter", "f"]
[[["apple"]], ["a", "e"]]
[null, 0]

["WordFilter", "f"]
[[["apple"]], ["apple", "e"]]
[null, 0]


["WordFilter", "f"]
[[["apple","applepie"]], ["apple", "e"]]
[null, 1]

["WordFilter", "f", "f", "f"]
[[["applepie","apple"]], ["apple", "e"], ["","applepie"], ["b","e"]]
[null, 1, 0, -1]


`
