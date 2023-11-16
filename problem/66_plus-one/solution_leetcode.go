package main

// leetcode 1 找出最长的后缀 999
func leetcode1(digits []int) []int {
	n := len(digits)
	for i := n - 1; i >= 0; i-- {
		if digits[i] != 9 {
			digits[i]++
			for j := i + 1; j < n; j++ {
				digits[j] = 0
			}
			return digits
		}
	}
	// digits 中所有的元素均为 9

	digits = make([]int, n+1)
	digits[0] = 1
	return digits
}

// leetcode 2 1.84MB
func leetcode2(digits []int) []int {
	var a int = 1
	for i := 0; i < len(digits); i++ {
		n := digits[len(digits)-a]
		if n != 9 {
			digits[len(digits)-a] = n + 1
			break
		}
		if n == 9 {
			digits[len(digits)-a] = 0
			a++
		} else {
			digits[len(digits)-a] = n + 1
		}
	}
	if digits[0] == 0 {
		digits[0] = 1
		digits = append(digits, 0)
	}
	return digits
}
