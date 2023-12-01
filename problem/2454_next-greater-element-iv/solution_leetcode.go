package main

// **前置知识**：[496. 下一个更大元素 I](https://leetcode.cn/problems/next-greater-element-i/)。
//
// 从左往右遍历 $\textit{nums}$，用（递减）单调栈 $s$ 记录元素，如果 $x=\textit{nums}[i]$ 比 $s$ 的栈顶大，则 $x$ 是栈顶的**下个**更大元素，弹出栈顶。最后把 $x$ 入栈（实际入栈的是下标 $i$）。
//
// 把弹出的元素加到另一个栈 $t$ 中（注意保持原始顺序），后续循环时，如果 $y=\textit{nums}[j]$ 比 $t$ 的栈顶大，则 $y$ 是栈顶的**下下个**更大元素，记录答案，弹出栈顶。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/next-greater-element-iv/description/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcodeEndless(nums []int) []int {
	ans := make([]int, len(nums))

	var t, s []int
	for i, x := range nums {
		ans[i] = -1
		for len(t) > 0 && nums[t[len(t)-1]] < x {
			ans[t[len(t)-1]] = x
			t = t[:len(t)-1]
		}
		j := len(s) - 1
		for j >= 0 && nums[s[j]] < x {
			j--
		}
		t = append(t, s[j+1:]...)
		s = append(s[:j+1], i)
	}
	return ans
}
