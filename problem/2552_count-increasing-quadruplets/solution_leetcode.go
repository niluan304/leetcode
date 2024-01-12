package main

// 这里的数组 count[:] 是统计以 j 为 "中间" 的、类似132顺序的三元组的个数
//
// 如果`nums[j]<nums[i]`，j之前的类似132顺序的三元组个数为 `count[j]`，总数 `ans += count[j]`
//
// 如果`nums[j]>nums[i]`，那么对于j，可以多z个132顺序的三元组，x 是 j 之前小于`nums[i]`的个数
//
// 作者：DestinyGod
// 链接：https://leetcode.cn/problems/count-increasing-quadruplets/solutions/2080664/by-destiny-god-4qc6/
func leetcode(nums []int) int64 {
	// 1324模式
	// i < j < k < k
	// nums[i] < nums[k] < nums[j] < nums[k]

	var n, ans = len(nums), 0

	// less[k][j] 表示在 [0, j] 这段区间，小于 nums[k] 的数量，可以计算出 `j为3, k为2` 的 `132模式` 中，有多少个 1。
	var less = make([][]int, n)
	for k := 0; k < n; k++ {
		less[k] = make([]int, k+1)
		x := 0 // x 记录 [0, j] 这段区间，有多少个数小于 nums[k]
		for j := 0; j < k; j++ {
			if nums[j] < nums[k] {
				x++
			}
			less[k][j] = x
		}
	}

	// count[k][j] 表示 在 [0, k] 这段区间内，以 `j为3` 的 `132模式` 数量
	var count = make([][]int, n)
	for k := 0; k < n; k++ {
		count[k] = make([]int, k+1)

		// 动态规划的思想
		for j := 0; j < k; j++ {
			// 在 [0, k] 区间，以 `j为3` 的 `132模式` 数量等于在 [0, k-1] 区间内的数量 + 以 `j为3, k为2` 的 `132模式` 的数量
			count[k][j] = count[k-1][j]

			// 如果 nums[j] > nums[k]，说明 nums[k] 可以作为新的 2，新增的数量：在 [0, j] 区间中小于 nums[k] 的个数，
			// 也就是以 `j为3, k为2` 的 `132模式` 中，1 的数量。
			if nums[j] > nums[k] {
				count[k][j] += less[k][j]
			}
			// 如果 nums[j] < nums[k]，说明不满足 132模式
		}
	}

	// 从小到大枚举 l 统计答案
	// 如果 nums[j] < nums[l]，那么有 `j为3`， l 为 4 的 1324模式
	// 答案应当加上以 `j为3` 的 `132模式` 的数量
	for l := 0; l < n; l++ {
		for j := 0; j < l; j++ {
			if nums[j] < nums[l] { // 符合 1324 模式
				ans += count[l][j]
			}
		}
	}

	return int64(ans)
}

func leetcode2(nums []int) int64 {
	// 1324 模式
	// i < j < k < l
	// nums[i] < nums[k] < nums[j] < nums[l]

	var n, ans = len(nums), 0

	// count[k][j] 表示 在 [0, k] 这段区间内，以 `j为3` 的 `132模式` 数量
	var count = make([][]int, n)
	for k := 0; k < n; k++ {
		count[k] = make([]int, k+1)

		x := 0 // x 记录 [0, j] 这段区间，有多少个数小于 nums[k]
		// 动态规划的思想
		for j := 0; j < k; j++ {
			// 在 [0, k] 区间，以 `j为3` 的 `132模式` 数量等于在 [0, k-1] 区间内的数量 + 以 `j为3`，以 `k为2` 的 `132模式` 的数量
			count[k][j] = count[k-1][j]

			// 如果 nums[j] > nums[k]，说明 nums[k] 可以作为新的 2，新增的数量：在 [0, j] 区间中小于 nums[k] 的个数，
			// 也就是以 `j为3, k为2` 的 `132模式` 中，1 的数量。
			if nums[j] > nums[k] {
				count[k][j] += x
			} else
			// 如果 nums[j] < nums[k]，说明不满足 `132模式`。但是小于 nums[k] 的数量 +1
			if nums[j] < nums[k] {
				x++
			}
		}
	}

	// 从小到大枚举 l 统计答案
	// 如果 nums[j] < nums[l]，那么存在 `j为3, l为4` 的 `1324模式`，答案应当加上以 `j为3` 的 `132模式` 的数量
	for l := 0; l < n; l++ {
		for j := 0; j < l; j++ {
			if nums[j] < nums[l] { // 符合 1324 模式
				ans += count[l][j]
			}
		}
	}

	return int64(ans)
}

// 空间压缩
func leetcode3(nums []int) int64 {
	// 1324 模式
	// i < j < l < l
	// nums[i] < nums[l] < nums[j] < nums[l]

	var n, ans = len(nums), 0

	// count[l][j] 表示 在 [0, l] 这段区间内，以 `j为3` 的 `132模式` 数量
	// 一边统计答案，一边预处理 count[l][j]，可以将 count[l][j] 为 一维数组
	var count = make([]int, n)
	for l := 0; l < n; l++ {
		x := 0 // x 记录 [0, j] 这段区间，有多少个数小于 nums[l]

		for j := 0; j < l; j++ {
			// 动态规划的思想
			// 在 [0, l] 区间，以 `j为3` 的 `132模式` 数量等于在 [0, l-1] 区间内的数量 + 以 `j为3`，以 `k为2` 的 `132模式` 的数量
			// 如果 nums[j] > nums[l]，说明 nums[l] 可以作为新的 2，而新增的数量为：在 [0, j] 区间中小于 nums[l] 的个数，
			// 也就是以 `j为3, k为2` 的 `132模式` 中，1 的数量。
			if nums[j] > nums[l] {
				count[j] += x
			} else
			// 如果 nums[j] < nums[l]，说明有 `j为3, l为4` 的 `1324` 模式
			// 并且 在 [0,j] 区间内 小于 nums[l] 的数量 +1
			if nums[j] < nums[l] {
				ans += count[j]
				x++
			}
		}
	}

	return int64(ans)
}
