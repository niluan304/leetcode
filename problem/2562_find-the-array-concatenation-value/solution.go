package main

func findTheArrayConcVal(nums []int) int64 {
	var ans = 0

	var left, right = 0, len(nums) - 1
	for left < right {
		x, y := nums[left], nums[right]
		for y > 0 {
			y /= 10
			x *= 10
		}
		//x := nums[left] * int(math.Pow(10, float64(len(strconv.Itoa(nums[right])))))
		ans += x + nums[right]

		left++
		right--
	}

	if left == right {
		ans += nums[left]
	}

	return int64(ans)
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

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
