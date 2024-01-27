package main

import (
	"math/rand"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_subarraySum(t *testing.T) {
	targetCaseNum := 0 // -1

	fs := []func(nums []int, k int) int{
		bruteForce,
		subarraySum,
		//subarraySum2,
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
[1,1,1]
2
2

[1,2,3]
3
2


`

func Test_validate(t *testing.T) {
	for i := 0; i < 10000; i++ {
		nums, k := generate()
		except := bruteForce(nums, k)
		err := tests.Validate2(nums, k, except, subarraySum)
		if err != nil {
			t.Error(err)
		}
	}
}

func generate() (nums []int, k int) {
	nums = make([]int, rand.Intn(1000)+1)
	for i, _ := range nums {
		nums[i] = rand.Intn(2000) - 1000
	}
	return nums, rand.Intn(2e7) - 1e7
}

func bruteForce(nums []int, k int) (ans int) {
	for i, _ := range nums {
		sum := 0
		for j := i; j >= 0; j-- {
			sum += nums[j]
			if sum == k {
				ans++
			}
		}
	}
	return ans
}
