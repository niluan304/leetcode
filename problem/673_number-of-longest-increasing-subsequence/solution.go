package main

// dp 动态规划
// 时间复杂度：O(n^2)
// 空间复杂度：O(n)
func findNumberOfLIS(nums []int) int {
	type Pair struct {
		Length int
		Kind   int
	}
	var n = len(nums)
	var dp = make([]Pair, n) // 以 nums[i] 为末尾的最长递增子序列长度
	var ans, mm = 0, 0
	for i := 0; i < n; i++ {
		length := 0
		kind := 1
		for j := 0; j < i; j++ {
			if nums[i] <= nums[j] {
				continue
			}

			vj := dp[j]
			// 寻找最长递增子序列的
			if length == vj.Length {
				kind += vj.Kind
			} else if length < vj.Length {
				length = vj.Length
				kind = vj.Kind
			}
		}

		length++
		dp[i] = Pair{
			Length: length,
			Kind:   kind,
		}

		if mm < length {
			mm = length
			ans = kind
		} else if mm == length {
			ans += kind
		}
	}

	return ans
}
