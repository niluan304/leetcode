package main

import "math"

func minimumSeconds(nums []int) int {
	index := make(map[int][]int)
	n, ans := len(nums), math.MaxInt32

	for i, num := range nums {
		index[num] = append(index[num], i)
	}
	for _, idx := range index {
		mx := idx[0] + n - idx[len(idx)-1]
		for i := 1; i < len(idx); i++ {
			mx = max(mx, idx[i]-idx[i-1])
		}
		ans = min(ans, mx/2)
	}
	return ans
}
