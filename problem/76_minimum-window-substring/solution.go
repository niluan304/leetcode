package main

import "math"

// 滑动窗口, 判断 Contain 时间复杂度为 O(mn), 总时间复杂度 O(mm)+O(n)
func minWindow(s string, t string) string {
	var (
		m = len(s)
		n = len(t)
	)

	if m < n {
		return ""
	}

	var (
		total = make(map[byte]int, 26*2)
		count = make(map[byte]int, 26*2)

		x, y = 0, math.MaxInt
		left = 0
	)

	// 时间复杂度: O(n)
	for i := 0; i < n; i++ {
		total[t[i]]++
	}

	Contain := func() bool {
		for b, v := range total {
			if count[b] < v {
				return false
			}
		}

		return true
	}

	// 时间复杂度: O(mn)
	for right := 0; right < m; right++ {
		count[s[right]]++

		// 时间复杂度: O(n)
		for Contain() {
			if y-x > right-left {
				x, y = left, right
			}

			count[s[left]]--
			left++
		}
	}

	if y == math.MaxInt {
		return ""
	}

	return s[x : y+1]
}

// 滑动窗口, 使用变量 distance 记录 total和count的差值, 时间复杂度 O(m+n)
func minWindow2(s string, t string) string {
	var (
		m = len(s)
		n = len(t)
	)

	if m < n {
		return ""
	}

	var (
		total = make(map[byte]int, 26*2)
		count = make(map[byte]int, 26*2)

		x, y = 0, math.MaxInt
		left = 0

		distance = n
	)

	// 时间复杂度: O(n)
	for i := 0; i < n; i++ {
		total[t[i]]++
	}

	// 时间复杂度: O(m)
	for right := 0; right < m; right++ {
		b := s[right]
		count[b]++

		if count[b] <= total[b] {
			distance--
		}

		// 等价于 for distance == 0
		for distance <= 0 {
			if y-x > right-left {
				x, y = left, right
			}

			b2 := s[left]
			count[b2]--
			left++

			if count[b2] < total[b2] {
				distance++
			}
		}
	}

	if y == math.MaxInt {
		return ""
	}

	return s[x : y+1]
}

// 滑动窗口, 使用变量 distance 记录 total和count的差值, 时间复杂度 O(m+n)
// 并使用数组替代哈希表, 以提升读写效率
func minWindow3(s string, t string) string {
	var (
		m = len(s)
		n = len(t)
	)

	if m < n {
		return ""
	}

	var (
		// total = [123]int{} // ascii 'z'  = 122, 与[128]等价
		total = [128]int{}
		count = [128]int{}

		x, y = 0, math.MaxInt // 保存最小窗口的左右边界
		left = 0              // 滑动窗口的左边界

		distance = n // 记录 total和count的差值
	)

	// 时间复杂度: O(n)
	for i := 0; i < n; i++ {
		total[t[i]]++
	}

	// 时间复杂度: O(m)
	for right := 0; right < m; right++ {
		b := s[right]
		if count[b] < total[b] {
			distance--
		}
		count[b]++

		// 等价于 for distance <= 0
		for distance == 0 {
			if y-x > right-left {
				x, y = left, right
			}

			b2 := s[left]
			if count[b2] == total[b2] {
				distance++
			}
			count[b2]--
			left++
		}
	}

	if y == math.MaxInt {
		return ""
	}

	return s[x : y+1]
}
