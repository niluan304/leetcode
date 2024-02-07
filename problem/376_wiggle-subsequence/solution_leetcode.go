package main

// #### 方法一：动态规划
//
// **思路及解法**
//
// 每当我们选择一个元素作为摆动序列的一部分时，这个元素要么是上升的，要么是下降的，这取决于前一个元素的大小。那么列出状态表达式为：
//
// 1. $\textit{up}[i]$ 表示以前 $i$ 个元素中的某一个为结尾的最长的「上升摆动序列」的长度。
//
// 2. $\textit{down}[i]$ 表示以前 $i$ 个元素中的某一个为结尾的最长的「下降摆动序列」的长度。
//
// 下面以 $\textit{up}[i]$ 为例，说明其状态转移规则：
//
// 1. 当 $\textit{nums}[i] \leq \textit{nums}[i - 1]$ 时，我们无法选出更长的「上升摆动序列」的方案。因为对于任何以 $\textit{nums}[i]$ 结尾的「上升摆动序列」，我们都可以将 $\textit{nums}[i]$ 替换为 $\textit{nums}[i - 1]$，使其成为以 $\textit{nums}[i - 1]$ 结尾的「上升摆动序列」。
//
// 2. 当 $\textit{nums}[i] > \textit{nums}[i - 1]$ 时，我们既可以从 $up[i - 1]$ 进行转移，也可以从 $\textit{down}[i - 1]$ 进行转移。下面我们证明从 $\textit{down}[i - 1]$ 转移是必然合法的，即必然存在一个 $\textit{down}[i - 1]$ 对应的最长的「下降摆动序列」的末尾元素小于 $\textit{nums}[i]$。
//
// - 我们记这个末尾元素在原序列中的下标为 $j$，假设从 $j$ 往前的第一个「谷」为 $\textit{nums}[k]$，我们总可以让 $j$ 移动到 $k$，使得这个最长的「下降摆动序列」的末尾元素为「谷」。
//
// - 然后我们可以证明在这个末尾元素为「谷」的情况下，这个末尾元素必然是从 $\textit{nums}[i]$ 往前的第一个「谷」。证明非常简单，我们使用反证法，如果这个末尾元素不是从 $\textit{nums}[i]$ 往前的第一个「谷」，那么我们总可以在末尾元素和 $\textit{nums}[i]$ 之间取得一对「峰」与「谷」，接在这个「下降摆动序列」后，使其更长。
//
// - 这样我们知道必然存在一个 $\textit{down}[i - 1]$ 对应的最长的「下降摆动序列」的末尾元素为 $\textit{nums}[i]$ 往前的第一个「谷」。这个「谷」必然小于 $\textit{nums}[i]$。证毕。
//
// 这样我们可以用同样的方法说明 $\textit{down}[i]$ 的状态转移规则，最终的状态转移方程为：
//
// $$
// \begin{aligned}
// &\textit{up}[i]=
// \begin{cases}
// \textit{up}[i - 1],& \textit{nums}[i] \leq \textit{nums}[i - 1] \\
// \max(\textit{up}[i - 1], \textit{down}[i - 1] + 1),& \textit{nums}[i] > \textit{nums}[i - 1]
// \end{cases} \\\\
// &\textit{down}[i]=
// \begin{cases}
// \textit{down}[i - 1],& \textit{nums}[i] \geq \textit{nums}[i - 1] \\
// \max(\textit{up}[i - 1] + 1, \textit{down}[i - 1]),& \textit{nums}[i] < \textit{nums}[i - 1]
// \end{cases}
// \end{aligned}
// $$
//
// 最终的答案即为 $\textit{up}[n-1]$ 和 $\textit{down}[n-1]$ 中的较大值，其中 $n$ 是序列的长度。
//
// **代码**
// *复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 是序列的长度。我们只需要遍历该序列一次。
//
// - 空间复杂度：$O(1)$。我们只需要常数空间来存放若干变量。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/wiggle-subsequence/submissions/475211844/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetocde1(nums []int) int {
	n := len(nums)
	if n < 2 {
		return n
	}
	up := make([]int, n)
	down := make([]int, n)
	up[0] = 1
	down[0] = 1
	for i := 1; i < n; i++ {
		if nums[i] > nums[i-1] {
			up[i] = max(up[i-1], down[i-1]+1)
			down[i] = down[i-1]
		} else if nums[i] < nums[i-1] {
			up[i] = up[i-1]
			down[i] = max(up[i-1]+1, down[i-1])
		} else {
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	return max(up[n-1], down[n-1])
}
