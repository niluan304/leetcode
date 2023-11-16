package main

import (
	"fmt"
	"strconv"
)

func addStrings(num1 string, num2 string) string {
	var sum = []byte{'0'}

	n1, n2 := len(num1), len(num2)
	for i := 0; i < n1 || i < n2; i++ {
		num := sum[len(sum)-1]
		if i < n1 {
			num += (num1[n1-i-1]) - '0'
		}
		if i < n2 {
			num += (num2[n2-i-1]) - '0'
		}

		var flag byte = '0'
		if num > 9+'0' {
			num -= 10
			flag += 1
		}
		sum[len(sum)-1] = num
		sum = append(sum, flag)
	}

	if sum[len(sum)-1] == '0' {
		sum = sum[:len(sum)-1]
	}
	n := len(sum)
	for i := 0; i < n/2; i++ {
		sum[i], sum[n-i-1] = sum[n-i-1], sum[i]
	}
	return string(sum)
}

func addStrings2(num1 string, num2 string) string {
	var sum = "0"
	n1, n2 := len(num1), len(num2)

	for i, j := 0, 0; i < n1 || j < n2; {
		num := sum[0] - '0'
		if i < n1 {
			num += (num1[n1-i-1]) - '0'
		}
		if j < n2 {
			num += (num2[n2-j-1]) - '0'
		}
		i++
		j++
		sum = fmt.Sprintf("%02d%s", num, sum[1:]) // 每次循环都会重置sum的排序, 非常耗时...
	}

	if sum[0] == '0' {
		sum = sum[1:]
	}
	return sum
}

func addStrings3(num1 string, num2 string) string {
	n1, n2, size := len(num1), len(num2), len(num2)

	if size < n1 {
		size = n1
	}

	var sum = append(make([]byte, 0, size+2), '0')

	for i := 0; i < n1 || i < n2; i++ {
		num := sum[len(sum)-1]
		if i < n1 {
			num += (num1[n1-i-1]) - '0'
		}
		if i < n2 {
			num += (num2[n2-i-1]) - '0'
		}

		var flag byte = '0'
		if num > 9+'0' {
			num -= 10
			flag += 1
		}
		sum[len(sum)-1] = num
		sum = append(sum, flag)
	}

	if sum[len(sum)-1] == '0' {
		sum = sum[:len(sum)-1]
	}
	n := len(sum)
	for i := 0; i < n/2; i++ {
		sum[i], sum[n-i-1] = sum[n-i-1], sum[i]
	}
	return string(sum)
}

func leetcode1(num1 string, num2 string) string {
	add := 0
	ans := ""
	for i, j := len(num1)-1, len(num2)-1; i >= 0 || j >= 0 || add != 0; i, j = i-1, j-1 {
		var x, y int
		if i >= 0 {
			x = int(num1[i] - '0')
		}
		if j >= 0 {
			y = int(num2[j] - '0')
		}
		result := x + y + add
		ans = strconv.Itoa(result%10) + ans
		add = result / 10
	}
	return ans
}

func leetcode2(num1 string, num2 string) string {
	i, j := len(num1)-1, len(num2)-1
	add := 0
	var res string
	for i >= 0 || j >= 0 || add > 0 {
		sum := 0
		if i >= 0 {
			sum += int(num1[i] - '0')
			i--
		}
		if j >= 0 {
			sum += int(num2[j] - '0')
			j--
		}
		sum += add
		add = sum / 10
		res = strconv.Itoa(sum%10) + res
	}
	return res
}
