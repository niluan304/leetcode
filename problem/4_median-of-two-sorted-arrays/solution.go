package main

import (
	"math"
	"sort"
)

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {

	nums1 = append(nums1, nums2...)
	sort.Ints(nums1)

	n := len(nums1)
	mid := n / 2

	if n%2 == 0 {
		return float64(nums1[mid]+nums1[mid-1]) / 2
	} else {
		return float64(nums1[mid])
	}
}

func FindMedianSortedArrays2(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	l := m + n
	mid1 := l / 2
	mid2 := mid1
	if l%2 == 0 {
		mid2 -= 1
	}

	var v1, v2, flag int

	p1, p2, min := 0, 0, 0
	for p1 < m || p2 < n {
		if p1 == m {
			min = nums2[p2]
			p2++
		} else if p2 == n {
			min = nums1[p1]
			p1++
		} else {
			if nums1[p1] < nums2[p2] {
				min = nums1[p1]
				p1++
			} else {
				min = nums2[p2]
				p2++
			}
		}

		if p1+p2-1 == mid1 {
			v1 = min
			flag++
		}
		if p1+p2-1 == mid2 {
			v2 = min
			flag++
		}
		if flag == 2 {
			break
		}
	}

	return float64(v1+v2) / 2
}

func FindMedianSortedArrays3(nums1 []int, nums2 []int) float64 {
	m, n := len(nums1), len(nums2)

	l := m + n
	mid1 := l / 2
	mid2 := mid1
	if l%2 == 0 {
		mid2 -= 1
	}

	nums1 = append(nums1, math.MaxInt)
	nums2 = append(nums2, math.MaxInt)

	var v1, v2, flag int
	p1, p2, min := 0, 0, 0
	for p1 < m || p2 < n {
		if nums1[p1] < nums2[p2] {
			min = nums1[p1]
			p1++
		} else {
			min = nums2[p2]
			p2++
		}

		if p1+p2-1 == mid1 {
			v1 = min
			flag++
		}
		if p1+p2-1 == mid2 {
			v2 = min
			flag++
		}
		if flag == 2 {
			break
		}
	}

	return float64(v1+v2) / 2
}
