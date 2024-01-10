package main

import "strings"

func minLength(s string) int {
	for strings.Contains(s, "AB") || strings.Contains(s, "CD") {
		s = strings.ReplaceAll(s, "AB", "")
		s = strings.ReplaceAll(s, "CD", "")
	}
	return len(s)
}

func minLength2(s string) int {
	var stack []byte
	for i := 0; i < len(s); i++ {
		stack = append(stack, s[i])

		for {
			n := max(0, len(stack)-2)
			tail := string(stack[n:])
			if tail != "AB" && tail != "CD" {
				break
			}
			stack = stack[:n]
		}
	}
	return len(stack)
}
