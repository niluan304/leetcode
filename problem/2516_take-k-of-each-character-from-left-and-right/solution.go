package main

func takeCharacters(s string, k int) int {
	var (
		set = [3]int{-k, -k, -k}
		n   = len(s)
	)

	for _, b := range s {
		set[b-'a']++
	}

	if set[0] < 0 || set[1] < 0 || set[2] < 0 {
		return -1
	}

	if set[0] == 0 && set[1] == 0 && set[2] == 0 {
		return n
	}

	var (
		left  = 0
		count = [3]int{}
		ans   = 0
	)

	for right, b := range s {
		count[b-'a']++

		for count[0] > set[0] || count[1] > set[1] || count[2] > set[2] {
			count[s[left]-'a']--
			left++
		}

		ans = max(ans, right-left+1)
	}

	return n - ans
}
