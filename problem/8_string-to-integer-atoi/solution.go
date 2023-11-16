package main

import (
	"math"
)

func myAtoi(s string) int {
	var (
		sign = 0 // 标识数字的正负号
		num  = 0
	)
	for _, c := range s {
		if sign == 0 {
			switch c {
			case '-':
				sign = -1
				continue
			case '+':
				sign = 1
				continue
			case ' ':
				continue
			}
		}

		if num > math.MaxInt32 || c < '0' || c > '9' {
			break
		}

		// 与1进行位或运算, 等价于 if sign == 0 {sign = 1}
		sign |= 1
		num = num*10 + int(c-'0')
	}

	num *= sign
	if num > math.MaxInt32 {
		return math.MaxInt32
	}
	if num < math.MinInt32 {
		return math.MinInt32
	}
	return num
}
