package main

func lengthOfLastWord(s string) (ans int) {
	index := len(s) - 1
	for s[index] == ' ' {
		index--
	}
	for index >= 0 && s[index] != ' ' {
		ans++
		index--
	}
	return
}
