package main

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func _min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func _abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// dp 动态规划, LIS
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func wiggleMaxLength(nums []int) int {
	type Item struct {
		Positive int // nums[i] > nums[i-1]
		Negative int // nums[i] < nums[i-1]
	}

	var n = len(nums)
	var dp = make([]Item, n)
	var ans = 0

	for i, num := range nums {
		positive, negative := 0, 0
		for j := i - 1; j >= 0; j-- {
			if num > nums[j] {
				positive = _max(positive, dp[j].Negative)
			}
			if num < nums[j] {
				negative = _max(negative, dp[j].Positive)
			}
		}

		dp[i] = Item{
			Positive: positive + 1,
			Negative: negative + 1,
		}

		ans = _max(ans, _max(dp[i].Positive, dp[i].Negative))
	}
	return ans
}
