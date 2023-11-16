package main

// 2.40 MB 的代码示例
func leetcodeMinMemory(nums []int) bool {
	// 计算数组中所有元素的和sum
	sum := 0
	for _, num := range nums {
		sum += num
	}
	// 如果sum是奇数，直接返回false
	if sum%2 == 1 {
		return false
	}
	// 计算目标和target，即sum/2
	target := sum / 2
	// 创建一个长度为target+1的布尔数组dp，其中dp[i]表示是否存在一个子集，使得它的和等于i
	dp := make([]bool, target+1)
	// 初始化dp[0]为true，表示空集的和为0
	dp[0] = true
	// 遍历数组中的每个元素num
	for _, num := range nums {
		// 从后往前更新dp数组，即对于每个i从target到num之间
		for i := target; i >= num; i-- {
			// 如果dp[i-num]为true，那么就表示存在一个子集，使得它的和等于i-num，
			// 那么我们就可以把num加入到这个子集中，使得它的和等于i，因此我们把dp[i]设为true
			if dp[i-num] {
				dp[i] = true
			}
		}
	}
	// 返回dp[target]的值，如果为true，就表示存在一个子集，使得它的和等于target，那么就可以分割成两个和相等的子集；如果为false，就表示不存在这样的子集，那么就无法分割成两个和相等的子集
	return dp[target]
}
