package main

import (
	"math"
)

func beautySum(s string) int {
	if len(s) < 3 {
		return 0
	}

	var sum = 0
	for i := 0; i <= len(s)-3; i++ {
		for j := i + 3; j <= len(s); j++ {
			sum += beauty(s[i:j])
		}
	}

	return sum
}

func beauty(s string) int {
	var (
		max   = 0
		min   = math.MaxInt32
		count = make(map[rune]int, len(s))
	)
	for _, b := range s {
		count[b]++
	}

	for _, v := range count {
		if max < v {
			max = v
		}
		if min > v {
			min = v
		}
	}

	return max - min
}

func beautySum3(s string) int {
	var sum = 0
	for i := range s {
		var (
			count = make(map[rune]int, len(s))
			max   = 1
		)
		for _, char := range s[i:] {
			count[char]++
			if count[char] > max {
				max = count[char]
			}

			var min = 500 // 1 <= s.length <= 500
			for _, cnt := range count {
				if cnt < min {
					min = cnt
				}
			}

			sum += max - min
		}
	}
	return sum
}

func leetcode(s string) (ans int) {
	for i := range s {
		cnt := [26]int{}
		maxFreq := 0
		for _, c := range s[i:] {
			cnt[c-'a']++
			maxFreq = max(maxFreq, cnt[c-'a'])
			minFreq := len(s)
			for _, c := range cnt {
				if c > 0 {
					minFreq = min(minFreq, c)
				}
			}
			ans += maxFreq - minFreq
		}
	}
	return
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
