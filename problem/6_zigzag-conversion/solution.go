package main

import (
	"bytes"
)

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	t := (numRows - 1) * 2
	zStr := make([][]byte, numRows)
	for i := 0; i < len(s); i++ {
		row := i % t
		if row >= numRows {
			row = numRows - (row - numRows) - 2
		}

		zStr[row] = append(zStr[row], s[i])
	}

	for i := 1; i < numRows; i++ {
		zStr[0] = append(zStr[0], zStr[i]...)
	}
	return string(zStr[0])
}

func convert2(s string, numRows int) string {
	n := len(s)
	if numRows < 2 || n <= numRows {
		return s
	}

	zStr := make([]byte, 0, len(s))
	t := (numRows - 1) * 2

	for i := 0; i < numRows; i++ {
		flag := t - 2*i
		for j := 0; i+j < n; j += t {
			k := i + j
			zStr = append(zStr, s[k])

			k += flag
			if i > 0 && i < numRows-1 && k < n {
				zStr = append(zStr, s[k])
			}
		}
	}
	return string(zStr)
}

func convert3(s string, numRows int) string {
	n := len(s)
	if numRows < 2 || n <= numRows {
		return s
	}

	zStr := make([]byte, 0, n)
	end := numRows - 1
	t := end * 2

	for j := 0; j < n; j += t {
		zStr = append(zStr, s[j])
	}
	for i := 1; i < end; i++ {
		flag := t - 2*i
		for j := 0; i+j < n; j += t {
			k := i + j
			zStr = append(zStr, s[k])

			k += flag
			if k < n {
				zStr = append(zStr, s[k])
			}
		}
	}
	for j := end; j < n; j += t {
		zStr = append(zStr, s[j])
	}

	return string(zStr)
}

// leetcode 1 0.048MB 内存
func leetcode1(s string, numRows int) string {
	if numRows < 2 || len(s) <= numRows {
		return s
	}
	var res []uint8

	for i := 0; i < numRows; i++ {
		down := true
		for j := i; j < len(s); {
			res = append(res, s[j])
			if i == 0 || i == numRows-1 {
				j += (numRows - 1) * 2
			} else {
				if down {
					j += (numRows - i - 1) * 2
				} else {
					j += i * 2
				}
				down = !down
			}
		}
	}

	return string(res)
}

func leetcode2(s string, numRows int) string {
	r := numRows
	if r == 1 || r >= len(s) {
		return s
	}
	mat := make([][]byte, r)
	t, x := r*2-2, 0
	for i, ch := range s {
		mat[x] = append(mat[x], byte(ch))
		if i%t < r-1 {
			x++
		} else {
			x--
		}
	}
	return string(bytes.Join(mat, nil))
}

func leetcode3(s string, numRows int) string {
	n, r := len(s), numRows
	if r == 1 || r >= n {
		return s
	}
	t := r*2 - 2
	ans := make([]byte, 0, n)
	for i := 0; i < r; i++ { // 枚举矩阵的行
		for j := 0; j+i < n; j += t { // 枚举每个周期的起始下标
			ans = append(ans, s[j+i]) // 当前周期的第一个字符
			if 0 < i && i < r-1 && j+t-i < n {
				ans = append(ans, s[j+t-i]) // 当前周期的第二个字符
			}
		}
	}
	return string(ans)
}

func leetcode4(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	var str = make([][]byte, numRows)
	i, flag := 0, -1
	for j := 0; j < len(s); j++ {
		str[i] = append(str[i], s[j])
		if i == 0 || i == numRows-1 {
			flag = -flag
		}

		i += flag
	}

	for k := 1; k < numRows; k++ {
		str[0] = append(str[0], str[k]...)
	}

	return string(str[0])
}
