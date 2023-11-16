package main

import "math/big"

// 方法二：组合数学
// 思路与算法
//
// 从左上角到右下角的过程中，我们需要移动 m+n−2 次，其中有 m−1 次向下移动，n−1 次向右移动。
// 因此路径的总数，就等于从 m+n−2 次移动中选择 m−1 次向下移动的方案数，即组合数：
//
// C m+n−2m−1=(m+n−2m−1)=(m+n−2)(m+n−3)⋯n / (m−1)! = [(m+n−2)!]/[(m−1)! (n−1)!]
//
// 因此我们直接计算出这个组合数即可。计算的方法有很多种：
//
// 如果使用的语言有组合数计算的 API，我们可以调用 API 计算；
//
// 如果没有相应的 API，我们可以使用 (m+n−2)(m+n−3)⋯n / (m−1)! 进行计算。
//
// 复杂度分析
//
// 时间复杂度：O(m)。由于我们交换行列的值并不会对答案产生影响，
// 因此我们总可以通过交换 m 和 n 使得 m≤nm ，这样空间复杂度降低至 O(min(m,n))。
//
// 空间复杂度：O(1)。
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/unique-paths/solutions/514311/bu-tong-lu-jing-by-leetcode-solution-hzjf/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode2(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}

// 1.71 MB 的代码示例
func leetcodeMinMemory(m int, n int) int {
	a := m - 1
	if m > n {
		a = n - 1
	}
	b := m + n - 2

	total := big.NewInt(int64(1))

	for i := 0; i < a; i++ {
		total.Mul(total, big.NewInt(int64(b-i)))
	}

	for i := 1; i <= a; i++ {
		total.Div(total, big.NewInt(int64(i)))
	}

	return int(total.Int64())
}

// 1.75 MB 的代码示例
// leetcode2 组合数学的代码实现
func leetcodeMinValue(m int, n int) int {
	ans := 1
	if m > n {
		n, m = m, n //保证m是较少数
	}
	//1,2... m-1   n,n+1,....n+m-2
	for x, y := n, 1; y < m; y++ {
		ans = ans * x / y //	ans = ans * x / y 不可以  ans *= x / y 因为先计算除法可能有小数
		x++
	}
	return ans
}
