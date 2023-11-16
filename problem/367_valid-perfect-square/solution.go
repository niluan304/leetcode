package main

func isPerfectSquare(num int) bool {
	var left, right = 0, num

	for left <= right {
		mid := left + (right-left)/2
		x := mid * mid
		if x == num {
			return true
		}

		if x > num {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return false
}
