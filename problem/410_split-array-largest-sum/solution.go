package main

import (
	"sort"
)

func splitArray(nums []int, k int) int {
	sum, mx := 0, 0
	for _, x := range nums {
		sum += x
		mx = max(mx, x)
	}
	left := max(mx, (sum-1)/k+1)
	right := sum

	// 必须是左闭右开区间
	ans := left + sort.Search(right-left, func(mx int) bool {
		mx += left
		cnt, s := 1, 0
		for _, x := range nums {
			if s+x <= mx {
				s += x
			} else { // 新划分一段
				if cnt == k {
					return false
				}
				cnt += 1
				s = x
			}
		}
		return true
	})
	return ans
}

// 二分答案
// 关键字：最大值最小（最小化最大值）
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func splitArray2(nums []int, k int) int {
	right, left := 0, 0
	for _, num := range nums {
		right += num
		left = max(left, num)
	}

	// 为什么需要 minSum+left ?
	// 因为寻找的值必定得 >= max(nums)，不然没有子数组能容纳数组中的最大值
	// 由于库函数的是在区间 [0, n] 内二分，而实际的二分区间为 [max(nums), sum(nums)]
	// 因此需要将二分区间左移 max(nums)
	var ans = left + sort.Search(right-left, func(minSum int) bool {
		count := 1
		sum := 0
		for _, num := range nums {
			sum += num
			if sum > minSum+left {
				count++
				sum = num
			}
		}
		return count <= k
	})

	return ans
}
