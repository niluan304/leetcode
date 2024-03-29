package main

import "sort"

// 为了简化描述，下文把 $\textit{satisfaction}$ 记作 $a$。
//
// ## 提示 1
//
// 枚举做 $k=0,1,2,\cdots,n$ 道菜。
//
// 假设要做 $2$ 道菜，应该选择 $a$ 中的哪两个数呢？
//
// 由于 $a[i]$ 越大，like-time 系数越大，所以应该选择 $a$ 中最大的两个数。
//
// ## 提示 2
//
// 假设 $a$ 中最大的两个数分别是 $4$ 和 $3$，先做哪道菜，后做哪道菜？
//
// - 如果先 $4$ 后 $3$，总和为 $1\cdot 4 + 2\cdot 3 = 10$。
// - 如果先 $3$ 后 $4$，总和为 $1\cdot 3 + 2\cdot 4 = 11$。
//
// 这说明 $a[i]$ 大的菜应该后做。可以根据 [排序不等式](https://baike.baidu.com/item/%E6%8E%92%E5%BA%8F%E4%B8%8D%E7%AD%89%E5%BC%8F/7775728) 得到该结论。
//
// ## 提示 3
//
// 为了方便计算，把 $a$ **从大到小**排序。
//
// 来看看做 $k=1,2,3$ 道菜时，对应的总和 $f(k)$ 是多少。
//
// - $k=1$ 时，总和为 $f(1)=a[0]$。
// - $k=2$ 时，总和为 $f(2)=2\cdot a[0] + a[1]$。
// - $k=3$ 时，总和为 $f(3)=3\cdot a[0] + 2\cdot a[1] + a[2]$。
//
// 为了快速地算出每个 $f(k)$，我们需要找到 $f(k)$ 的**递推式**。观察上面列出的式子，你能找到递推式吗？
//
// 先把 $f(k)$ 的式子列出来：
//
// $$
// f(k)=k\cdot a[0] + (k-1)\cdot a[1] + \cdots + 2\cdot a[k-2] + a[k-1]
// $$
//
// 每一项去掉一个 $a[i]$，得到：
//
// $$
// (k-1)\cdot a[0] + (k-2)\cdot a[1] + \cdots + a[k-2]
// $$
//
// 这正是 $f(k-1)$。
//
// 所以有
//
// $$
// f(k) = f(k-1) + (a[0] + a[1] + \cdots + a[k-1])
// $$
//
// 右边的和式是 $a$ 的前缀和，我们可以一边遍历 $a$，一边把 $a[i]$ 累加到一个变量 $s$ 中。这样就可以 $\mathcal{O}(1)$ 地从 $f(k-1)$ 递推得到 $f(k)$ 了。
//
// 答案为 $f(0), f(1), f(2),\cdots, f(n)$ 中的最大值。
//
// ## 实现细节
//
// 想一想 $s$ 是怎么变化的。由于数组是从大到小排序的，（一般地）会先遇到正数，再遇到负数，所以（一般地）$s$ 会先变大，再变小。
//
// 如果 $s\le 0$，那么后面的 $a[i]$ 必然都是负数，我们不可能得到更大的 $f(k)$，退出循环。
//
// 代码实现时，可以只用一个变量表示 $f$。由于在退出循环之前 $s$ 都是大于 $0$ 的，所以 $f(k) > f(k-1)$，因此退出循环时的 $f$ 就是最终答案。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $a$ 的长度。瓶颈在排序上。
// - 空间复杂度：$\mathcal{O}(1)$。忽略排序的栈开销。
func leetcodeEndlesscheng(satisfaction []int) int {
	sort.Sort(sort.Reverse(sort.IntSlice(satisfaction)))
	f := 0 // f(0) = 0
	s := 0 // satisfaction 的前缀和
	for _, x := range satisfaction {
		s += x
		if s <= 0 { // 后面不可能找到更大的 f(k)
			break
		}
		f += s // f(k) = f(k-1) + s
	}
	return f
}
