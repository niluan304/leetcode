package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_capacity_to_ship_packages_within_d_days(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		shipWithinDays,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,4,5,6,7,8,9,10]
5
15

[3,2,2,4,1,4]
3
6

[1,2,3,1,1]
4
3

`
