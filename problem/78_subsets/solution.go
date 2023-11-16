package main

func subsets(nums []int) [][]int {
	var ans = [][]int{{}}
	for _, num := range nums {
		for _, v := range ans {
			item := append(append([]int{}, v...), num)
			ans = append(ans, item)

			//ans = append(ans, append(v, num))  // 直接修改v似乎会有问题
		}
	}
	return ans
}

func subsets2(nums []int) [][]int {
	var ans = [][]int{{}}
	var path []int

	var dfs func(idx int)
	dfs = func(idx int) {
		for i := idx; i < len(nums); i++ {
			path = append(path, nums[i])
			ans = append(ans, append([]int{}, path...))
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}
