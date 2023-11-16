package main

// 差分数组
// 时间复杂度：O(N)
// 空间复杂度：O(N)
func carPooling(trips [][]int, capacity int) bool {
	const N = 1000 + 2        // 差分数组/原数组的长度, 题目说明 to_i <= 1000, 也可以遍历找到 max(to_i), 数组长度 +2 是为避免数组索引越界
	var diff = make([]int, N) // 原数组为 [0]*N, 差分数组也为 [0]*N
	for _, trip := range trips {
		x, left, right := trip[0], trip[1], trip[2]
		diff[left] += x
		//if right < N {
		// diff[right] -= x
		//}
		diff[right] -= x //  right 处下车, 因此是 [left, right-1] 区间 + x
	}

	// 从差分数组还原为原数组
	for i := 0; i < N; i++ {
		if i > 0 {
			diff[i] += diff[i-1]
		}

		if diff[i] > capacity {
			return false
		}
	}
	return true
}
