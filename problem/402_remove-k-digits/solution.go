package main

import (
	"strings"
)

// 贪心 + 单调栈
//
// - 时间复杂度：O(n)。
// - 空间复杂度：O(1)。
func removeKdigits(num string, k int) string {
	n := len(num)
	stack := make([]byte, 0, n) // 递增单调栈
	stack = append(stack, '0')
	for i := 0; i < n; i++ {
		digit := num[i]
		for m := len(stack); m > 0 && k > 0; m-- {
			if digit >= stack[m-1] {
				break
			}
			stack = stack[:m-1]
			k--
		}
		stack = append(stack, digit)
	}

	// 去掉 前缀 0，如果为空字符串，还需要返回 "0"
	stack = stack[:len(stack)-k]
	ans := strings.TrimLeft(string(stack), "0")
	if ans == "" {
		return "0"
	}
	return ans
}
