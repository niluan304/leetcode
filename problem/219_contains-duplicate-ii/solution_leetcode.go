package main

// 方法二：滑动窗口
// 考虑数组 nums 中的每个长度不超过 k+1 的滑动窗口，同一个滑动窗口中的任意两个下标差的绝对值不超过 k。
// 如果存在一个滑动窗口，其中有重复元素，则存在两个不同的下标 i 和 j 满足 nums[i]=nums[j] 且 ∣i−j∣≤k。
// 如果所有滑动窗口中都没有重复元素，则不存在符合要求的下标。
// 因此，只要遍历每个滑动窗口，判断滑动窗口中是否有重复元素即可。
//
// 如果一个滑动窗口的结束下标是 i，则该滑动窗口的开始下标是 max(0,i−k)。可以使用哈希集合存储滑动窗口中的元素。
// 从左到右遍历数组 nums，当遍历到下标 iii 时，具体操作如下：
//
// 如果 i>k，则下标 i−k−1 处的元素被移出滑动窗口，因此将 nums[i−k−1] 从哈希集合中删除；
//
// 判断 nums[i] 是否在哈希集合中，如果在哈希集合中则在同一个滑动窗口中有重复元素，返回 true，
// 如果不在哈希集合中则将其加入哈希集合。
//
// 当遍历结束时，如果所有滑动窗口中都没有重复元素，返回 false。
func leetcode2(nums []int, k int) bool {
	set := map[int]struct{}{}
	for i, num := range nums {
		if i > k {
			delete(set, nums[i-k-1])
		}
		if _, ok := set[num]; ok {
			return true
		}
		set[num] = struct{}{}
	}
	return false
}

// 76 ms 的代码示例
func leetcodeMinTime(nums []int, k int) bool {
	n := len(nums)
	m := make(map[int]bool, n)
	for i, num := range nums {
		if _, ok := m[num]; ok {
			return true
		}
		m[num] = true
		if len(m) == k+1 {
			delete(m, nums[i-k])
		}
	}
	return false
}

// 7.28 MB 的代码示例
func leetcodeMinMemory(nums []int, k int) bool {
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= i+k && j < len(nums); j++ {
			if nums[i] == nums[j] {
				return true
			}
		}
	}

	return false
}
