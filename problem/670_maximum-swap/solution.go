package main

import (
	"strconv"
)

func maximumSwap(num int) int {
	var s = []byte(strconv.Itoa(num))
	var n = len(s)

	var maxSuffix = make([]int, n)
	maxSuffix[n-1] = n - 1
	for i := n - 2; i >= 0; i-- {
		maxSuffix[i] = i
		if j := maxSuffix[i+1]; s[i] <= s[j] {
			maxSuffix[i] = j
		}
	}

	for i, j := range maxSuffix {
		if s[i] != s[j] {
			s[i], s[j] = s[j], s[i]
			break
		}
	}

	var ans = 0
	for _, b := range s {
		ans = ans*10 + int(b-'0')
	}
	return ans
}

// #### 方法一：直接遍历
//
// 由于对于整数 $\textit{num}$ 的十进制数字位长最长为 $8$ 位，任意两个数字交换一次最多有 $28$ 种不同的交换方法，因此我们可以尝试遍历所有可能的数字交换方法即可，并找到交换后的最大数字即可。
// + 我们将数字存储为长度为 $n$ 的列表，其中 $n$ 为整数 $\textit{num}$ 的十进制位数的长度。对于位置为 $\text{(i, j)}$ 的每个候选交换，我们交换数字并记录组成的新数字是否大于当前答案；
// + 对于前导零的问题，我们也不需要特殊处理。
// + 由于数字只有 $8$ 位，所以我们不必考虑交换后溢出的风险；
func maximumSwap2(num int) int {
	ans := num
	s := []byte(strconv.Itoa(num))
	for i := range s {
		for j := range s[:i] {
			s[i], s[j] = s[j], s[i]
			v, _ := strconv.Atoi(string(s))
			ans = max(ans, v)
			s[i], s[j] = s[j], s[i]
		}
	}
	return ans
}
