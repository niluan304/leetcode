package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_corporate_flight_bookings(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		corpFlightBookings,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[[1,2,10],[2,3,20],[2,5,25]]
5
[10,55,45,25,25]

[[1,2,10],[2,2,15]]
2
[10,25]

`
