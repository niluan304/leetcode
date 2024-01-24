package main

import (
	"math"
)

const mod = 1_000_000_007

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时
func bruteForce(arr []int) int {
	var ans = 0
	for i, _ := range arr {
		v := math.MaxInt
		for j := i; j >= 0; j-- {
			v = min(v, arr[j])
			ans = (ans + v) % mod
		}
	}
	return ans
}

// 单调栈 + 动态规划
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func sumSubarrayMins(arr []int) int {
	arr = append([]int{0}, arr...) // 1 <= arr[i] <= 3 * 10^4

	var n = len(arr)
	var dp = make([]int, n)
	var stack []int
	var ans = 0
	for i, num := range arr {
		for m := len(stack); m > 0; m-- {
			j := stack[m-1]
			if num > arr[j] {
				// 处理后 arr[0] 一定为最小值
				dp[i] = dp[j] + num*(i-j)%mod
				ans = (ans + dp[i]) % mod
				break
			}
			stack = stack[:m-1]
		}
		stack = append(stack, i)
	}
	return ans
}

// 动态规划 + 单调栈
// dp[i-1] 表示区间 [0, i-1] 内 min(b) 的和
// 有一个新元素 nums[i]，假设 nums[i] > nums[i-1]，
// 那将 nums[i] 拼到 区间 [0 : i-1] 后面，min(b) 是不会变化的，
// dp[i] 可以 **继承** dp[i-1] 的 min(b) 的和，并额外增加 nums[i] 作为 min(b) 时的值
// 再将这个情况推广，会发现，dp[i] 能继承的是上一个小于 nums[i] 的数，记为 nums[j]，
// 那么有递推公式：dp[i] = dp[j] + [i-j+1] * num
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func sumSubarrayMins2(arr []int) int {
	var n = len(arr)
	var dp = make([]int, n+1)
	var stack = []int{-1} // 哨兵
	var ans = 0

	for i, num := range arr {
		m := len(stack)
		for m > 1 && num <= arr[stack[m-1]] { // 单调递减栈
			m--
		}

		// top 是上一个小于 nums[i] 的元素索引，top == -1 哨兵时，表示 nums[i] 在区间 [0, i] 是最小值的
		top := stack[m-1]
		// nums[i] 作为新的最小值入栈，小于于等于 nums[i] 的元素出栈
		stack = append(stack[:m], i)

		// 递推公式： dp[i] = dp[top] + v
		// v = nums[i]*(i - top)
		dp[i+1] = (dp[top+1] + num*(i-top)%mod) % mod

		ans = (ans + dp[i+1]) % mod
	}
	return ans
}

// 单调栈 + 贡献值
// 时间复杂度: O(n)
// 空间复杂度: O(n)
// 贡献值的本质是扩散：
// 也就是 [84. 柱状图中最大的矩形](https://leetcode.cn/problems/largest-rectangle-in-histogram/) 变形题
// 依次遍历柱形的高度，对于每一个高度分别向两边扩散，求出以当前高度为矩形的最大宽度多少。
func sumSubarrayMins3(arr []int) int {
	var n = len(arr)
	var right, left = make([]int, n+1), make([]int, n+1)

	// 1 <= arr[i] <= 3 * 10^4
	// 因此可以使用 -1 作为哨兵，以判断边界
	arr = append(arr, -1) // 新增 -1 作为边界哨兵
	var stack = []int{-1}
	for i, v := range arr {
		// 在方法一中，我们在对位置 i 进行入栈操作时，确定了它的左边界。
		// 从直觉上来说，与之对应的我们在对位置 i 进行出栈操作时可以确定它的右边界！

		m := len(stack)
		for m > 1 && v <= arr[stack[m-1]] {
			right[stack[m-1]] = i
			m--
		}
		left[i] = stack[m-1]
		stack = append(stack[:m], i)
	}

	var ans = 0
	for i := 0; i < n; i++ { // arr[n] = -1 为 哨兵，不应该遍历
		l := i - left[i]
		r := right[i] - i
		ans = (ans + arr[i]*l*r) % mod
	}
	return ans
}

// sumSubarrayMins3 的空间优化版本
func sumSubarrayMins4(arr []int) int {
	var n = len(arr)
	var left = make([]int, n+1)

	// 1 <= arr[i] <= 3 * 10^4
	// 因此可以使用 -1 作为哨兵，以判断边界
	arr = append(arr, -1) // 新增 -1 作为边界哨兵
	var stack = []int{-1}
	var ans = 0
	for i, v := range arr {
		// 在方法一中，我们在对位置 i 进行入栈操作时，确定了它的左边界。
		// 从直觉上来说，与之对应的我们在对位置 i 进行出栈操作时可以确定它的右边界！

		m := len(stack)
		for m > 1 && v <= arr[stack[m-1]] {
			j := stack[m-1]
			l, r := j-left[j], i-j
			ans = (ans + arr[j]*l*r) % mod
			m--
		}
		left[i] = stack[m-1]
		stack = append(stack[:m], i)
	}
	return ans
}
