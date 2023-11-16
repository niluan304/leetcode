package main

import "math"

// 暴力解法
// 时间复杂度: O(n^4)
// 空间复杂度: O(1)
func fourSumCount(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		ans = 0
	)

	for _, n1 := range nums1 {
		for _, n2 := range nums2 {
			for _, n3 := range nums3 {
				for _, n4 := range nums4 {
					if n1+n2+n3+n4 == 0 {
						ans++
					}
				}
			}
		}

	}

	return ans
}

// 使用哈希表保存数组出现次数
// 时间复杂度: O(n^3)
// 空间复杂度: O(n)
func fourSumCount2(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		ans = 0

		m1, m2, m3, m4 = map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}
	)

	for i := 0; i < len(nums1); i++ {
		m1[nums1[i]]++
		m2[nums2[i]]++
		m3[nums3[i]]++
		m4[nums4[i]]++
	}

	for n1, c1 := range m1 {
		for n2, c2 := range m2 {
			for n3, c3 := range m3 {
				// 不存在则 c4 = 0
				c4 := m4[0-(n1+n2+n3)]
				ans += c1 * c2 * c3 * c4
			}
		}
	}

	return ans
}

// 使用哈希表保存数组出现次数
// 时间复杂度: O(n^3)
// 空间复杂度: O(n)
// 额外校验最大值和最小值
func fourSumCount3(nums1 []int, nums2 []int, nums3 []int, nums4 []int) int {
	var (
		ans = 0

		m1, m2, m3, m4 = map[int]int{}, map[int]int{}, map[int]int{}, map[int]int{}

		min1, min2, min3, min4 = math.MaxInt, math.MaxInt, math.MaxInt, math.MaxInt
		max1, max2, max3, max4 = math.MinInt, math.MinInt, math.MinInt, math.MinInt
	)

	for i := 0; i < len(nums1); i++ {
		var v1, v2, v3, v4 = nums1[i], nums2[i], nums3[i], nums4[i]

		min1 = _min(min1, v1)
		min2 = _min(min2, v2)
		min3 = _min(min3, v3)
		min4 = _min(min4, v4)

		max1 = _max(max1, v1)
		max2 = _max(max2, v2)
		max3 = _max(max3, v3)
		max4 = _max(max4, v4)

		m1[nums1[i]]++
		m2[nums2[i]]++
		m3[nums3[i]]++
		m4[nums4[i]]++
	}

	if min1+min2+min3+min4 > 0 || max1+max2+max3+max4 < 0 {
		return 0
	}

	for n1, c1 := range m1 {
		for n2, c2 := range m2 {
			for n3, c3 := range m3 {
				// 不存在则 c4 = 0
				c4 := m4[0-(n1+n2+n3)]
				ans += c1 * c2 * c3 * c4
			}
		}
	}

	return ans
}

func _min(x, y int) int {
	if x < y {
		return x
	}

	return y
}

func _max(x, y int) int {
	if x > y {
		return x
	}

	return y
}
