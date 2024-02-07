package main

func trap(height []int) int {

	n := len(height)
	if n < 3 {
		return 0
	}

	var sum = 0
	for i := 1; i < n; i++ {
		v := height[i]

		var left, right = v, v

		for j := i; j < n; j++ {
			if right < height[j] {
				right = height[j]
			}
		}

		for j := 0; j < i; j++ {
			if left < height[j] {
				left = height[j]
			}
		}

		mn := left
		if mn > right {
			mn = right
		}

		if mn <= v {
			continue
		}

		sum += mn - v
	}

	return sum
}

func trap2(height []int) int {
	var (
		ans = 0
		n   = len(height)

		maxPrefix, maxSuffix = 0, 0

		prefix = make([]int, n)
		suffix = make([]int, n)
	)

	for i := 0; i < n; i++ {
		maxPrefix = max(maxPrefix, height[i])
		prefix[i] = maxPrefix

		j := n - i - 1
		maxSuffix = max(maxSuffix, height[j])
		suffix[j] = maxSuffix
	}

	for i, h := range height {
		ans += min(prefix[i], suffix[i]) - h
	}

	return ans
}

// 前后缀分解
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func trap3(height []int) int {
	var n = len(height)
	var prefix, suffix = make([]int, n+1), make([]int, n+1)
	for i := 0; i < n; i++ {
		prefix[i+1] = max(prefix[i], height[i])
	}
	for i := n - 1; i >= 0; i-- {
		suffix[i] = max(suffix[i+1], height[i])
	}

	var ans = 0
	for i, h := range height {
		wall := min(prefix[i+1], suffix[i+1])
		ans += max(0, wall-h)
	}
	return ans
}

// 单调栈
// 时间复杂度：O(n)
// 空间复杂度：O(n)
func trap4(height []int) int {
	var stack []int

	var ans = 0
	for i, h := range height {
		for len(stack) > 0 {
			bottom := height[stack[len(stack)-1]]
			if h < bottom {
				break
			}
			stack = stack[:len(stack)-1]
			if len(stack) == 0 {
				break
			}
			left := stack[len(stack)-1]
			dh := min(height[left], h) - bottom
			w := i - left - 1
			ans += dh * w
		}
		stack = append(stack, i)
	}
	return ans
}
