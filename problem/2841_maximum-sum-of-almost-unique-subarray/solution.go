package main

// 滑动窗口
// 看到「长度固定的子数组」就要想到滑动窗口！
// 维护窗口内的元素出现次数 $\textit{cnt}$，以及元素和 $\textit{sum}$。
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
// - 空间复杂度：$\mathcal{O}(k)$。哈希表的大小不会超过窗口长度，即 $k$。
func maxSum(nums []int, m int, k int) int64 {
	var ans, left, n = 0, 0, len(nums)

	var sum = 0                      // 记录滑动窗口的和
	var count = make(map[int]int, n) // 使用哈希表记录滑动窗口的出现次数，
	for right, num := range nums {
		// 滑动窗口添加元素
		count[num]++
		sum += num

		if right-left+1 < k { // 窗口大小小于 k，不需要移动 left，先跳过
			continue
		}

		if len(count) >= m { // 当前的滑动窗口满足「几乎唯一」的条件，尝试更新最大和
			ans = max(ans, sum)
		}

		// 滑动窗口删除元素
		v := nums[left]
		sum -= v
		count[v]--
		if count[v] == 0 {
			delete(count, v)
		}

		left++
	}
	return int64(ans)
}
