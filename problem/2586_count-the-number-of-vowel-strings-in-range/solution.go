package main

func vowelStrings(words []string, left int, right int) int {
	var ans = 0
	for _, word := range words[left : right+1] {
		if isVowel(word[0]) && isVowel(word[len(word)-1]) {
			ans++
		}
	}
	return ans
}

func isVowel(c byte) bool {
	switch c {
	case 'a', 'e', 'i', 'o', 'u':
		return true
	}
	return false
}
