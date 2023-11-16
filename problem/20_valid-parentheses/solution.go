package main

import "strings"

func isValid(s string) bool {
	var stack = make([]byte, 0, len(s))
	for i := range s {
		if len(stack) != 0 {
			switch s[i] {
			case ')':
				if stack[len(stack)-1] == '(' {
					stack = stack[:len(stack)-1]
					continue
				}
			case ']':
				if stack[len(stack)-1] == '[' {
					stack = stack[:len(stack)-1]
					continue
				}
			case '}':
				if stack[len(stack)-1] == '{' {
					stack = stack[:len(stack)-1]
					continue
				}
			}
		}

		stack = append(stack, s[i])
	}

	return len(stack) == 0
}

func isValid2(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var pattern = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}
	var stack = make([]byte, 0, len(s))
	for i := range s {
		v, ok := pattern[s[i]]
		if !ok {
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != v {
			return false
		}
		stack = stack[:len(stack)-1]

	}

	return len(stack) == 0
}

func isValid3(s string) bool {
	if len(s)%2 != 0 {
		return false
	}
	var pattern = map[byte]byte{
		')': '(',
		']': '[',
		'}': '{',
	}

	var stack = make([]byte, 0, len(s))
	for i := range s {
		switch s[i] {
		case '(', '[', '{':
			stack = append(stack, s[i])
			continue
		}

		if len(stack) == 0 || stack[len(stack)-1] != pattern[s[i]] {
			return false
		}
		stack = stack[:len(stack)-1]
	}

	return len(stack) == 0
}

func isValid4(s string) bool {
	if len(s)%2 != 0 {
		return false
	}

	n := len(s) / 2
	for i := 0; i < n; i++ {
		s = strings.ReplaceAll(s, "()", "")
		s = strings.ReplaceAll(s, "{}", "")
		s = strings.ReplaceAll(s, "[]", "")
	}

	return len(s) == 0
}
