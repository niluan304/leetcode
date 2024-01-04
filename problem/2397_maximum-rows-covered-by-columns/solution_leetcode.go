package main

import "math/bits"

// #### 方法一：二进制枚举
//
// **思路与算法**
//
// 由于题目给定的 $m \times n$ 矩阵 $\textit{mat}$ 的列数 $n$ 满足 $1 \le n \le 12$，所以我们可以用一个整数 $S$ 来表示当前我们选中列的序号集合：从低位到高位，第 $i$ 位为 $1$ 则表示第 $i$ 列被选择，否则表示第 $i$ 列没被选择，同时使用 $mask$ 数组表示矩阵每一行的排列情况，数组中的元素为该行二进制表示的数。然后我们
// 可以通过枚举选择列的全部情况，得到符合题目要求的情况中被覆盖的最大行数：
//
// - 若选择的列的序号状态 $S$ 中 $1$ 的个数不为 $\textit{numSelect}$，则该情况不符合题目要求，跳过该情况。
// - 否则我们计算 $mask$ 数组的每一个值与序号状态的与，若结果等于数组的值则表示该行被覆盖。计算被覆盖的行数，并全局维护该值的最大值即可。
//
// **代码**
// *复杂度分析**
//
// - 时间复杂度：$O(m \times C_m^{\textit{numSelect}})$，其中 $m$ 和 $n$ 分别为矩阵 $\textit{mat}$ 的行数和列数, $\textit{numSelect}$ 为题目要求选择的列的个数。
// - 空间复杂度：$O(m)$，其中 $m$ 为矩阵 $\textit{matrix}$ 的行数。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-rows-covered-by-columns/solutions/2587986/bei-lie-fu-gai-de-zui-duo-xing-shu-by-le-5kb9/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode1(matrix [][]int, numSelect int) int {
	m, n := len(matrix), len(matrix[0])
	mask := make([]int, m)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			mask[i] += matrix[i][j] << (n - j - 1)
		}
	}
	res, limit := 0, 1<<n
	for cur := 1; cur < limit; cur++ {
		if bits.OnesCount(uint(cur)) != numSelect {
			continue
		}
		t := 0
		for j := 0; j < m; j++ {
			if (mask[j] & cur) == mask[j] {
				t++
			}
		}
		res = max(res, t)
	}
	return res
}
