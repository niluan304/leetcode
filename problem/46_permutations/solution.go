package main

func permute(nums []int) [][]int {
	var (
		ans  [][]int
		path []int
	)

	var dfs func(nums []int)
	dfs = func(nums []int) {
		if len(nums) == 0 {
			ans = append(ans, append([]int{}, path...))
			return
		}

		for i, v := range nums {
			path = append(path, v)

			var list = append(append([]int{}, nums[:i]...), nums[i+1:]...)
			dfs(list)

			path = path[:len(path)-1]
		}
	}

	dfs(nums)

	return ans
}
