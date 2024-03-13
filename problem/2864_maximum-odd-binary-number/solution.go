package main

import "strings"

func maximumOddBinaryNumber(s string) string {
	ones := strings.Count(s, "1")
	return strings.Repeat("1", ones-1) + strings.Repeat("0", len(s)-ones) + "1"
}
