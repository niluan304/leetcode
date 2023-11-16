package main

import "sort"

// 贪心
// 时间复杂度：O(nlogn)
// 空间复杂度：O(n)
func splitNum(num int) int {
	var splits = make([]int, 0, 10)
	for num != 0 {
		splits = append(splits, num%10)
		num /= 10
	}
	sort.Ints(splits)

	var num1, num2 = 0, 0
	for i, v := range splits {
		if i%2 == 0 {
			num1 = num1*10 + v
		} else {
			num2 = num2*10 + v
		}
	}
	return num1 + num2
}
