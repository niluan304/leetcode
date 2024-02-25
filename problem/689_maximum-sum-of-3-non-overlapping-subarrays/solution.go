package main

import (
	"cmp"
)

// 前后缀分解 + 枚举中间的子数组
//
// 1. 将 子数组 nums[i:i+k-1] 的和，转变为 sum[i]，这样更方便计算
// 2. 前后缀分解
// 2.1 通过前缀最大和数组 preMax 求出 sum[:i] 的最大值的索引
// 2.2 通过后缀最大和数组 sufMax 求出 sum[i:] 的最大值的索引
// 3. 枚举 sum[k: n-2k]
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maxSumOfThreeSubarrays(nums []int, k int) []int {
	n := len(nums)
	m := n - k + 1

	sum := make([]int, m)
	sum[0] = Sum(nums[:k])
	for i := 0; i < m-1; i++ {
		sum[i+1] = sum[i] + nums[i+k] - nums[i]
	}

	// 前后缀分解

	preMax := make([]int, m)
	preMax[0] = 0
	for i := 1; i < m; i++ {
		preMax[i] = i
		if j := preMax[i-1]; sum[j] >= sum[i] { // 从前向后遍历，如果 sum[j] == sum[i]，j 的字典序更小，需要更新
			preMax[i] = j
		}
	}

	sufMax := make([]int, m)
	sufMax[m-1] = m - 1
	for i := m - 2; i >= 0; i-- {
		sufMax[i] = i
		if j := sufMax[i+1]; sum[j] > sum[i] { // 从后向前遍历，如果 sum[j] == sum[i]，i 的字典序更小，无需更新
			sufMax[i] = j
		}
	}

	mx := 0
	var ans []int
	for i := k; i < m-k; i++ { // 枚举中间的子数组
		l, r := preMax[i-k], sufMax[i+k]
		s := sum[l] + sum[i] + sum[r]
		if s > mx {
			ans = []int{l, i, r}
			mx = s
		}
	}
	return ans
}

func Sum[S ~[]E, E cmp.Ordered](x S) E {
	var m E
	for i := 0; i < len(x); i++ {
		m += x[i]
	}
	return m
}

// 要计算三个无重叠子数组的最大和，我们可以枚举第三个子数组的位置，同时维护前两个无重叠子数组的最大和及其位置。
//
// 要计算两个无重叠子数组的最大和，我们可以枚举第二个子数组的位置，同时维护第一个子数组的最大和及其位置。
//
// 因此，我们首先来解决单个子数组的最大和问题，然后解决两个无重叠子数组的最大和问题，最后解决三个无重叠子数组的最大和问题。
//
// #### 前言一：单个子数组的最大和
//
// 我们用滑动窗口来解决这一问题。
//
// 滑动窗口是数组/字符串问题中常用的抽象概念。**窗口**通常是指在数组/字符串中由开始和结束索引定义的一系列元素的集合，即闭区间 $[i,j]$。而**滑动窗口**是指可以将两个边界向某一方向「滑动」的窗口。例如，我们将 $[i,j]$ 向右滑动 $1$ 个元素，它将变为 $[i+1,j+1]$。
//
// 设 $\textit{sum}_1$ 为大小为 $k$ 的窗口的元素和，当窗口从 $[i-k+1,i]$ 向右滑动 $1$ 个元素后，$\textit{sum}_1$ 增加了 $\textit{nums}[i+1]$，减少了 $\textit{nums}[i-k+1]$。据此我们可以 $O(1)$ 地计算出向右滑动 $1$ 个元素后的窗口的元素和。
//
// 我们从 $[0,k-1]$ 开始，不断地向右滑动窗口，直至窗口右端点到达数组末尾时停止。统计这一过程中的 $\textit{sum}_1$ 的最大值（记作 $\textit{maxSum}_1$）及其对应位置。
//
// #### 前言二：两个无重叠子数组的最大和
//
// 我们使用两个大小为 $k$ 的滑动窗口。设 $\textit{sum}_1$ 为第一个滑动窗口的元素和，该滑动窗口从 $[0,k-1]$ 开始；$\textit{sum}_2$ 为第二个滑动窗口的元素和，该滑动窗口从 $[k,2k-1]$ 开始。
//
// 我们同时向右滑动这两个窗口，并维护 $\textit{sum}_1$ 的最大值 $\textit{maxSum}_1$ 及其对应位置。每次滑动时，计算当前 $\textit{maxSum}_1$ 与 $\textit{sum}_2$ 之和。统计这一过程中的 $\textit{maxSum}_1+\textit{sum}_2$ 的最大值（记作 $\textit{maxSum}_{12}$）及其对应位置。
// #### 方法一：滑动窗口
//
// 回到本题，我们使用三个大小为 $k$ 的滑动窗口。设 $\textit{sum}_1$ 为第一个滑动窗口的元素和，该滑动窗口从 $[0,k-1]$ 开始；$\textit{sum}_2$ 为第二个滑动窗口的元素和，该滑动窗口从 $[k,2k-1]$ 开始；$\textit{sum}_3$ 为第三个滑动窗口的元素和，该滑动窗口从 $[2k,3k-1]$ 开始。
//
// 我们同时向右滑动这三个窗口，按照前言二的方法并维护 $\textit{maxSum}_{12}$ 及其对应位置。每次滑动时，计算当前 $\textit{maxSum}_{12}$ 与 $\textit{sum}_3$ 之和。统计这一过程中的 $\textit{maxSum}_{12}+\textit{sum}_3$ 的最大值及其对应位置。
//
// 对于题目要求的最小字典序，由于我们是从左向右遍历的，并且仅当元素和超过最大元素和时才修改最大元素和，从而保证求出来的下标列表是字典序最小的。
//
// **复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。
//
// - 空间复杂度：$O(1)$。我们只需要常数空间来存放若干变量。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-sum-of-3-non-overlapping-subarrays/solutions/1145000/san-ge-wu-zhong-die-zi-shu-zu-de-zui-da-4a8lb/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func maxSumOfThreeSubarrays2(nums []int, k int) []int {
	n := len(nums)
	sum1, sum2, sum3 := 0, 0, 0
	maxSum1, maxSum12, maxTotal := 0, 0, 0
	maxSum1Idx, maxSum12Idx1, maxSum12Idx2 := 0, 0, 0
	var ans []int
	for i := 2 * k; i < n; i++ {
		sum1 += nums[i-2*k]
		sum2 += nums[i-k]
		sum3 += nums[i]
		if i >= 3*k-1 {
			if sum1 > maxSum1 {
				maxSum1 = sum1
				maxSum1Idx = i - 3*k + 1
			}
			if maxSum1+sum2 > maxSum12 {
				maxSum12 = maxSum1 + sum2
				maxSum12Idx1, maxSum12Idx2 = maxSum1Idx, i-2*k+1
			}
			if maxSum12+sum3 > maxTotal {
				maxTotal = maxSum12 + sum3
				ans = []int{maxSum12Idx1, maxSum12Idx2, i - k + 1}
			}
			sum1 -= nums[i-3*k+1]
			sum2 -= nums[i-2*k+1]
			sum3 -= nums[i-k+1]
		}
	}
	return ans
}

// todo solution with dp
func maxSumOfThreeSubarrays3(nums []int, k int) []int {
	n := len(nums)
	m := n - k + 1

	sum := make([]int, m)
	sum[0] = Sum(nums[:k])
	for i := 0; i < m-1; i++ {
		sum[i+1] = sum[i] + nums[i+k] - nums[i]
	}

	return nil
}
