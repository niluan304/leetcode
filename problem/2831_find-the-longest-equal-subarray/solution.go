package main

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// Deprecated: 超时
func bruteForce(nums []int, k int) int {
	var ans = 0 // math.MaxInt32 // math.MinInt32

	for i, num := range nums {
		cnt := 0
		j := i
		for ; j < len(nums); j++ {
			if nums[j] != num {
				cnt++
			}
			if cnt > k {
				break
			}
		}

		ans = max(ans, j-i-cnt)
	}
	return ans
}

// 哈希表 + 滑动窗口
// 哈希表 map[int][]int 预处理，将相同 num 的索引保存到同一个数组中
// 然后使用滑动窗口对索引数组 "提纯"，要求杂质数量 <= k
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func longestEqualSubarray(nums []int, k int) int {
	var ans = 0
	var pos = map[int][]int{}
	for i, num := range nums {
		pos[num] = append(pos[num], i)
	}

	for _, p := range pos {
		if len(p) <= ans {
			continue
		}

		ans = max(ans, maxLength2(p, k))
	}
	return ans
}

// 滑动窗口
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func maxLength(pos []int, k int) (ans int) {
	var left = 0
	for right := 0; right < len(pos); right++ {
		for (pos[right]-pos[left])-(right-left) > k {
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}

// 滑动窗口
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(1)$。
// find a tip: let say we have a fixed sliding window size key, what's the longest sequence?
func maxLength2(pos []int, k int) int {
	var left = 0
	var ans = 1 // right = 0 的初始值
	var cnt = 0
	for right := 1; right < len(pos); right++ {
		cnt += pos[right] - pos[right-1] - 1 // "杂质" 变多
		for cnt > k {                        // "杂质" 数量过多
			cnt -= pos[left+1] - pos[left] - 1 // 筛掉 "杂质"
			left++
		}
		ans = max(ans, right-left+1)
	}
	return ans
}
