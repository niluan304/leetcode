package main

// ### 前置知识：同向双指针
//
// 见 [同向双指针【基础算法精讲 01】](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。
//
// > 注：我一般把窗口大小不固定的叫做**双指针**，窗口大小固定的叫做**滑动窗口**。
//
// ### 思路
//
// 我们可以强制让第二条线段的右端点恰好落在奖品上，设**第二条**线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，我们需要知道在 $\textit{prizePositions}[\textit{left}]$ 左侧的第一条线段最多可以覆盖多少个奖品。
//
// 那么，先想想只有一条线段要怎么做。
//
// 使用双指针，设线段右端点在 $\textit{prizePositions}[\textit{right}]$ 时，左端点最远覆盖了 $\textit{prizePositions}[\textit{left}]$，那么当前覆盖的奖品个数为 $\textit{right} - \textit{left} + 1$。
//
// 同时，用一个数组 $\textit{pre}[\textit{right}+1]$ 记录线段右端点**不超过** $\textit{prizePositions}[\textit{right}]$ 时最多可以覆盖多少个奖品。下标错开一位是为了方便下面计算。
//
// 初始 $\textit{pre}[0]=0$。根据 $\textit{pre}$ 的定义，有
//
// $$
// \textit{pre}[\textit{right}+1] = \max(\textit{pre}[\textit{right}],\textit{right} - \textit{left} + 1)
// $$
//
// 回到第二条线段的计算，根据开头说的，此时最多可以覆盖的奖品数为
//
// $$
// \textit{right}-\textit{left}+1+\textit{pre}[\textit{left}]
// $$
//
// 这里 $\textit{pre}[\textit{left}]$ 表示**第一条**线段右端点**不超过** $\textit{prizePositions}[\textit{left}-1]$ 时最多可以覆盖多少个奖品。
//
// 遍历过程中取上式的最大值，即为答案。
//
// 由于我们遍历了所有的奖品作为第二条线段的右端点，且通过 $\textit{pre}[\textit{left}]$ 保证第一条线段与第二条线段没有任何交点，且第一条线段覆盖了第二条线段左侧的最多奖品。那么这样遍历后，算出的答案就一定是所有情况中的最大值。
//
// 如果脑中没有一幅直观的图像，可以看看 [视频讲解【双周赛 97】](https://www.bilibili.com/video/BV1rM4y1X7z9/)的第三题。
//
// ### 复杂度分析
//
// - 时间复杂度：$O(n)$，其中 $n$ 为 $\textit{prizePositions}$ 的长度。虽然写了个二重循环，但是内层循环中对 $\textit{left}$ 加一的**总**执行次数不会超过 $n$ 次，所以总的时间复杂度为 $O(n)$。
// - 空间复杂度：$O(n)$。
func endlessCheng(positions []int, k int) int {
	var n = len(positions)
	var prefix = make([]int, n+1)

	var ans, left = 0, 0
	for right, pos := range positions {
		for pos-positions[left] > k { // 通过滑动窗口，避免了二分里的重复操作
			left++
		}

		cnt := right - left + 1
		ans = max(ans, cnt+prefix[left])          // 第二条线段 + 前缀最大的线段
		prefix[right+1] = max(prefix[right], cnt) // 更新前缀最大值
	}
	return ans
}
