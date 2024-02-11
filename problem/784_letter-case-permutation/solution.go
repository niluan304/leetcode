package main

// 子集型回溯 模板1
//
// 从输入的角度 (选与不选)
//
// 1.当前操作？
// 变不变这个元素
// 2.子问题？
// 构造字符串>=i的部分
// 3.下一个子问题？
// 构造字符串>=i+1的部分
//
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func letterCasePermutation(s string) []string {
	var ans []string
	n, path := len(s), []byte(s)

	var dfs func(i int)
	dfs = func(i int) {
		if i == n {
			ans = append(ans, string(path))
			return
		}
		x := s[i]
		dfs(i + 1)
		if '0' <= x && x <= '9' {
			return
		}

		// 小技巧
		//'a' ^ 32 = 'A'
		//'A' ^ 32 = 'a'
		path[i] ^= 32 // 转换大小写
		dfs(i + 1)
	}
	dfs(0)
	return ans
}

// 子集型回溯 模板2
//
// 从答案的角度
//
// 1.当前操作？
// 构造一个下标 j>=i 的字符串
// 2.子问题？
// 从下标 >= i 的字符串中构造子集
// 3.下一个子问题？
// 从下表 >= j+1 的数字中构造子集
//
// - 时间复杂度：$\mathcal{O}(n \cdot 2^n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func letterCasePermutation2(s string) (ans []string) {
	n, path := len(s), []byte(s)

	var dfs func(int)
	dfs = func(i int) {
		ans = append(ans, string(path))
		if i == n {
			return // 可以省去这个判断
		}
		for j := i; j < n; j++ {
			x := path[j]
			if '0' <= x && x <= '9' {
				continue
			}

			// 小技巧
			//'a' ^ 32 = 'A'
			//'A' ^ 32 = 'a'
			path[j] ^= 32 // 改变大小写
			dfs(j + 1)
			path[j] ^= 32 // 恢复
		}
	}
	dfs(0)
	return
}
