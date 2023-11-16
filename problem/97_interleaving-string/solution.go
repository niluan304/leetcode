package main

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func isInterleave(s1 string, s2 string, s3 string) bool {
	var n, m = len(s1), len(s2)
	var cache = make([][]int, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]int, m+1)
	}

	const Empty, No, Yes = 0, 1, 2
	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// 边界条件
		if i == 0 {
			return s2[:j] == s3[:j]
		}
		if j == 0 {
			return s1[:i] == s3[:i]
		}

		v := cache[i][j]
		if v != Empty {
			return v == Yes
		}

		// 末尾字符必须有一个一样
		b := s3[i+j-1]
		if b != s1[i-1] && b != s2[j-1] {
			return false
		}

		cache[i][j] = No
		if b == s1[i-1] && dfs(i-1, j) {
			cache[i][j] = Yes
		}
		if b == s2[j-1] && dfs(i, j-1) {
			cache[i][j] = Yes
		}
		return cache[i][j] == Yes
	}

	return dfs(n, m)
}

// error
func isInterleave2(s1 string, s2 string, s3 string) bool {
	var str2 = make([]byte, 0, len(s2))

	var i1 = 0
	for i3 := 0; i3 < len(s3); i3++ {
		b3 := s3[i3]
		if i1 < len(s1) && b3 == s1[i1] {
			i1++
		} else {
			str2 = append(str2, b3)
		}
	}
	return i1 == len(s1) && string(str2) == s2
}

// dfs + 记忆化搜索
// 时间复杂度：O(mn)
// 空间复杂度：O(mn)
func isInterleave3(s1 string, s2 string, s3 string) bool {
	var n, m = len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	var cache = make([][]*bool, n+1)
	for i := 0; i <= n; i++ {
		cache[i] = make([]*bool, m+1)
	}

	var dfs func(i, j int) bool
	dfs = func(i, j int) bool {
		// 边界条件
		if i == 0 {
			return s2[:j] == s3[:j]
		}
		if j == 0 {
			return s1[:i] == s3[:i]
		}

		if cache[i][j] != nil {
			return *(cache[i][j])
		}

		b := s3[i+j-1]
		v := false
		if b == s1[i-1] {
			v = v || dfs(i-1, j)
		}
		if b == s2[j-1] {
			v = v || dfs(i, j-1)
		}
		cache[i][j] = &v
		return v
	}

	return dfs(n, m)
}
