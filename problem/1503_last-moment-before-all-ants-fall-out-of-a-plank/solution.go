package main

func getLastMoment(n int, left []int, right []int) int {
	var ans = 0

	for _, v := range left {
		ans = _max(ans, v)
	}

	for _, v := range right {
		ans = _max(ans, n-v)
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
