package main

import (
	"sort"
)

func singleNumber(nums []int) int {
	var ans = 0
	for i := range nums {
		ans ^= nums[i]
	}
	return ans
}

func singleNumber2(nums []int) int {
	// 先排序, 再查找
	sort.Ints(nums)

	// 长度一定是 2N+1
	for i := 0; i < len(nums)-1; i += 2 {
		if nums[i] != nums[i+1] {
			// 数字必定出现在奇数位置
			return nums[i]
		}
	}

	return nums[len(nums)-1]
}

func leetcode1(nums []int) int {
	for i := 1; i < len(nums); i++ {
		nums[0] ^= nums[i]
	}

	return nums[0]
}
