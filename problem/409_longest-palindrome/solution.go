package main

func longestPalindrome(s string) int {
	slice := make([]byte, 26*2)
	longest := 0
	for i := 0; i < len(s); i++ {
		char := s[i]
		if char > 'Z' {
			char = char - 'a' + 'Z'
		}

		slice[char-'A']++

		if slice[char-'A'] == 2 {
			longest += 2
			slice[char-'A'] = 0
		}
	}

	for _, n := range slice {
		if n == 1 {
			longest++
			break
		}
	}

	return longest
}

func longestPalindrome2(s string) int {
	slice := make([]byte, 26*2)
	n, longest := len(s), 0
	for i := 0; i < n; i++ {
		char := s[i]
		if char > 'Z' {
			char = char - 'a' + 'Z'
		}

		slice[char-'A']++

		if slice[char-'A'] == 2 {
			longest += 2
			slice[char-'A'] = 0
		}
	}

	if longest < n {
		longest++
	}
	return longest
}

func longestPalindrome3(s string) int {
	slice := make([]byte, 'z'-'A'+1)
	n, longest := len(s), 0
	for i := 0; i < n; i++ {
		char := s[i]
		slice[char-'A']++
		if slice[char-'A'] == 2 {
			longest += 2
			slice[char-'A'] = 0
		}
	}

	if longest < n {
		longest++
	}
	return longest
}

func leetcode1(s string) int {
	cnt := make([]int, 128)
	for _, c := range s {
		cnt[c]++
	}
	oddCnt := 0
	for _, e := range cnt {
		oddCnt += e % 2
	}
	n := len(s)
	if oddCnt == 0 {
		return n
	}
	return n - oddCnt + 1
}
