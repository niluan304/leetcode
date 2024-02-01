package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_minimum_time_to_repair_cars(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		repairCars,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[4,2,3,1]
10
16

[5,1,8]
6
16

`
