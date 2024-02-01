package main

import (
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_reward_top_k_students(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []interface{}{
		topStudents,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
["smart","brilliant","studious"]
["not"]
["this student is studious","the student is smart"]
[1,2]
2
[1,2]

["smart","brilliant","studious"]
["not"]
["this student is not studious","the student is smart"]
[1,2]
2
[2,1]

`
