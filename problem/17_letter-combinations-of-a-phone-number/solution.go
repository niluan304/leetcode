package main

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}

	var numbers = []string{"abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	var ans []string
	var path []rune
	var dfs func(i int)
	dfs = func(i int) {
		if i == len(digits) {
			ans = append(ans, string(path))
			return
		}

		for _, char := range numbers[digits[i]-'2'] {
			path = append(path, char)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}
	dfs(0)
	return ans
}
