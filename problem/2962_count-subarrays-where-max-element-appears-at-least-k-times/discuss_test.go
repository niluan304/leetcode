package main

import (
	"math"
	"math/rand"
	"testing"

	"github.com/niluan304/leetcode/tests"
)

func Test_validateDiscuss(t *testing.T) {
	fs := []func(nums []int, k int) int64{
		discuss,
		discuss2,
		discuss3,
		discuss4,
	}

	for i := 0; i < 1000; i++ {
		nums, k := generate()
		except := bruteForce(nums, k)
		err := tests.Validate2(nums, k, except, fs...)
		if err != nil {
			t.Error(err)
		}
	}
}

func generate() ([]int, int) {
	var n = rand.Intn(1000) + 1 // rand.Intn(n) 随机数组范围为 [0, n-1]
	var k = rand.Intn(5) + 1
	var nums = make([]int, n)
	for j, _ := range nums {
		nums[j] = rand.Intn(50) + 1
	}

	return nums, k
}

func bruteForce(nums []int, k int) int64 {
	var n, ans = len(nums), 0
	for i := 0; i < n; i++ {
		var count, mx = map[int]int{}, math.MinInt
		for j := i; j < n; j++ {
			count[nums[j]]++
			mx = max(mx, nums[j])
			if count[mx] >= k {
				ans++
			}
		}
	}
	return int64(ans)
}
