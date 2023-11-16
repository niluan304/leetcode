package main

import "sort"

// 方法一：迭代法实现子集枚举
// 考虑数组 [1,2,2]，选择前两个数，或者第一、三个数，都会得到相同的子集。
// 也就是说，对于当前选择的数 x，若前面有与其相同的数 y，且没有选择 y，
// 此时包含 x 的子集，必然会出现在包含 y 的所有子集中。
//
// 我们可以通过判断这种情况，来避免生成重复的子集。代码实现时，可以先将数组排序；
// 迭代时，若发现没有选择上一个数，且当前数字与上一个数相同，则可以跳过当前生成的子集。
func leetcode1(nums []int) (ans [][]int) {
	sort.Ints(nums)
	n := len(nums)
outer:
	for mask := 0; mask < 1<<n; mask++ {
		t := []int{}
		for i, v := range nums {
			if mask>>i&1 > 0 {
				if i > 0 && mask>>(i-1)&1 == 0 && v == nums[i-1] {
					continue outer
				}
				t = append(t, v)
			}
		}
		ans = append(ans, append([]int(nil), t...))
	}
	return
}

// 方法二：递归法实现子集枚举
// 与方法一类似，在递归时，若发现没有选择上一个数，
// 且当前数字与上一个数相同，则可以跳过当前生成的子集。
func leetcode2(nums []int) (ans [][]int) {
	sort.Ints(nums)
	t := []int{}
	var dfs func(bool, int)
	dfs = func(choosePre bool, cur int) {
		if cur == len(nums) {
			ans = append(ans, append([]int(nil), t...))
			return
		}
		dfs(false, cur+1)
		if !choosePre && cur > 0 && nums[cur-1] == nums[cur] {
			return
		}
		t = append(t, nums[cur])
		dfs(true, cur+1)
		t = t[:len(t)-1]
	}
	dfs(false, 0)
	return
}
