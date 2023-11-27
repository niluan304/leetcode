package main

import "sort"

// 单调栈：从右向左 二分查找索引更新位置
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func dailyTemperatures(temperatures []int) []int {
	var n = len(temperatures)
	var ans = make([]int, n)

	//// 额外保存 value 的写法
	//type Value struct{ Idx, T int }
	//var values = []Value{{Idx: n - 1, T: 0}}
	//
	//for i := n - 1; i >= 0; i-- {
	//	t := temperatures[i]
	//	m := len(values)
	//	j := m - sort.Search(m, func(i int) bool {
	//		return t < values[m-1-i].T
	//	})
	//	if j > 0 {
	//		ans[i] = values[j-1].Idx - i
	//	}
	//	values = values[:j]
	//	values = append(values, Value{Idx: i, T: t})
	//}

	// 只存下标的写法
	var stack []int
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]

		m := len(stack)
		j := m - sort.Search(m, func(i int) bool {
			// 本题单调栈是单调递减的，因此需要 "取反" 索引
			return t < temperatures[stack[m-1-i]]
		})
		if j > 0 {
			ans[i] = stack[j-1] - i
		}
		stack = stack[:j] // 延迟出栈
		stack = append(stack, i)
	}
	return ans
}

// 单调栈：从右向左 遍历查找索引更新位置
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func dailyTemperatures2(temperatures []int) []int {
	var n = len(temperatures)
	var ans = make([]int, n)

	// 只存下标的写法
	var stack []int
	for i := n - 1; i >= 0; i-- {
		t := temperatures[i]

		m := len(stack)
		for m > 0 && t >= temperatures[stack[m-1]] {
			m--
		}
		if m > 0 {
			ans[i] = stack[m-1] - i
		}
		stack = stack[:m]
		stack = append(stack, i)
	}
	return ans
}

// 单调栈：从左向右
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func dailyTemperatures3(temperatures []int) []int {
	var n = len(temperatures)
	var ans = make([]int, n)
	var stack []int

	for i, t := range temperatures {
		m := len(stack)
		for m > 0 && t > temperatures[stack[m-1]] {
			m--
			j := stack[m]
			ans[j] = i - j
		}
		stack = stack[:m]
		stack = append(stack, i)
	}
	return ans
}
