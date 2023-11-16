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

		min := left
		if min > right {
			min = right
		}

		if min <= v {
			continue
		}

		sum += min - v
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
		maxPrefix = _max(maxPrefix, height[i])
		prefix[i] = maxPrefix

		j := n - i - 1
		maxSuffix = _max(maxSuffix, height[j])
		suffix[j] = maxSuffix
	}

	for i, h := range height {
		ans += _min(prefix[i], suffix[i]) - h
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}

	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}

	return y
}
