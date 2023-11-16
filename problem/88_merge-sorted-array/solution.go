package main

import (
	"sort"
)

func merge(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	nums1 = append(nums1[:m], nums2...)
	sort.Ints(nums1)
}

func merge2(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}
	p1, p2 := m-1, n-1
	for tail := m + n - 1; p1 >= 0 && p2 >= 0; tail-- {
		if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
	}

	if p1 == -1 {
		copy(nums1[:p2+1], nums2[:p2+1])
	}

	return
}

func merge3(nums1 []int, m int, nums2 []int, n int) {
	if n == 0 {
		return
	}

	// 前序遍历不好存放变临时变量, 如果借用 nums1 为0的空间, 就可以不用存放临时变量了。
	for p1, p2, tail := m-1, n-1, m+n-1; p1 >= 0 || p2 >= 0; tail-- {
		if p2 == -1 {
			break
		} else if p1 == -1 {
			copy(nums1[:p2+1], nums2[:p2+1])
			break
		} else if nums1[p1] > nums2[p2] {
			nums1[tail] = nums1[p1]
			p1--
		} else {
			nums1[tail] = nums2[p2]
			p2--
		}
	}
	return
}
