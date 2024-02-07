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
		mx    = 0
		mn    = math.MaxInt32
		count = make(map[rune]int, len(s))
	)
	for _, b := range s {
		count[b]++
	}

	for _, v := range count {
		if mx < v {
			mx = v
		}
		if mn > v {
			mn = v
		}
	}

	return mx - mn
}

func beautySum3(s string) int {
	var sum = 0
	for i := range s {
		var (
			count = make(map[rune]int, len(s))
			mx    = 1
		)
		for _, char := range s[i:] {
			count[char]++
			if count[char] > mx {
				mx = count[char]
			}

			var mn = 500 // 1 <= s.length <= 500
			for _, cnt := range count {
				if cnt < mn {
					mn = cnt
				}
			}

			sum += mx - mn
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
