package main

func isBalance(x int) bool {
	count := make([]int, 10)
	for x > 0 {
		count[x%10]++
		x /= 10
	}
	for i := 0; i < 10; i++ {
		if count[i] > 0 && count[i] != i {
			return false
		}
	}
	return true
}

// 方法一：枚举
// 思路与算法
//
// 题目给一个整数 n ，要求返回严格大于 n 的最小数值平衡数，我们直接按照题目的要求进行模拟即可。
//
// 观察到 0 <= n <= 10^60<=n<=10, 我们可能返回的数值平衡数最大是 1224444，这个范围可以在时间要求内找到答案。
//
// 我们依次枚举大于 nnn 的整数，统计所有数字的出现频数，判断是否是数值平衡数即可。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/next-greater-numerically-balanced-number/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode1(n int) int {
	for i := n + 1; i <= 1224444; i++ {
		if isBalance(i) {
			return i
		}
	}
	return -1
}
