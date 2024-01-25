package main

// 暴力穷举
// - 时间复杂度：$\mathcal{O}(n^2)$。
// - 空间复杂度：$\mathcal{O}(1)$。
func bruteForce(nums []int, indexDifference int, valueDifference int) []int {
	var n = len(nums)
	for j := indexDifference; j < n; j++ {
		for i := 0; i <= j-indexDifference; i++ { // 可以看出预处理 前缀 最小值/最大值 的影子。
			if Abs(nums[i]-nums[j]) >= valueDifference {
				return []int{i, j}
			}
		}
	}

	return []int{-1, -1}
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// 前缀和的技巧
//
// 令 i <= j，遍历 j，那么 i ∈ [0, j-indexDiff]
// 如果 nums[j] 是 (i, j) 的更大者，区间 [0, j-indexDiff] 中最小值，是最可能满足题目要求的
// 如果 nums[j] 是 (i, j) 的更小者，区间 [0, j-indexDiff] 中最大值，是最可能满足题目要求的
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func findIndices(nums []int, indexDifference int, valueDifference int) []int {
	var n = len(nums)

	type Index struct {
		MinIndex int // 记录前缀中最小值的索引
		MaxIndex int // 记录前缀中最大值的索引
	}
	var prefix = make([]Index, n)
	for i := 1; i < n; i++ {
		prefix[i] = prefix[i-1]
		if nums[i] < nums[prefix[i].MinIndex] {
			prefix[i].MinIndex = i // nums[i] 更小，更新最小值索引
		}
		if nums[i] > nums[prefix[i].MaxIndex] {
			prefix[i].MaxIndex = i // nums[i] 更大，更新最大值索引
		}
	}

	for j := indexDifference; j < n; j++ {
		index := prefix[j-indexDifference]

		// 假设 nums[j] 是满足条件的更大值，寻找 [0, j-indexDiff] 的最小值
		minIndex := index.MinIndex
		if (nums[j] - nums[minIndex]) >= valueDifference {
			return []int{minIndex, j}
		}

		// 假设 nums[j] 是满足条件的更小值，寻找 [0, j-indexDiff] 的最大值
		maxIndex := index.MaxIndex
		if (nums[maxIndex] - nums[j]) >= valueDifference {
			return []int{maxIndex, j}
		}
	}
	return []int{-1, -1}
}
