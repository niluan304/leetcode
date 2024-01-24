package main

// 前后缀分解 + 单调栈
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumSumOfHeights(maxHeights []int) int64 {

	var n = len(maxHeights)
	var ans = 0 // math.MaxInt32 // math.MinInt32
	var left, right = make([]int, n+1), make([]int, n+1)

	var stack = []int{-1}
	for i := 0; i < n; i++ {
		m := len(stack)
		for m > 1 && maxHeights[i] <= maxHeights[stack[m-1]] {
			m--
		}
		top := stack[m-1]
		stack = append(stack[:m], i)

		left[i+1] = (i-top)*maxHeights[i] + left[top+1]
	}

	stack = []int{n}
	for i := n - 1; i >= 0; i-- {
		m := len(stack)
		for m > 1 && maxHeights[i] <= maxHeights[stack[m-1]] {
			m--
		}
		top := stack[m-1]
		stack = append(stack[:m], i)

		right[i] = (top-i)*maxHeights[i] + right[top]
	}

	for i, height := range maxHeights {
		ans = max(ans, left[i+1]+right[i]-height)
	}

	return int64(ans)
}
