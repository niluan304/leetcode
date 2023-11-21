package main

// 贪心
// 删除元素后，剩余元素都会往前移动，需要分类讨论索引的变动，这样的方式来计算删除数会比较繁琐。
// 美化数组是什么？
// 1. len(nums) 是偶数
// 2. 两两分组后，组内元素不能相等
// 那么反向思考：题目可以等价于 数组原长度 - 最长的美化数组长度
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func minDeletion(nums []int) int {
	var n = len(nums)
	var idx = 0
	var last = -1
	for _, num := range nums {
		if idx%2 == 1 && num == last {
			continue
		}
		idx++
		last = num
	}
	if idx%2 == 1 {
		idx--
	}
	return n - idx
}
