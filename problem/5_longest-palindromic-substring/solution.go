package main

func longestPalindrome(s string) string {
	n := len(s)
	longest := ""
	for i := 0; i < n; i++ {
		for j := i + 1 + len(longest); j < n+1; j++ {
			str := s[i:j]
			if IsPalindrome(str) {
				longest = str
			}
		}
	}

	return longest
}

func longestPalindrome2(s string) string {

	// TODO 写出更高效的算法
	panic("写出更高效的算法:" + s)

	return s
}

func IsPalindrome(s string) bool {
	n := len(s)

	for i := 0; i < n/2; i++ {
		if s[i] != s[n-1-i] {
			return false
		}
	}

	return true
}
