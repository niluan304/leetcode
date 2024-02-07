package main

import (
	"strings"
)

func lengthOfLongestSubstring(s string) int {
	var (
		count = 0
		left  = 0
	)

	for right, b := range s {
		idx := strings.IndexRune(s[left:right], b)
		if idx != -1 {
			left += idx + 1
		}

		count = max(count, right+1-left)
	}
	return count
}

func lengthOfLongestSubstring2(s string) int {
	var (
		count = 0
		left  = -1
	)

	var idx = make(map[rune]int, 100)
	for right, b := range s {
		i, ok := idx[b]
		if !ok {
			i = -1
		}

		left = max(i, left)
		count = max(count, right-left)

		idx[b] = right
	}
	return count
}

func lengthOfLongestSubstring3(s string) int {
	var (
		count = 0
		left  = 0
	)

	var set = map[rune]int{}
	for right, b := range s {
		set[b]++
		for set[b] > 1 {
			set[rune(s[left])]--
			left++
		}
		count = max(count, right-left+1)
	}
	return count
}

func lengthOfLongestSubstring4(s string) int {
	var (
		count = 0
		left  = 0
	)

	var set = [256]int{}
	for right := 0; right < len(s); right++ {
		b := s[right]
		set[b]++
		for set[b] > 1 {
			set[s[left]]--
			left++
		}
		count = max(count, right-left+1)

	}
	return count
}
