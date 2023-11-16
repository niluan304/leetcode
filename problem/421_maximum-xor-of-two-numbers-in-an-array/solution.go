package main

import (
	"math/bits"
	"slices"
)

func findMaximumXOR(nums []int) (ans int) {
	n := bits.Len(uint(slices.Max(nums)))
	for i := n - 1; i >= 0; i-- {
		ans = ans<<1 + 1
		if !findTwoSum(nums, ans, i) {
			ans -= 1
		}
	}
	return ans
}

func findTwoSum(nums []int, target int, offset int) bool {
	var cache = map[int]bool{}
	for _, num := range nums {
		x := num >> offset
		y := x ^ target
		if cache[y] {
			return true
		}
		cache[x] = true
	}
	return false
}

func findMaximumXOR2(nums []int) (ans int) {
	highBit := bits.Len(uint(slices.Max(nums))) - 1
	seen := map[int]bool{}
	mask := 0
	for i := highBit; i >= 0; i-- { // 从最高位开始枚举
		clear(seen)
		mask |= 1 << i
		newAns := ans | 1<<i // 这个比特位可以是 1 吗？
		for _, x := range nums {
			x &= mask // 低于 i 的比特位置为 0
			if seen[newAns^x] {
				ans = newAns // 这个比特位可以是 1
				break
			}
			seen[x] = true
		}
	}
	return
}

func findMaximumXOR3(nums []int) (ans int) {
	n := bits.Len(uint(slices.Max(nums)))
	var cache = map[int]bool{}

Loop:
	for i := n - 1; i >= 0; i-- {
		ans = ans<<1 + 1
		clear(cache)
		for _, num := range nums {
			x := num >> i
			y := x ^ ans
			if cache[y] {
				continue Loop
			}
			cache[x] = true
		}
		ans -= 1
	}
	return ans
}
