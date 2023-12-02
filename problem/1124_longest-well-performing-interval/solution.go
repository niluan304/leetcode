package main

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func longestWPI(hours []int) int {
	var ans = 0
	var n = len(hours)
	for i, _ := range hours {
		var wpi = 0
		for j := i; j < n; j++ {
			if hours[j] > 8 {
				wpi++
			} else {
				wpi--
			}
			if wpi > 0 {
				ans = max(ans, j-i+1)
			}
		}
	}
	return ans
}

// 前缀和 + 哈希表
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func longestWPI2(hours []int) int {
	var ans = 0
	var n = len(hours)

	var preSum = make([]int, n+1)
	for i, h := range hours {
		preSum[i+1] = preSum[i] - 1
		if h > 8 {
			preSum[i+1] = preSum[i] + 1
		}
	}

	var minIdx = map[int]int{}
	for i := 1; i < n+1; i++ {
		sum := preSum[i]
		if sum > 0 {
			ans = max(ans, i)
			continue
		}

		if j, ok := minIdx[sum-1]; ok {
			ans = max(ans, i-j)
		}
		if _, ok := minIdx[sum]; !ok { // preSum[0] = 0，无法用于设置索引，所以 i 必须从 1 开始
			minIdx[sum] = i
		}
	}
	return ans
}

// 贪心思想：前缀和 + 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func longestWPI3(hours []int) int {
	var n = len(hours)

	var stack = []int{0}
	var preSum = make([]int, n+1)
	for i, h := range hours {
		j := i + 1
		preSum[j] = preSum[i] - 1
		if h > 8 {
			preSum[j] = preSum[i] + 1
		}
		if preSum[j] < preSum[stack[len(stack)-1]] {
			stack = append(stack, j)
		}
	}

	var ans = 0
	for i := n; i > 0; i-- {
		for m := len(stack); m > 0; m-- {
			j := stack[m-1]
			if preSum[i] <= preSum[j] {
				break
			}
			ans = max(ans, i-j)
			stack = stack[:m-1]
		}
	}
	return ans
}
