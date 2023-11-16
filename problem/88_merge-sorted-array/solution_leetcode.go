package main

import "sort"

func leetcode1(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	copy(nums1[m:], nums2)
	sort.Ints(nums1)
}

func leetcode2(nums1 []int, m int, nums2 []int, n int) {
	sorted := make([]int, 0, m+n)
	p1, p2 := 0, 0
	for {
		if p1 == m {
			sorted = append(sorted, nums2[p2:]...)
			break
		}
		if p2 == n {
			sorted = append(sorted, nums1[p1:]...)
			break
		}
		if nums1[p1] < nums2[p2] {
			sorted = append(sorted, nums1[p1])
			p1++
		} else {
			sorted = append(sorted, nums2[p2])
			p2++
		}
	}
	copy(nums1, sorted)
}

// leetcode 3 方法二中，之所以要使用临时变量，是因为如果直接合并到数组 `nums1` 中，
// `nums1` 中的元素可能会在取出之前被覆盖。那么如何直接避免覆盖`nums1`中的元素呢？
// 观察可知，`nums1` 的后半部分是空的，可以直接覆盖而不会影响结果。因此可以指针设置
// 为从后向前遍历，每次取两者之中的较大者放进 `nums1` 的最后面。
func leetcode3(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		var cur int
		if p1 == -1 {
			cur = nums2[p2]
			p2--
		} else if p2 == -1 {
			cur = nums1[p1]
			p1--
		} else if nums1[p1] > nums2[p2] {
			cur = nums1[p1]
			p1--
		} else {
			cur = nums2[p2]
			p2--
		}
		nums1[tail] = cur
	}
}
