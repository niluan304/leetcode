package main

func maximumNumberOfStringPairs(words []string) int {
	var count = map[string]int{}
	var ans = 0

	for _, word := range words {
		key := reverse(word)
		if count[key] > 0 {
			ans++
			count[key]--
		} else {
			count[word]++
		}
	}

	return ans
}

func reverse(s string) string {
	var n = len(s)
	var res = make([]byte, n)
	for i := 0; i < n; i++ {
		res[n-1-i] = s[i]
	}
	return string(res)
}
