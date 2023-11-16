package main

import "sort"

// 方法二：二分+贪心
// 也可以用贪心做。
//
// 考虑到只需要计算个数，在从左到右遍历的情况下只要当前房子可以偷，就立刻偷。
//
// 例如 nums=[1,2,3,4], mx=3，如果不偷 nums[0]=1 去偷 nums[1]=2，那么只能偷一间房子。
// 而如果偷 nums[0]=1 和 nums[2]=3，就可以偷两间房子。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/house-robber-iv/submissions/467618405/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcodeEndlesscheng(nums []int, k int) int {
	return sort.Search(1e9, func(mx int) bool {
		cnt, n := 0, len(nums)
		for i := 0; i < n; i++ {
			if nums[i] <= mx {
				cnt++ // 偷 nums[i]
				i++   // 跳过下一间房子
			}
		}
		return cnt >= k
	})
}
