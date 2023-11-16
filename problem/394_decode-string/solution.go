package main

import (
	"bytes"
)

type Stack struct {
	n     int
	value []byte
}

func decodeString(s string) string {
	var n = 0
	var stacks = make([]Stack, 1)
	for _, char := range s {
		j := len(stacks) - 1

		switch char {
		case ']':
			stacks[j-1].value = append(stacks[j-1].value, bytes.Repeat(stacks[j].value, stacks[j].n)...)
			stacks = stacks[:j]
			n = 0
		case '[':
			stacks = append(stacks, Stack{n: n})
			n = 0
		default:
			if char >= '0' && char <= '9' {
				n = n*10 + int(char-'0')
			} else {
				stacks[j].value = append(stacks[j].value, byte(char))
			}
		}
	}

	return string(stacks[0].value)
}
