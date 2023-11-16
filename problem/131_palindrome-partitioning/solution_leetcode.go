package main

// leetcode 1 回溯 + 动态规划预处理
func leetcode1(s string) (ans [][]string) {
	n := len(s)
	f := make([][]bool, n)
	for i := range f {
		f[i] = make([]bool, n)
		for j := range f[i] {
			f[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			f[i][j] = s[i] == s[j] && f[i+1][j-1]
		}
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if f[i][j] {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}

// leetcode 2 回溯 + 记忆化搜索
func leetcode2(s string) (ans [][]string) {
	n := len(s)
	f := make([][]int8, n)
	for i := range f {
		f[i] = make([]int8, n)
	}

	// 0 表示尚未搜索，1 表示是回文串，-1 表示不是回文串
	var isPalindrome func(i, j int) int8
	isPalindrome = func(i, j int) int8 {
		if i >= j {
			return 1
		}
		if f[i][j] != 0 {
			return f[i][j]
		}
		f[i][j] = -1
		if s[i] == s[j] {
			f[i][j] = isPalindrome(i+1, j-1)
		}
		return f[i][j]
	}

	splits := []string{}
	var dfs func(int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, append([]string(nil), splits...))
			return
		}
		for j := i; j < n; j++ {
			if isPalindrome(i, j) > 0 {
				splits = append(splits, s[i:j+1])
				dfs(j + 1)
				splits = splits[:len(splits)-1]
			}
		}
	}
	dfs(0)
	return
}
