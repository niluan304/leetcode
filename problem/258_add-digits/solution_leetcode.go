package main

// 模拟
// 最直观的方法是模拟各位相加的过程，直到剩下的数字是一位数。
func leetcode1(num int) int {
	for num >= 10 {
		sum := 0
		for ; num > 0; num /= 10 {
			sum += num % 10
		}
		num = sum
	}
	return num
}

func leetcode2(num int) int {
	return (num-1)%9 + 1
}
