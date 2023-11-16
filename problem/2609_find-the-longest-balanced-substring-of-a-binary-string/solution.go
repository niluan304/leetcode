package main

import "strings"

func findTheLongestBalancedSubstring(s string) int {
	var ans = 0
	var zero, one = 0, 0
	for _, b := range s {
		switch b {
		case '0':
			if one > 0 {
				one = 0
				zero = 0
			}
			zero++
		case '1':
			if zero > 0 {
				one++
				ans = max(ans, min(zero, one)*2)
			}
		}
	}

	return ans
}

func findTheLongestBalancedSubstring2(s string) int {
	sub := ""
	for strings.Contains(s, sub) {
		sub = "0" + sub + "1"
	}
	return len(sub) - 2
}
