package main

// 方法一：二进制（子集）枚举
// 「组合中只允许含有 1−9 的正整数，并且每种组合中不存在重复的数字」意味着这个组合中最多包含 9 个数字。我们可以把原问题转化成集合 S={1,2,3,4,5,6,7,8,9} ，我们要找出 S 的当中满足如下条件的子集：
// 大小为 k
// 集合中元素的和为 n
func leetcode1(k int, n int) (ans [][]int) {
	var temp []int
	check := func(mask int) bool {
		temp = nil
		sum := 0
		for i := 0; i < 9; i++ {
			if 1<<i&mask > 0 {
				temp = append(temp, i+1)
				sum += i + 1
			}
		}
		return len(temp) == k && sum == n
	}

	for mask := 0; mask < 1<<9; mask++ {
		if check(mask) {
			ans = append(ans, append([]int(nil), temp...))
		}
	}
	return
}

// 方法二：组合枚举
// 我们可以换一个思路：我们需要在 9 个数中选择 k 个数，让它们的和为 n。
// 这样问题就变成了一个组合枚举问题。组合枚举有两种处理方法——递归法和字典序法，在「77. 组合的官方题解」中有详细的说明。
// 这里我们要做的是做一个「在 9 个数中选择 k 个数」的组合枚举，对于枚举到的所有组合，判断这个组合内元素之和是否为 n。
func leetcode2(k int, n int) (ans [][]int) {
	var temp []int
	var dfs func(cur, rest int)
	dfs = func(cur, rest int) {
		// 找到一个答案
		if len(temp) == k && rest == 0 {
			ans = append(ans, append([]int(nil), temp...))
			return
		}
		// 剪枝：跳过的数字过多，后面已经无法选到 k 个数字
		if len(temp)+10-cur < k || rest < 0 {
			return
		}
		// 跳过当前数字
		dfs(cur+1, rest)
		// 选当前数字
		temp = append(temp, cur)
		dfs(cur+1, rest-cur)
		temp = temp[:len(temp)-1]
	}
	dfs(1, n)
	return
}
