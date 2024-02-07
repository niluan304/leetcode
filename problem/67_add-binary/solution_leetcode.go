package main

import (
	"strconv"
)

// leetcode 1 模拟
func leetcode1(a string, b string) string {
	ans := ""
	carry := 0
	lenA, lenB := len(a), len(b)
	n := max(lenA, lenB)

	for i := 0; i < n; i++ {
		if i < lenA {
			carry += int(a[lenA-i-1] - '0')
		}
		if i < lenB {
			carry += int(b[lenB-i-1] - '0')
		}
		ans = strconv.Itoa(carry%2) + ans
		carry /= 2
	}
	if carry > 0 {
		ans = "1" + ans
	}
	return ans
}

// leetcode 2 0ms示例
func leetcode2(a string, b string) string {
	arr := []byte{}
	add := byte(0)
	i := 0
	for i < len(a) || i < len(b) {
		t := add
		if i < len(a) {
			t += a[len(a)-i-1] - '0'
		}
		if i < len(b) {
			t += b[len(b)-i-1] - '0'
		}

		arr = append(arr, t&byte(1)+'0')
		add = t / 2
		i++
	}
	if add > 0 {
		arr = append(arr, byte(1)+'0')
	}
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return string(arr[:])
}
