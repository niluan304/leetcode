package main

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func largestRectangleArea(heights []int) int {
	var ans = 0

	for i, height := range heights {
		x, y := i, i
		for x < len(heights) && heights[x] >= height {
			x++
		}
		for y >= 0 && heights[y] >= height {
			y--
		}
		ans = max(ans, height*(x-y-1))
	}
	return ans
}

// 前后单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func largestRectangleArea2(heights []int) int {
	var n = len(heights)
	var preWeights = make([]int, n)
	var preStack []int
	for i := 0; i < n; i++ {
		preWeights[i] = 1
		for m := len(preStack); m > 0; m-- {
			j := preStack[m-1]
			if heights[i] > heights[j] {
				break
			}
			preStack = preStack[:m-1]
			preWeights[i] += preWeights[j]
		}
		preStack = append(preStack, i)
	}

	var sufWeights = make([]int, n)
	var sufStack []int
	for i := n - 1; i >= 0; i-- {
		sufWeights[i] = 1
		for m := len(sufStack); m > 0; m-- {
			j := sufStack[m-1]
			if heights[i] > heights[j] {
				break
			}
			sufStack = sufStack[:m-1]
			sufWeights[i] += sufWeights[j]
		}
		sufStack = append(sufStack, i)
	}

	var ans = 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(preWeights[i]+sufWeights[i]-1))
	}
	return ans
}

func largestRectangleArea3(heights []int) int {
	var n = len(heights)
	f := func(i int, stack *[]int, weights []int) {
		weights[i] = 1
		for m := len(*stack); m > 0; m-- {
			j := (*stack)[m-1]
			if heights[i] > heights[j] {
				break
			}
			*stack = (*stack)[:m-1]
			weights[i] += weights[j]
		}
		*stack = append(*stack, i)
	}

	var stack []int
	var preWeights = make([]int, n)
	for i := 0; i < n; i++ {
		f(i, &stack, preWeights)
	}

	stack = stack[:0]
	var sufWeights = make([]int, n)
	for i := n - 1; i >= 0; i-- {
		f(i, &stack, sufWeights)
	}

	var ans = 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(preWeights[i]+sufWeights[i]-1))
	}
	return ans
}

// 单调栈
// 在方法一中，我们在对位置 i 进行入栈操作时，确定了它的左边界。
// 从直觉上来说，与之对应的我们在对位置 i 进行出栈操作时可以确定它的右边界！
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func largestRectangleArea4(heights []int) int {
	var n = len(heights)
	var right, left = make([]int, n+1), make([]int, n+1)

	// 0 <= heights[i] <= 10^4
	// 因此可以使用 -1 作为哨兵，以判断边界
	heights = append(heights, -1) // 新增 -1 作为边界哨兵
	var stack = []int{-1}
	for i, height := range heights {
		// 在方法一中，我们在对位置 i 进行入栈操作时，确定了它的左边界。
		// 从直觉上来说，与之对应的我们在对位置 i 进行出栈操作时可以确定它的右边界！
		for m := len(stack); m > 1; m-- { // stack[0] = -1, 边界哨兵
			j := stack[m-1]
			if height > heights[j] {
				break
			}
			stack = stack[:m-1]
			right[j] = i
		}

		left[i] = stack[len(stack)-1]
		stack = append(stack, i)
	}

	var ans = 0
	for i := 0; i < n; i++ {
		ans = max(ans, heights[i]*(right[i]-left[i]-1))
	}
	return ans
}
