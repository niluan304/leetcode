package main

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	}

	var ans []string
	var path []rune
	var numString = []string{
		//"",     // 0 digits[i] 是范围 ['2', '9'] 的一个数字。
		//"",     // 1
		"abc",  // 2
		"def",  // 3
		"ghi",  // 4
		"jkl",  // 5
		"mno",  // 6
		"pqrs", // 7
		"tuv",  // 8
		"wxyz", // 9
	}

	var dfs func(s []byte)
	dfs = func(s []byte) {
		if len(s) == 0 {
			ans = append(ans, string(path))
			return
		}

		for _, b := range numString[s[0]-'2'] {
			path = append(path, b)
			dfs(s[1:])
			path = path[:len(path)-1]
		}
	}

	dfs([]byte(digits))
	return ans
}
