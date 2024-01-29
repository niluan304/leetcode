package main

import (
	"math/rand"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_longest_consecutive_sequence(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int) int{
		sortInt,
		longestConsecutive,
	}

	for _, f := range fs {
		err := tests.RunFunc(t, f, samples, targetCaseNum)
		if err != nil {
			t.Error(err)
		}
	}
}

var samples = `
[100,4,200,1,3,2]
4

[0,3,7,2,5,8,4,6,0,1]
9



`

func Test_validate(t *testing.T) {
	for i := 0; i < 100000; i++ {
		nums := generate()
		except := sortInt(nums)
		err := tests.Validate(nums, except, longestConsecutive)
		if err != nil {
			t.Error(err)
		}
	}
}

func generate() []int {
	var n = rand.Intn(1000) + 1
	var nums = make([]int, 1000)
	for i, _ := range nums {
		nums[i] = rand.Intn(n)
	}
	return nums
}
