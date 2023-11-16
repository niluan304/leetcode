package main

func plusOne(digits []int) []int {
	var flag = true

	for i := len(digits) - 1; i >= 0; i-- {
		if !flag {
			break
		}

		digits[i]++
		if digits[i] == 10 {
			digits[i] = 0
			flag = true
		} else {
			flag = false
		}
	}

	if flag {
		return append([]int{1}, digits...)
	}
	return digits
}

func plusOne2(digits []int) []int {

	for i := len(digits) - 1; i >= 0; i-- {
		digits[i]++
		if digits[i] != 10 {
			return digits
		}

		digits[i] = 0
	}

	digits[0] = 1
	return append(digits, 0)
}
