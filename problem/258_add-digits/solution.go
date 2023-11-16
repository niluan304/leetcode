package main

import (
	"strconv"
	"strings"
)

func addDigits(num int) int {
	for {
		ans := 0
		for _, s := range strings.Split(strconv.Itoa(num), "") {
			ans += int(s[0] - '0')
		}

		if ans < 10 {
			return ans
		}

		num = ans
	}
}
