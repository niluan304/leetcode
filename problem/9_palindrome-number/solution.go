package main

func isPalindrome(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	var list []int
	for x != 0 {
		list = append(list, x%10)
		x = x / 10

	}

	n := len(list)
	for i := range list {
		if list[i] != list[n-1-i] {
			return false
		}

		if (i+1)*2 >= n {
			return true
		}
	}

	return true
}

func isPalindrome2(x int) bool {
	if x < 10 {
		if x < 0 {
			return false
		}
		return true
	}

	if x%10 == 0 {
		return false
	}

	var y = 0
	for y < x {
		y = y*10 + x%10
		x /= 10
	}

	return y == x || y/10 == x
}
