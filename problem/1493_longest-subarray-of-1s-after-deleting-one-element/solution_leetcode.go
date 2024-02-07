package main

// 方法二：递推优化
// 思路
//
// 我们也可以修改递推的方式使用一次就可以得到答案。
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/longest-subarray-of-1s-after-deleting-one-element/solutions/327132/shan-diao-yi-ge-yuan-su-yi-hou-quan-wei-1-de-zui-c/
func leetcode1(nums []int) int {
	var (
		ans = 0
		p0  = 0
		p1  = 0
	)

	for _, num := range nums {
		if num == 0 {
			p1, p0 = p0, 0
		} else {
			p0++
			p1++
		}

		ans = max(ans, p1)
	}

	if ans == len(nums) {
		ans--
	}

	return ans
}

// 24 ms 的代码示例
func leetcodeMinTime(nums []int) int {
	left := 0
	ans := 0
	sum := 0
	for right := 0; right < len(nums); right++ {
		sum += nums[right]
		for left <= right && sum+1 < right-left+1 {
			sum -= nums[left]
			left++
		}
		if ans < right-left+1 {
			ans = right - left + 1
		}
	}
	return ans - 1
}
