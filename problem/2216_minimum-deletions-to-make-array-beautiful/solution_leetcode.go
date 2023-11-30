package main

// #### 方法一：一次遍历
//
// **思路与算法**
//
// 我们记满足 $i \bmod 2 = 0$ 并且 $\textit{nums}[i] = \textit{nums}[i+1]$ 的下标 $i$ 为「坏下标」。根据题目描述，我们需要删掉最少数量的元素，使得数组中不存在「坏下标」，并且数组的长度是偶数。
//
// 当数组 $\textit{nums}$ 中存在「坏下标」时，我们找出最小的那个「坏下标」记为 $i_0$，那么：
//
// - 当 $i < i_0$ 时，$i$ 一定不是「坏下标」；
// - 当 $i > i_0$ 时，$i$ 可能是「坏下标」，也可能不是「坏下标」。
//
// 为了将 $i_0$ 变得不是「坏下标」，我们必须从数组 $\textit{nums}$ 中删除元素。根据删除元素的位置 $i$ 与 $i_0$ 的关系，会有下面的 $3$ 种情况：
//
// 1. 当 $i > i_0 + 1$ 时，删除 $\textit{nums}[i]$ 会使得 $\textit{nums}[i_0]$ 和 $\textit{nums}[i_0+1]$ 均不发生任何变化，此时 $i_0$ 仍然是一个「坏下标」，因此是无意义的操作；
//
// 2. 当 $i < i_0$ 时，删除 $\textit{nums}[i]$ 会使得原先的 $\textit{nums}[i_0]$ 和 $\textit{nums}[i_0+1]$ 向左移动一个单位，变为 $\textit{nums}[i_0-1]$ 和 $\textit{nums}[i_0]$。虽然它们相等，但是 $(i_0-1) \bmod 2 \neq 0$，所以 $i_0-1$ 不是「坏下标」。但此时：
// - $(1)$ 可能会有更小的「坏下标」出现。例如数组为 $[1, 2, 1, 1]$，$i_0 = 2$ 是最小的「坏下标」，我们删除 $i = 1$ 位置的元素，得到数组 $[1, 1, 1]$，此时出现了更小的「坏下标」$i_0 = 0$；
// - $(2)$ 可能 $i_0$ 仍然是一个「坏下标」。例如数组为 $[1, 2, 3, 3, 3]$，$i_0 = 2$ 是最小的「坏下标」，我们删除 $i = 1$ 位置的元素，得到数组 $[1, 3, 3, 3]$，此时 $i_0 = 2$ 仍然是一个「坏下标」。
//
// 3. 当 $i = i_0$ 或 $i = i_0 + 1$ 时，这两种情况会得到完全相同的数组。我们不妨考虑 $i = i_0$，可以发现，上面一种情况中的 $(2)$ 仍然可能出现，但 $(1)$ 一定不会出现。
//
// 根据上面的分类讨论，我们可以得出结论：
//
// - 删除 $i > i_0 + 1$ 位置的元素是没有必要的；
// - 删除 $i < i_0$ 与 $i = i_0$ 位置的元素，它们对于所有 $i > i_0$ 的元素的影响是**完全相同**的。对于 $i \leq i_0$ 位置的元素，前者可能会有更小的「坏下标」出现，而后者一定不会。
//
// 因此，删除 $i = i_0$ 位置的元素是**严格最优**的决策。这样一来，我们的解决方法就变得非常简单：我们只需要对数组 $\textit{nums}$ 进行一次遍历，当遍历到一个「坏下标」时，我们就将答案增加 $1$（表示删除这个「坏下标」对应的元素）。需要注意的是，删除一个元素后，所有后续元素下标的奇偶性发生了变化，所以我们还需要使用一个
// 变量记录当前的下标时偶数还是奇数。
//
// 最终得到的数组的长度需要是偶数，因此，如果我们最后得到了一个奇数长度的数组，还需要额外删除一个元素。显然删除最后一个元素就可以满足要求。
//
// **代码**
// *复杂度分析**
//
// - 时间复杂度：$O(n)$，其中 $n$ 是数组 $\textit{nums}$ 的长度。
//
// - 空间复杂度：$O(1)$。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/minimum-deletions-to-make-array-beautiful/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode(nums []int) int {
	n, ans, check := len(nums), 0, true
	for i := 0; i+1 < n; i++ {
		if nums[i] == nums[i+1] && check {
			ans++
		} else {
			check = !check
		}
	}
	if (n-ans)%2 != 0 {
		ans++
	}
	return ans
}