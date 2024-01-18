package main

// 在方法一中，我们仅需要统计当前窗口内的最大值与最小值，因此我们也可以分别使用两个单调队列解决本题。
//
// 在实际代码中，我们使用一个单调递增的队列 queMin 维护最小值，一个单调递减的队列 queMax 维护最大值。
// 这样我们只需要计算两个队列的队首的差值，即可知道当前窗口是否满足条件。
//
// *复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 是数组长度。我们最多遍历该数组两次，两个单调队列入队出队次数也均为 $O(n)$。
//
// - 空间复杂度：$O(n)$，其中 $n$ 是数组长度。最坏情况下单调队列将和原数组等大。
func leetcode2(nums []int) int64 {
	const limit = 2
	var minQ, maxQ []int
	left, ans := 0, 0
	for right, v := range nums {
		for len(minQ) > 0 && minQ[len(minQ)-1] > v {
			minQ = minQ[:len(minQ)-1]
		}
		minQ = append(minQ, v)
		for len(maxQ) > 0 && maxQ[len(maxQ)-1] < v {
			maxQ = maxQ[:len(maxQ)-1]
		}
		maxQ = append(maxQ, v)
		for len(minQ) > 0 && len(maxQ) > 0 && maxQ[0]-minQ[0] > limit {
			if nums[left] == minQ[0] {
				minQ = minQ[1:]
			}
			if nums[left] == maxQ[0] {
				maxQ = maxQ[1:]
			}
			left++
		}
		ans += right - left + 1
	}
	return int64(ans)
}

// 代码框架是 [滑动窗口（双指针）](https://www.bilibili.com/video/BV1hd4y1r7Gq/)。在遍历数组的同时，维护窗口内的数字。
//
// 由于绝对差至多为 $2$，所以用平衡树或者哈希表维护都行，反正至多维护 $3$ 个数，添加删除可以视作是 $\mathcal{O}(1)$ 的。（如果用哈希表，还需记录数字的出现次数。）
//
// 如果窗口内的最大值与最小值的差大于 $2$，就不断移动左端点 $\textit{left}$，减少窗口内的数字。
//
// 最后
//
// $$
// [\textit{left},\textit{right}],[\textit{left}+1,\textit{right}],\cdots,[\textit{right},\textit{right}]
// $$
//
// 这一共 $\textit{right}-\textit{left}+1$ 个子数组都是合法的，加入答案。
//
// #### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(1)$。注意至多维护 $3$ 个数，仅用到常量额外空间。
//
// #### 相似题目
//
// - [1438. 绝对差不超过限制的最长连续子数组](https://leetcode.cn/problems/longest-continuous-subarray-with-absolute-diff-less-than-or-equal-to-limit/)
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/continuous-subarrays/solutions/2327219/shuang-zhi-zhen-ping-heng-shu-ha-xi-biao-4frl/
func endlessCheng(a []int) (ans int64) {
	cnt := map[int]int{}
	left := 0
	for right, x := range a {
		cnt[x]++
		for {
			mx, mn := x, x
			for k := range cnt {
				mx = max(mx, k)
				mn = min(mn, k)
			}
			if mx-mn <= 2 {
				break
			}
			y := a[left]
			if cnt[y]--; cnt[y] == 0 {
				delete(cnt, y)
			}
			left++
		}
		ans += int64(right - left + 1)
	}
	return
}
