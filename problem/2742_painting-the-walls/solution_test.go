package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_paintWalls(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(cost []int, time []int) int{
		//bruteForce,
		paintWalls,
		paintWalls2,
		paintWalls3,
		paintWalls4,
		paintWalls5,
		//leetcode,
		//endlessCheng,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[1,2,3,2]
[1,2,3,2]
3

[2,3,4,2]
[1,1,1,1]
4

[1,99,99,99,99]
[99,1,1,1,1]
1

[43,40,23,59,70,79,51,22,63]
[01,01,01,02,03,01,03,01,01]
132

[29,232,109,14,69,171,15,305,175,223,100,270,109,77,294,29,77,271]
[03,002,002,01,01,001,02,002,004,001,004,001,004,01,001,01,04,002]
233

`
