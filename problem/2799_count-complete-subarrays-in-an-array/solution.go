package main

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(n)
func countCompleteSubarrays(nums []int) int {
	var set = map[int]struct{}{}
	for _, num := range nums {
		set[num] = struct{}{}
	}

	var ans, n, sum = 0, len(nums), len(set)

	// 可以使用滑动窗口优化
	for i := 0; i < n; i++ {
		clear(set)
		for j := i; j < n; j++ {
			set[nums[j]] = struct{}{}
			if len(set) == sum {
				ans += n - j
				break
			}
		}
	}
	return ans
}

func countCompleteSubarrays2(nums []int) int {
	var count = map[int]int{}
	for _, num := range nums {
		count[num] = 0
	}

	var ans, n, sum = 0, len(nums), len(count)
	clear(count)

	var left = 0
	for right, num := range nums {
		count[num]++
		for len(count) >= sum {
			key := nums[left]
			count[key]--
			if count[key] == 0 {
				delete(count, key)
			}
			left++

			ans += n - right
		}
	}
	return ans
}
