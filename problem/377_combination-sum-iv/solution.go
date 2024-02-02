package main

import (
	"slices"
	"sort"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

func combinationSum4(nums []int, target int) int {
	return waysToTarget(target, nums)
}

// 爬楼梯
// 需要 target 阶你才能到达楼顶。你有 ways[:] 方式向上爬
// 你从 0 开始出发，有多少种方式爬到楼顶？
//
// - 时间复杂度：$\mathcal{O}(n \codt m)$。
// - 空间复杂度：$\mathcal{O}(n \codt m)$。
//
// 题目保证 ways[i] >= 1
// 但如果 ways[i] 可以为负数，该怎么处理？
func waysToTarget(target int, ways []int) int {
	slices.Sort(ways[:])

	var dfs func(n int) int
	dfs = func(n int) int {
		if n == 0 {
			return 1
		}

		res := 0
		for _, way := range ways {
			if way > n {
				break
			}
			res += dfs(n - way)
		}
		return res
	}

	MemorySearch(&dfs)
	ans := dfs(target)
	return ans
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func combinationSum4II(nums []int, target int) int {
	// 升序排序, 方便剪枝
	sort.Ints(nums)

	var cache = make(map[int]int, target+1)
	cache[0] = 1
	var dfs func(target int) int
	dfs = func(target int) int {
		v, ok := cache[target]
		if ok {
			return v
		}

		for _, num := range nums {
			if target < num {
				break
			}
			v += dfs(target - num)
		}
		cache[target] = v
		return v
	}

	return dfs(target)
}

// dfs + 记忆化搜索
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func combinationSum4III(nums []int, target int) int {
	// 升序排序, 方便剪枝
	sort.Ints(nums)

	var cache = make([]int, target+1)
	const Empty = -1
	for i := 1; i <= target; i++ {
		cache[i] = Empty
	}
	cache[0] = 1

	var dfs func(target int) int
	dfs = func(target int) int {
		if cache[target] != Empty {
			return cache[target]
		}

		v := 0
		for _, num := range nums {
			if target < num {
				break
			}
			v += dfs(target - num)
		}
		cache[target] = v
		return v
	}

	//// debug: 打印 dp数组
	//defer fmt.Println(cache)

	return dfs(target)
}

// dp 动态规划
// 时间复杂度：O(mn)
// 空间复杂度：O(n)
func combinationSum4IV(nums []int, target int) int {
	// 升序排序, 方便剪枝
	sort.Ints(nums)

	var dp = make([]int, target+1)
	// 初始化 dp数组
	// dp[0] 可以理解成target=0有几种排列，显然只有一种，那就是一个数都不放
	dp[0] = 1
	for i := nums[0]; i <= target; i++ {
		for _, num := range nums {
			if i < num {
				break
			}
			// why?
			// 个人觉得这道题不是背包问题，也没必要套用背包问题的代码框架，避免纠结于物品遍历内外层循环的问题。
			// 其实这道题换个问法就会觉得很简单： 楼梯的长度为target,每次爬楼梯可选的层数从nums数组中挑选，问有几种爬法？ 这不是入门dp问题么？
			//
			//作者：有个人
			//链接：https://leetcode.cn/problems/combination-sum-iv/solutions/2268310/zhe-dao-ti-bu-jiu-shi-pa-lou-ti-yao-by-j-h5nn/
			//来源：力扣（LeetCode）
			//著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
			dp[i] += dp[i-num]
		}
	}

	//// debug: 打印 dp数组
	//fmt.Println(dp)

	return dp[target]
}
