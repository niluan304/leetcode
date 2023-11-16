package main

import "math/bits"

// ![](https://pic.leetcode.cn/1697089221-bDdgkc-lc260-c.png)
//
// 代码实现时，需要找到异或和中的某个值为 $1$ 的比特位。
//
// 一种方式是计算 lowbit，只保留二进制最低位的 $1$，举例如下：
//
// ```cpp
//
//	 s = 101100
//	~s = 010011
//
// (~s)+1 = 010100 // 根据补码的定义，这就是 -s   最低 1 左侧取反，右侧不变
// s & -s = 000100 // lowbit
// ```
//
// 该技巧收录在 [从集合论到位运算，常见位运算技巧分类总结！](https://leetcode.cn/circle/discuss/CaOJ45/)
//
// 也可以通过计算 $\textit{xorAll}$ **尾零的个数**，直接取得 $\textit{nums}[i]$ 在该比特位上的值，从而避免逻辑判断。
//
// > 注：如果没有计算尾零个数的库函数，可以改为计算二进制的长度减一。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(1)$。仅用到若干额外变量。
//
// ## 相似题目
//
// - [136. 只出现一次的数字](https://leetcode.cn/problems/single-number/)
// - [137. 只出现一次的数字 II](https://leetcode.cn/problems/single-number-ii/)
//
// 欢迎关注 [B站@灵茶山艾府](https://b23.tv/JMcHRRp)
//
// 更多精彩题解，请看 [往期题解精选（按 tag 分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
func leetcode2(nums []int) []int {
	xorAll := 0
	for _, x := range nums {
		xorAll ^= x
	}
	tz := bits.TrailingZeros(uint(xorAll))
	ans := make([]int, 2)
	for _, x := range nums {
		ans[x>>tz&1] ^= x
	}
	return ans
}
