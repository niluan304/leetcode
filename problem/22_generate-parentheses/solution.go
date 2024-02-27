package main

import (
	"bytes"
	"slices"

	. "github.com/niluan304/leetcode/container"
	. "github.com/niluan304/leetcode/copypasta/dp"
)

// 组合型回溯模板1
// 从输入的角度出发：选或不选
//
// 回溯三问：
// 当前操作？枚举 path[i] 是左括号还是右括号
// 子问题？构造字符串 >= i 的部分
// 下一个子问题？构造字符串 >= i+1 的部分
//
// - 时间复杂度：$\mathcal{O}(n \cdot C^{n}_{2n})$。
// - 空间复杂度：$\mathcal{O}(n)$。
func generateParenthesis(n int) []string {
	var ans []string
	path := make([]byte, n*2) // 由于答案长度是固定，因此替代数组的 push, pop 操作

	var dfs func(i int, open int)
	dfs = func(i int, open int) {
		if i == 2*n {
			ans = append(ans, string(path))
			return
		}

		if open < n {
			path[i] = '('
			dfs(i+1, open+1)
		}
		if open*2 > i {
			path[i] = ')'
			dfs(i+1, open)
		}
	}

	dfs(0, 0)

	slices.Sort(ans[:]) // 额外做升序排列，仅用于 debug
	return ans
}

// 组合型回溯模板2
// 从答案的角度出发：枚举下一个左括号的位置
//
// 回溯三问：
// 当前操作？
// 子问题？
// 下一个子问题？
//
// - 时间复杂度：$\mathcal{O}(n \cdot C^{n}_{2n})$。
// - 空间复杂度：$\mathcal{O}(n)$。
func generateParenthesis2(n int) (ans []string) {
	path := []int{}
	var dfs func(int, int)
	dfs = func(i, balance int) { // balance = 左括号个数 - 右括号个数
		if len(path) == n {
			s := bytes.Repeat([]byte{')'}, n*2)
			for _, j := range path {
				s[j] = '('
			}
			ans = append(ans, string(s))
			return
		}
		// 可以填 0 到 balance 个右括号
		for c := 0; c <= balance; c++ { // 填 c 个右括号
			path = append(path, i+c) // 填 1 个左括号
			dfs(i+c+1, balance-c+1)
			path = path[:len(path)-1]
		}
	}
	dfs(0, 0)

	slices.Sort(ans[:]) // 额外做升序排列，仅用于 debug
	return
}

// 递归版本的有效括号定义
// 1.空字符串 "" 是一个合法的括号
// 2. 如果 "x" 是合法的括号，那么 "(x)" 也是合法的括号
// 3. 如果 "x","y" 都是合法的括号，那么 "xy" 也是合法的括号
// 定义的 dfs 函数为纯函数
func generateParenthesis3(n int) []string {
	var dfs func(n int) []string
	dfs = func(n int) []string {
		if n == 0 {
			return []string{""} // 1. 空字符串 "" 是一个合法的括号
		}

		set := NewSet[string]()
		for _, value := range dfs(n - 1) {
			set.Add("(" + value + ")") // 2. 如果 "x" 是合法的括号，那么 "(x)" 也是合法的括号
		}

		for i := 1; i < n/2+1; i++ {
			for _, x := range dfs(i) {
				for _, y := range dfs(n - i) {
					set.Add(x+y, y+x) // 3. 如果 "x","y" 都是合法的括号，那么 "xy" 也是合法的括号
				}
			}
		}
		return set.Values()
	}

	MemorySearch(&dfs)

	ans := dfs(n)
	slices.Sort(ans[:]) // 额外做升序排列，仅用于 debug
	return ans
}
