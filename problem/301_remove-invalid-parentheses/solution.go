package main

import (
	. "github.com/niluan304/leetcode/container"
)

// 组合型回溯模板1
// 答案的视角，选或不选
// 回溯三问：
// 当前操作？当前字符选或不选
// 子问题？从字符串 s[i:] 中拼凑出平衡的括号
// 下一个子问题？从字符串 s[i+1:] 中拼凑出平衡的括号
func removeInvalidParentheses(s string) []string {
	set := NewSet[string]()

	var path []byte
	n, mx := len(s), 0
	values := map[byte]int{'(': +1, ')': -1}

	var dfs func(i int, balance int)
	dfs = func(i int, balance int) {
		if i == n {
			if balance == 0 {
				if len(path) > mx {
					set.Clear()
					mx = len(path)
				}
				if len(path) == mx {
					set.Add(string(path))
				}
			}
			return
		}
		// 剪枝
		if len(path)+n-i < mx || balance > n-i {
			return
		}

		x := s[i]
		switch x {
		case '(':
			dfs(i+1, balance) // 不选
		case ')':
			dfs(i+1, balance) // 不选
			if balance <= 0 {
				return // 不合法的方案，不可以选
			}
		default:
			// 剪枝优化：如果是字母，必须要选
		}

		path = append(path, x) // 选
		dfs(i+1, balance+values[x])
		path = path[:len(path)-1]
	}

	dfs(0, 0)
	return set.Values()
}
