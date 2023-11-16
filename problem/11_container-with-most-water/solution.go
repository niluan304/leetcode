package main

func maxArea(height []int) int {
	var max = 0
	for i := range height {
		for j := i + 1; j < len(height); j++ {
			area := min(height[i], height[j]) * (j - i)
			if area > max {
				max = area
			}
		}
	}

	return max
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func maxArea2(height []int) int {
	var (
		i   = 0
		j   = len(height) - 1
		max = 0
	)

	for i <= j {
		area := j - i

		if height[i] > height[j] {
			area *= height[j]
			j--
		} else {
			area *= height[i]
			i++
		}

		if area > max {
			max = area
		}

	}

	return max
}

func maxArea3(height []int) int {
	var (
		i   = 0
		j   = len(height) - 1
		max = 0
	)

	for i <= j {
		var area = 0

		if height[i] > height[j] {
			area = (j - i) * height[j]
			j--
		} else {
			area = (j - i) * height[i]
			i++
		}

		if area > max {
			max = area
		}

	}

	return max
}
