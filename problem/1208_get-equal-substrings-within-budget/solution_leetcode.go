package main

import "sort"

// 前言
// 假定字符串 s 和 t 的长度均为 n，对于任意 0≤i<n，将 s[i] 变成 t[i] 的开销是 ∣s[i]−t[i]∣，
// 因此可以创建一个长度为 n 的数组 diff，其中 diff[i]=∣s[i]−t[i]∣。
//
// 创建数组 diff 之后，问题转化成计算数组 diff 的元素和不超过 maxCost 的最长子数组的长度。
// 有两种方法可以解决，第一种方法是前缀和 + 二分查找，第二种方法是滑动窗口。
//
// 方法一：前缀和 + 二分查找
// 首先计算数组 diff 的前缀和，创建长度为 n+1 的数组 accDiff，其中 accDiff[0]=0，
// 对于 0≤i<n0 ，有 accDiff[i+1]=accDiff[i]+diff[i]。
//
// 即当 1≤i≤n1 时，accDiff[i] 为 diff 从下标 0 到下标 i−1 的元素和：
//
// 当 diff 的子数组以下标 j 结尾时，需要找到最小的下标 k（k≤j），
// 使得 diff 从下标 k 到 j 的元素和不超过 maxCost，此时子数组的长度是 j−k+1。
// 由于已经计算出前缀和数组 accDiff，因此可以通过 accDiff 得到 diff 从下标 k 到 j 的元素和：
//
// 因此，找到最小的下标 k（k≤jk），使得 diff 从下标 k 到 j 的元素和不超过 maxCost，
// 等价于找到最小的下标 k（k≤jk），使得 accDiff[j+1]−accDiff[k]≤maxCost。
//
// 由于 diff 的的每个元素都是非负的，因此 accDiff 是递增的，对于每个下标 j，
// 可以通过在 accDiff 内进行二分查找的方法找到符合要求的最小的下标 k。
//
// 以下是具体实现方面的细节。
//
// 不需要计算数组 diff 的元素值，而是可以根据字符串 s 和 t 的对应位置字符直接计算 accDiff 的元素值。
//
// 对于下标范围 [1,n] 内的每个 i，通过二分查找的方式，在下标范围 [0,i] 内找到最小的下标 start，
// 使得 accDiff[start]≥accDiff[i]−maxCost，此时对应的 diff 的子数组的下标范围是从 start 到 i−1，
// 子数组的长度是 i−start。
//
// 遍历下标范围 [1,n] 内的每个 i 之后，即可得到符合要求的最长子数组的长度，即字符串可以转化的最大长度。
func leetcode1(s string, t string, maxCost int) (maxLen int) {
	n := len(s)
	accDiff := make([]int, n+1)
	for i, ch := range s {
		accDiff[i+1] = accDiff[i] + abs(int(ch)-int(t[i]))
	}
	for i := 1; i <= n; i++ {
		start := sort.SearchInts(accDiff[:i], accDiff[i]-maxCost)
		maxLen = max(maxLen, i-start)
	}
	return
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
