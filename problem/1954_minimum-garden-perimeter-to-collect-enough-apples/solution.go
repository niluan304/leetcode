package main

import (
	"math"
	"sort"
)

// ### 复杂度分析
//
// - 时间复杂度：$O(\log m)$，其中 $m$ 表示题目中的 $\textit{neededApples}$，即为二分查找需要的时间。
//
// - 空间复杂度：$O(1)$。
func minimumPerimeter(neededApples int64) int64 {
	ans := sort.Search(1000000, func(i int) bool {
		return int64(i*(i+1)*(2*i+1)/6)*12 >= neededApples
	})

	return int64(ans * 2 * 4)
}

// ![LC1954-c.png](https://pic.leetcode.cn/1703216655-CnqSqn-LC1954-c.png)
//
// 设正方形的**周长**为 $8n$，则其**边长**为 $2n$。
//
// 问题相当于求最小的 $n$，满足
//
// $$
// 2n(n+1)(2n+1)\ge \textit{neededApples}
// $$
//
// 上式变形为
//
// $$
// n(n+1)(n+\dfrac{1}{2})\ge \dfrac{1}{4}\textit{neededApples}
// $$
//
// 设 $m = \left\lfloor \sqrt[3]{\dfrac{1}{4}\textit{neededApples}} \right\rfloor$。
//
// - 由于 $(m-1)m(m-\dfrac{1}{2}) < m^3 \le \dfrac{1}{4}\textit{neededApples}$，所以 $m-1$ 必不满足要求。
// - 由于 $(m+1)(m+2)(m+\dfrac{3}{2}) > (m+1)^3 > \dfrac{1}{4}\textit{neededApples}$，所以 $m+1$ 必满足要求。注意 $m+1>\left\lceil\sqrt[3]{\dfrac{1}{4}\textit{neededApples}}\right\rceil$。
// - $m$ 是否满足要求？计算一下就知道了。
//
// 因此，直接计算出 $n = \left\lfloor \sqrt[3]{\dfrac{1}{4}\textit{neededApples}} \right\rfloor$，如果 $2n(n+1)(2n+1)< \textit{neededApples}$ 则将 $n$ 加一。
//
// > 注：在本题的数据范围下，`cbrt` 算出的整数部分是正确的，不会因为浮点误差导致对 `xxx.999999` 下取整的错误。
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(1)$。开立方有专用的计算函数 `cbrt`，时间可以视作 $\mathcal{O}(1)$。
// - 空间复杂度：$\mathcal{O}(1)$。
//
// #### 相似题目
//
// - [1739. 放置盒子](https://leetcode.cn/problems/building-boxes/)
//
// 欢迎关注 [B站@灵茶山艾府](https://space.bilibili.com/206214)
//
// 更多精彩题解，请看 [往期题解精选（已分类）](https://github.com/EndlessCheng/codeforces-go/blob/master/leetcode/SOLUTIONS.md)
func leetcodeEndless(neededApples int64) int64 {
	n := int64(math.Cbrt(float64(neededApples) / 4))
	if 2*n*(n+1)*(2*n+1) < neededApples {
		n++
	}
	return 8 * n
}
