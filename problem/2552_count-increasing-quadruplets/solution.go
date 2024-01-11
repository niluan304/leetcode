package main

import "sort"

// - 时间复杂度：$\mathcal{O}(n^2\log n)$。
// - 空间复杂度：$\mathcal{O}(n^2)$。
func countQuadruplets(nums []int) int64 {
	var ans = 0
	var n = len(nums)
	var less, great = make([][]int, n), make([][]int, n)

	for j := n - 1; j >= 0; j-- {
		great[j] = make([]int, 0, n)
		for l := j + 1; l < n; l++ {
			if nums[j] < nums[l] {
				great[j] = append(great[j], l) // great[j] 记录在 j 处所有满足条件的索引 l，即 对于 l ∈ great[j][:]，都有 nums[j] < nums[l]
			}
		}
	}
	for k := 0; k < n; k++ {
		less[k] = make([]int, 0, n)
		for i := 0; i < k; i++ {
			if nums[k] > nums[i] {
				less[k] = append(less[k], i) //  less[k] 记录在 k 处所有满足条件的索引 i，即 对于 i ∈ less[k][:]，都有 nums[i] < nums[k]
			}
		}
	}

	for j := 0; j < n; j++ {
		for k := j + 1; k < n; k++ {
			if nums[j] > nums[k] { // 枚举所有满足条件的 j, k：nums[j] > nums[k]
				x := sort.SearchInts(less[k], j)  // 在  less[k][:] 中寻找索引小于 j 的位置，即在 j 的左边有多少个 i 满足 nums[i] < nums[k]
				y := sort.SearchInts(great[j], k) // 在 great[j][:] 中寻找索引大于 k 的位置，即在 k 的右边有多少个 l 满足 nums[l] > nums[j]
				ans += x * (len(great[j]) - y)
			}
		}
	}

	return int64(ans)
}

// ### 提示 1
//
// 枚举 $j$ 和 $k$ 这两个**中间**的，会更容易计算。
//
// 这个技巧在去年的周赛题 [2242. 节点序列的最大得分](https://leetcode.cn/problems/maximum-score-of-a-node-sequence/) 出现过。
//
// 需要计算哪些信息？
//
// ### 提示 2
//
// 需要计算：
//
// - 在 $k$ 右侧的比 $\textit{nums}[j]$ 大的元素个数，记作 $\textit{great}[k][\textit{nums}[j]]$；
// - 在 $j$ 左侧的比 $\textit{nums}[k]$ 小的元素个数，记作 $\textit{less}[j][\textit{nums}[k]]$。
//
// 对于固定的 $j$ 和 $k$，根据乘法原理，对答案的贡献为
//
// $$
// \textit{less}[j][\textit{nums}[k]]\cdot \textit{great}[k][\textit{nums}[j]]
// $$
//
// 如何维护个数？
//
// ### 提示 3
//
// 维护方法如下：
//
// - 倒序遍历 $\textit{nums}$，设 $x < \textit{nums}[k+1]$，对于 $x$，大于它的数的个数加一，即 $\textit{great}[k][x]$ 加一；
// - 正序遍历 $\textit{nums}$，设 $x > \textit{nums}[j-1]$，对于 $x$，小于它的数的个数加一，即 $\textit{less}[j][x]$ 加一。
//
// 代码实现时，可以在枚举 $j$ 的同时更新 $\textit{less}$，并且只需要一个数组。
//
// 附：[视频讲解](https://www.bilibili.com/video/BV1mD4y1E7QK/)
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n^2)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$O(n^2)$。
func countQuadruplets2(nums []int) int64 {
	var ans, n = 0, len(nums)
	var less, great = make([][]int, n+1), make([][]int, n+1)

	great[n] = make([]int, n+1)
	for k := n - 1; k >= 2; k-- {
		great[k] = append([]int{}, great[k+1]...) // great[k] 记录 所有 nums中 > nums[k] 的 个数
		for x := 1; x < nums[k]; x++ {
			great[k][x]++ // x < nums[k]，对于 x，大于它的数的个数 +1
		}
	}

	less[0] = make([]int, n+1)
	for j := 0; j < n-2; j++ {
		less[j+1] = append([]int{}, less[j]...) // less[k] 记录 所有 nums中 < nums[k] 的 个数
		for x := nums[j] + 1; x <= n; x++ {
			less[j+1][x]++ // x > nums[j]，对于 x，小于它的数的个数 +1
		}
	}

	for j := 1; j < n-2; j++ {
		x := nums[j]
		for k := j + 1; k < n; k++ {
			y := nums[k]
			if x < y {
				continue
			}
			ans += less[j][y] * great[k][x]
		}
	}
	return int64(ans)
}
