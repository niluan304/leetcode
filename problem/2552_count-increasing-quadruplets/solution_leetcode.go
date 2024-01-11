package main

// 这里的数组v是统计以 j 为 “中间“ 的、类似132顺序的三元组的个数
//
// 如果`nums[j]<nums[i]`，j之前的类似132顺序的三元组个数为`v[j]`，总数`sum+=v[j]`
//
// 如果`nums[j]>nums[i]`，那么对于j，可以多z个132顺序的三元组，z是j之前小于`nums[i]`的个数
//
// 作者：DestinyGod
// 链接：https://leetcode.cn/problems/count-increasing-quadruplets/solutions/2080664/by-destiny-god-4qc6/
func leetcode(nums []int) (ans int64) {
	// i < j < k < l
	// nums[i] < nums[k] < nums[j] < nums[l]
	n := len(nums)
	cnt := make([]int64, n) // 132个数
	var large int64
	// 先枚举 i, 再枚举 j(j < i)
	// 先从小到大枚举 j 统计答案
	// 这里的 i 相当于题目描述的第四个下标 "l"，j 相当于题目中的第二个下标 "j"
	// 此时，如果 nums[j] < nums[i]，那么答案就应当加上以 j 为 "中间" 的 "132" 数量
	for i := 0; i < n; i++ {
		large = 0
		for j := 0; j < i; j++ {
			// 只要 nums[j] > nums[i]，那么以 j 为中间的模式数量（v[j]）
			// 应当加上 nums[j] 之前的、小于 nums[i] 的数量, 也就是 count
			if nums[i] > nums[j] {
				// i < j < k且nums[i] < nums[k] < nums[j] 的个数
				ans += cnt[j]
				// 前面小于nums[i]的个数
				large++
			} else {
				// i < j < k且nums[i] < nums[k] < nums[j] 的个数
				cnt[j] += large
			}
		}
	}
	return
}
