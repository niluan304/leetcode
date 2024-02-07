package main

import "sort"

// 方法一：排序 + 二分查找
// 对于正整数 a,b,ca, b, ca,b,c，它们可以作为三角形的三条边，当且仅当：
// a+b>c
// a+c>b
// b+c>a
// 均成立。如果我们将三条边进行升序排序，使它们满足 a≤b≤c ，那么 a+c>b 和 b+c>a 使一定成立的，我们只需要保证 a+b>c。
//
// 因此，我们可以将数组 nums 进行升序排序，随后使用二重循环枚举 a 和 b。设 a=nums[i],b=nums[j]，
// 为了防止重复统计答案，我们需要保证 i<j。剩余的边 c 需要满足 c<nums[i]+nums[j]，
// 我们可以在 [j+1,n−1] 的下标范围内使用二分查找（其中 n 是数组 nums 的长度），找出最大的满足 nums[k]<nums[i]+nums[j] 的下标 k，
// 这样一来，在 [j+1,k] 范围内的下标都可以作为边 c 的下标，我们将该范围的长度 k−j 累加入答案。
//
// 当枚举完成后，我们返回累加的答案即可。
func leetcode1(nums []int) (ans int) {
	sort.Ints(nums)
	for i, v := range nums {
		for j := i + 1; j < len(nums); j++ {
			ans += sort.SearchInts(nums[j+1:], v+nums[j])
		}
	}
	return
}

// 方法二：排序 + 双指针
// 我们可以对方法一进行优化。
//
// 我们将当 a=nums[i],b=nums[j] 时，最大的满足 nums[k]<nums[i]+nums[j] 的下标 k 记为 k[i,j] 。
// 可以发现，如果我们固定 i，那么随着 j 的递增，不等式右侧 nums[i]+nums[j]也是递增的，因此 k[i,j]	也是递增的。
//
// 这样一来，我们就可以将 j 和 k 看成两个同向（递增）移动的指针，将方法一进行如下的优化：
// 我们使用一重循环枚举 i。当 i 固定时，我们使用双指针同时维护 j 和 k，它们的初始值均为 i；
// 我们每一次将 j 向右移动一个位置，即 j←j+1，并尝试不断向右移动 k，使得 k 是最大的满足 nums[k]<nums[i]+nums[j] 的下标。
// 我们将 max(k−j,0) 累加入答案。
//
// 当枚举完成后，我们返回累加的答案即可。
func leetcode2(nums []int) (ans int) {
	n := len(nums)
	sort.Ints(nums)
	for i, v := range nums {
		k := i
		for j := i + 1; j < n; j++ {
			for k+1 < n && nums[k+1] < v+nums[j] {
				k++
			}
			ans += max(k-j, 0)
		}
	}
	return
}
