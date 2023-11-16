package main

func generateParenthesis(n int) []string {
	switch n {
	case 0:
		return []string{}
	case 1:
		return []string{"()"}
	case 2:
		return []string{"(())", "()()"}
	}

	lastSum := generateParenthesis(n - 1)
	var m = map[string]struct{}{}
	var curSum = make([]string, 0, len(lastSum)*3)
	for _, last := range lastSum {
		for i := 0; i <= len(last); i += 2 {
			for j := i; j <= len(last); j += 2 {
				cur := last[:i] + "(" + last[i:j] + ")" + last[j:]
				if _, ok := m[cur]; !ok {
					curSum = append(curSum, cur)
					m[cur] = struct{}{}
				}
			}
		}
	}
	return curSum
}

// leetcode 2 动态规划 https://leetcode.cn/problems/generate-parentheses/solutions/9251/zui-jian-dan-yi-dong-de-dong-tai-gui-hua-bu-lun-da/
func leetcode2(n int) []string {
	return dp(n)[n]
}

func dp(n int) map[int][]string {
	if n == 0 {
		return map[int][]string{0: {""}}
	}
	if n == 1 {
		return map[int][]string{0: {""}, 1: {"()"}}
	}
	lastMap := dp(n - 1)
	var oneRes []string
	for i := 0; i < n; i++ {
		inners := lastMap[i]
		outers := lastMap[n-1-i]
		for _, inner := range inners {
			for _, outer := range outers {
				oneRes = append(oneRes, "("+inner+")"+outer)
			}
		}
	}
	lastMap[n] = oneRes
	return lastMap
}
