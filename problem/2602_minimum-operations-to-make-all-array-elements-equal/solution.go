package main

import "sort"

// 前缀和 + 二分查找
// 时间复杂度: O(nlogn + mlogn)
// 空间复杂度: O(n)
func minOperations(nums []int, queries []int) []int64 {
	var ans = make([]int64, len(queries))

	var n = len(nums)
	sort.Ints(nums)
	var prefixSum = make([]int, n+1)
	for i, num := range nums {
		prefixSum[i+1] = prefixSum[i] + num
	}

	for i, query := range queries {
		j := sort.SearchInts(nums, query)                  // [0,j-1] 都小于 query，[j,n-1] 的数都大于等于 query
		left := query*j - prefixSum[j]                     // 小于 query 的操作：query*小于的数量(j-1 - 0 + 1) - [0,j-1]的前缀和
		right := prefixSum[n] - prefixSum[j] - query*(n-j) // 大于 query 的操作：[j,n-1]的前缀和 - query * 大于的数量(n-1 - j + 1)
		ans[i] = int64(left + right)
	}
	return ans
}
