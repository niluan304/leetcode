package main

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func alternatingSubarray(nums []int) int {

	var n = len(nums)
	var ans = 0 // math.MaxInt32 // math.MinInt32

	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if nums[j]-nums[i] != (j-i)%2 {
				break
			}
			ans = max(ans, j-i+1)
		}
	}

	if ans <= 1 {
		return -1
	}

	return ans
}

// **适用场景**：按照题目要求，数组会被分割成若干组，且每一组的判断/处理逻辑是一样的。
//
// **核心思想**：
//
// - 外层循环负责遍历组之前的准备工作（记录开始位置），和遍历组之后的统计工作（更新答案最大值）。
// - 内层循环负责遍历组，找出这一组最远在哪结束。
//
// 这个写法的好处是，各个逻辑块分工明确，也不需要特判最后一组。以我的经验，这个写法是所有写法中最不容易出 bug 的，推荐大家记住。
//
// 对于本题来说，在内层循环时，假设这一组的第一个数是 $3$，那么这一组的数字必须形如 $3,4,3,4,\cdots$。设第一个数字的下标是 $i_0$，我们可以根据当前数字的下标 $i$ 分类讨论：如果 $i$ 和 $i_0$ 的奇偶性相同，那么 $\textit{nums}[i]$ 必须和 $\textit{nums}[i_0]$ 相等，否则必须比 $\textit{nums}[i_0]$ 多 $1$，即
//
// $$
// \textit{nums}[i] = \textit{nums}[i_0] + (i-i_0)\bmod 2
// $$
//
// 另外，对于 $[3,4,3,4,5,4,5]$ 这样的数组，第一组交替子数组为 $[3,4,3,4]$，第二组交替子数组为 $[4,5,4,5]$，这两组有一个数是**重叠**的，所以下面代码在外层循环末尾要把 $i$ 减一。
// 会一个模板是远远不够的，需要大量练习才能灵活运用。
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/longest-alternating-subarray/solutions/2615916/jiao-ni-yi-ci-xing-ba-dai-ma-xie-dui-on-r57bz/
func alternatingSubarray2(nums []int) int {
	ans := -1
	i, n := 0, len(nums)
	for i < n-1 {
		if nums[i+1]-nums[i] != 1 {
			i++ // 直接跳过
			continue
		}
		i0 := i // 记录这一组的开始位置
		// i 和 i+1 已经满足要求，从 i+2 开始判断
		for i += 2; i < n; i++ {
			if nums[i] != nums[i0]+(i-i0)%2 {
				break
			}
		}
		// 从 i0 到 i-1 是满足题目要求的（并且无法再延长的）子数组
		ans = max(ans, i-i0)
		i--
	}
	return ans
}

// 分组循环
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func alternatingSubarray3(nums []int) int {
	var n = len(nums)
	var ans = -1

	for i := 0; i < n-1; {
		if nums[i]+1 != nums[i+1] {
			i++
			continue
		}

		j := i
		for i++; i < n; i++ {
			if nums[i] != nums[j]+(i-j)%2 {
				break
			}
		}
		ans = max(ans, i-j)
		i--
	}
	return ans
}
