package main

func maxProduct(words []string) int {
	var n = len(words)

	var nums = make([]int, n)
	for i, word := range words {
		for _, c := range word {
			nums[i] |= 1 << (c - 'a')
		}
	}

	var ans = 0
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[i]&nums[j] == 0 {
				ans = max(ans, len(words[i])*len(words[j]))
			}
		}
	}
	return ans
}
