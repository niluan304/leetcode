package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reconstruct_itinerary(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		findItinerary,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[["MUC","LHR"],["JFK","MUC"],["SFO","SJC"],["LHR","SFO"]]
["JFK","MUC","LHR","SFO","SJC"]

[["JFK","SFO"],["JFK","ATL"],["SFO","ATL"],["ATL","JFK"],["ATL","SFO"]]
["JFK","ATL","JFK","SFO","ATL","SFO"]

`
