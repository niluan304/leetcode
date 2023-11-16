package main

import (
	"bytes"
	"regexp"
	"strconv"
	"strings"
)

// leetcode 1 方法一：栈操作
// https://leetcode.cn/problems/decode-string/solutions/264391/zi-fu-chuan-jie-ma-by-leetcode-solution/
func leetcode1(s string) string {
	stk := []string{}
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if cur >= '0' && cur <= '9' {
			digits := getDigits(s, &ptr)
			stk = append(stk, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stk = append(stk, string(cur))
			ptr++
		} else {
			ptr++
			sub := []string{}
			for stk[len(stk)-1] != "[" {
				sub = append(sub, stk[len(stk)-1])
				stk = stk[:len(stk)-1]
			}
			for i := 0; i < len(sub)/2; i++ {
				sub[i], sub[len(sub)-i-1] = sub[len(sub)-i-1], sub[i]
			}
			stk = stk[:len(stk)-1]
			repTime, _ := strconv.Atoi(stk[len(stk)-1])
			stk = stk[:len(stk)-1]
			t := strings.Repeat(getString(sub), repTime)
			stk = append(stk, t)
		}
	}
	return getString(stk)
}

func getDigits(s string, ptr *int) string {
	ret := ""
	for ; s[*ptr] >= '0' && s[*ptr] <= '9'; *ptr++ {
		ret += string(s[*ptr])
	}
	return ret
}

func getString(v []string) string {
	ret := ""
	for _, s := range v {
		ret += s
	}
	return ret
}

var (
	src string
	ptr int
)

// leetcode 2 方法二：递归
// https://leetcode.cn/problems/decode-string/solutions/264391/zi-fu-chuan-jie-ma-by-leetcode-solution/
func leetcode2(s string) string {
	src = s
	ptr = 0
	return getStringLeetcode2()
}

func getStringLeetcode2() string {
	if ptr == len(src) || src[ptr] == ']' {
		return ""
	}
	cur := src[ptr]
	repTime := 1
	ret := ""
	if cur >= '0' && cur <= '9' {
		repTime = getDigitsLeetcode2()
		ptr++
		str := getStringLeetcode2()
		ptr++
		ret = strings.Repeat(str, repTime)
	} else if cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z' {
		ret = string(cur)
		ptr++
	}
	return ret + getStringLeetcode2()
}

func getDigitsLeetcode2() int {
	ret := 0
	for ; src[ptr] >= '0' && src[ptr] <= '9'; ptr++ {
		ret = ret*10 + int(src[ptr]-'0')
	}
	return ret
}

var idx int

// leetcode 3 内存消耗最低解法
func leetcode3(s string) string {
	idx = 0
	return decode(s)
}

func decode(s string) string {
	if idx >= len(s) || s[idx] == ']' {
		return ""
	}

	n := getDigitsLeetcode3(s)
	var res string
	if n != 0 {
		idx++
		ss := decode(s)
		res = strings.Repeat(ss, n)
	} else {
		res += string(s[idx])
	}
	idx++
	return res + decode(s)
}

func getDigitsLeetcode3(s string) int {
	var n int
	for s[idx] >= '0' && s[idx] <= '9' {
		n = n*10 + int(s[idx]-'0')
		idx++
	}
	return n
}

var exp = regexp.MustCompile(`\d+\[[a-z]+]`)

// leetcode 4 正则表达式
func leetcode4(s string) string {
	data := []byte(s)

	for bytes.IndexByte(data, '[') != -1 {
		data = exp.ReplaceAllFunc(data, func(item []byte) []byte {
			var n, idx = 0, 0
			for i, char := range item {
				if char < '0' || char > '9' {
					idx = i + 1
					break
				}
				n = n*10 + int(char-'0')
			}
			return bytes.Repeat(item[idx:len(item)-1], n)
		})
	}

	return string(data)
}
