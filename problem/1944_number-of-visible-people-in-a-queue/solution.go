package main

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
// Deprecated: 超时
func canSeePersonsCount(heights []int) []int {
	var n = len(heights)
	var ans = make([]int, n)

	for i := 0; i < n; i++ {
		maxH := 0
		for j := i + 1; j < n; j++ {
			h := heights[j]
			if h > maxH { // 高度各不相同
				maxH = h
				ans[i]++
			}
			if h > heights[i] {
				break
			}
		}
	}
	return ans
}

// 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func canSeePersonsCount2(heights []int) []int {
	var n = len(heights)
	var ans = make([]int, n)
	var stack []int
	for i, h := range heights {
		m := len(stack)
		for ; m > 0; m-- {
			j := stack[m-1]
			if heights[j] > h { // 题目高度各不相同，因此不用判断等于
				break
			}

			ans[j]++ // 出栈代表新到来的比 heights[j] 更高，看到比自己高的，就代表这个位置的观测结束了。
		}

		stack = append(stack[:m], i)
		if m > 0 {
			ans[stack[m-1]]++ // 入栈代表新到来的比 栈顶的人 还矮，栈顶的看到的人数 +1。
		}
	}

	return ans
}
