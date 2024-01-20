package main

import (
	"strings"
)

func splitWordsBySeparator(words []string, separator byte) []string {
	var ans []string
	for _, word := range words {
		split := strings.Split(word, string([]byte{separator}))
		for _, s := range split {
			if s == "" {
				continue
			}
			ans = append(ans, s)
		}
	}

	return ans
}
