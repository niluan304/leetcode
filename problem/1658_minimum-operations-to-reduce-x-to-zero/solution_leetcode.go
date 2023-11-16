package main

// 方法二：直接双指针
// 如果要正向计算也是可以的，就是写起来稍微有点麻烦：首先算出最长的元素和不超过 x 的后缀，然后不断枚举前缀长度，另一个指针指向后缀最左元素，答案就是前缀+后缀长度之和的最小值。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/minimum-operations-to-reduce-x-to-zero/solutions/1/ni-xiang-si-wei-pythonjavacgo-by-endless-b4jt/
func leetcodePrefixSum(nums []int, x int) int {
	s, n := 0, len(nums)
	right := n
	for right > 0 && s+nums[right-1] <= x { // 计算最长后缀
		right--
		s += nums[right]
	}
	if right == 0 && s < x { // 全部移除也无法满足要求
		return -1
	}
	ans := n + 1
	if s == x {
		ans = n - right
	}
	for left, num := range nums {
		s += num
		for ; right < n && s > x; right++ { // 缩小后缀长度
			s -= nums[right]
		}
		if s > x { // 缩小失败，说明前缀过长
			break
		}
		if s == x {
			ans = min(ans, left+1+n-right) // 前缀+后缀长度
		}
	}
	if ans > n {
		return -1
	}
	return ans
}

func min(a, b int) int {
	if b < a {
		return b
	}
	return a
}

// 7.94 MB 的代码示例
func leetcodeMinMemory(nums []int, x int) int {

	ans := len(nums)
	sum := 0
	left := 0
	k := 0
	if len(nums) < 1 {
		return -1
	}
	if nums[0] > x && nums[len(nums)-1] > x {
		return -1
	}
	for _, v := range nums {
		sum = sum + v
	}
	if sum < x {
		return -1
	}
	if sum == x {
		return len(nums)
	}
	m := sum - x
	for right, c := range nums {
		k = k + c
		for k > m {
			k = k - nums[left]
			left++
		}
		if k == m {
			ans = min(ans, len(nums)-right+left-1)
		}

	}
	return ans
}
