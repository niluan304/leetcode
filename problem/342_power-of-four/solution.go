package main

import (
	"strconv"
	"strings"
)

func isPowerOfFour(n int) bool {
	if n <= 0 {
		return false
	}

	const mask = 0b10101010101010101010101010101010
	return n&(n-1) == 0 && n&mask == 0
	return n&-n == n && n&mask == 0 // 等价
}
func isPowerOfFour2(n int) bool {
	switch n {
	case 1, 4, 16, 64, 256, 1024, 4096, 16384, 65536, 262144, 1048576, 4194304, 16777216, 67108864, 268435456, 1073741824:
		return true
	default:
		return false
	}
}

func isPowerOfFour3(n int) bool {
	if n <= 0 {
		return false
	}
	s := strconv.FormatInt(int64(n), 2)
	i := strings.LastIndex(s, "1")
	return i == 0 && len(s)%2 != 0
}
