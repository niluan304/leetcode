package main

import (
	"strings"
)

func wordPattern(pattern string, s string) bool {
	var m1 = map[byte]string{}
	var m2 = map[string]byte{}

	words := strings.Split(s, " ")

	if len(words) != len(pattern) {
		return false
	}

	for i, word := range words {
		m1[pattern[i]] = word
		m2[word] = pattern[i]
	}

	if len(m1) != len(m2) {
		return false
	}

	for i, word := range words {
		char := pattern[i]
		if m2[word] > 0 && m2[word] != char || m1[char] != "" && m1[char] != word {
			return false
		}
		m2[word] = char
		m1[char] = word
	}
	return true
}
