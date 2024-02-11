package main

import (
	"strconv"
	"strings"
)

func isAdditiveNumber(num string) bool {
	n := len(num)
	for i := 1; i < n-1; i++ {

		for j := i + 1; j < n; j++ {
			dp0 := num[:i]
			dp1 := num[i:j]
			if dp0 == "00" || dp1 == "00" {
				break
			}

			x := ""
			for len(x) < n-j {
				dp1, dp0 = addStrings(dp0, dp1), dp1
				x += dp1
			}
			if x == num[j:] {
				return true
			}
		}
	}
	return false
}

func addStrings(num1 string, num2 string) string {
	result := ""
	for i, j, carry := len(num1)-1, len(num2)-1, 0; i >= 0 || j >= 0 || carry > 0; i, j = i-1, j-1 {
		digit1 := 0
		digit2 := 0

		if i >= 0 {
			digit1 = int(num1[i] - '0')
		}

		if j >= 0 {
			digit2 = int(num2[j] - '0')
		}

		sum := digit1 + digit2 + carry
		carry = sum / 10
		digit := sum % 10

		result = strconv.Itoa(digit) + result
	}

	return result
}

// 子集型回溯 模板2
//
// 从答案的角度
//
// 1.当前操作？
// 选取字符串截止位置
// 2.子问题？
// 从 num[i:j+1] 找到 path[-1] + path[-2] 的前缀
// 3.下一个子问题？
// 从 num[j:k+1] 找到 path[0] + path[-1] 的前缀
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func isAdditiveNumber2(num string) bool {
	const MinLength = 3 // 题目要求的最低长度
	n := len(num)
	ans := false
	var path []string

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = ans || len(path) >= MinLength
			return // 不合法的方案
		}

		for j := i; j < n; j++ {
			x := num[i : j+1]
			if x == "00" {
				break
			}

			pLen := len(path)
			if pLen >= 2 && x != addStrings(path[pLen-2], path[pLen-1]) {
				continue
			}

			path = append(path, x)
			dfs(j + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}

// 子集型回溯 模板1
//
// 从输入的角度
//
// 1.当前操作？
// 选取字符串截止位置
// 2.子问题？
// 从 num[i:j+1] 找到 path[-1] + path[-2] 的前缀
// 3.下一个子问题？
// 从 num[j:k+1] 找到 path[0] + path[-1] 的前缀
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func isAdditiveNumber3(num string) bool {
	const MinLength = 3 // 题目要求的最低长度
	n := len(num)
	ans := false
	var path []string

	var dfs func(start, end int)
	dfs = func(start, end int) {
		if end == n {
			ans = ans || len(path) >= MinLength
			return // 不合法的方案
		}
		if end < n-1 {
			dfs(start, end+1) // 不选
		}

		x := num[start : end+1] // 选
		if strings.HasPrefix(x, "00") {
			return
		}
		pLen := len(path)
		if pLen >= 2 && x != addStrings(path[pLen-2], path[pLen-1]) {
			return
		}

		path = append(path, x)
		dfs(end+1, end+1)
		path = path[:pLen]
	}

	dfs(0, 0)
	return ans
}
