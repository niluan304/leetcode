package main

import "bytes"

// 差分数组
// 时间复杂度：O(m+n), m=len(shifts), n=len(s)
// 空间复杂度：O(n)
func shiftingLetters(s string, shifts [][]int) string {
	var n = len(s)
	var diff = make([]int, n+1) // 差分数组
	for _, shift := range shifts {
		left, right, x := shift[0], shift[1], shift[2]
		if x == 0 {
			x = -1
		}
		diff[left] += x
		diff[right+1] -= x
	}

	// 在差分数组上还原数组
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	var ans = make([]byte, n)
	for i := 0; i < n; i++ {
		move := diff[i] % 26
		if move < 0 {
			move += 26
		}
		v := int(s[i]) + move
		if v > 'z' {
			v = v - 'z' + 'a' - 1
		}
		ans[i] = byte(v)
	}
	return string(ans)
}

// 差分数组
// 时间复杂度：O(m+n), m=len(shifts), n=len(s)
// 空间复杂度：O(n)
func shiftingLetters2(s string, shifts [][]int) string {
	var n = len(s)
	var diff = make([]int, n+1) // 差分数组
	for _, shift := range shifts {
		left, right, x := shift[0], shift[1], shift[2]
		if x == 0 {
			x = -1
		}
		diff[left] += x
		diff[right+1] -= x
	}

	// 在差分数组上还原数组
	for i := 1; i < n; i++ {
		diff[i] += diff[i-1]
	}

	var buff = new(bytes.Buffer)
	for i := 0; i < n; i++ {
		offset := int(s[i]-'a') + diff[i] // 偏移量
		offset %= 26
		if offset < 0 {
			offset += 26
		}
		buff.WriteByte(byte(offset + 'a'))
	}

	return buff.String()
}
