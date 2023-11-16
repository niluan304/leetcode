package main

import "math"

var (
	temp []int
	ans  [][]int

	n int
)

// 方法一：二进制枚举 + 哈希
// 思路与算法
// 我们可以采取最朴素的思路，即枚举出所有的子序列，然后判断当前的子序列是否是非严格递增的。
// 那么我们可以用什么办法来枚举所有的子序列呢？我们需要从原序列中选出一些数，来形成新的序列，即原序列的子序列。
// 对于原序列的每一个数来说，都有两种可能的状态，即被选中或者不被选中。
// 如果我们用 1 代表被选中，0 代表不被选中，假设一个序列长度为 3，那么选出的子序列就对应着下面的八种状态：
func leetcode1(nums []int) [][]int {
	n = len(nums)
	var ans [][]int
	set := map[int]bool{}
	for i := 0; i < 1<<n; i++ {
		findSubsequences1(i, nums)
		hashValue := getHash(263, int(1e9+7))
		if check() && !set[hashValue] {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
			set[hashValue] = true
		}
	}
	return ans
}

func findSubsequences1(mask int, nums []int) {
	temp = []int{}
	for i := 0; i < n; i++ {
		if (mask & 1) != 0 {
			temp = append(temp, nums[i])
		}
		mask >>= 1
	}
}

func getHash(base, mod int) int {
	hashValue := 0
	for _, x := range temp {
		hashValue = hashValue*base%mod + (x + 101)
		hashValue %= mod
	}
	return hashValue
}

func check() bool {
	for i := 1; i < len(temp); i++ {
		if temp[i] < temp[i-1] {
			return false
		}
	}
	return len(temp) >= 2
}

// 方法二：递归枚举 + 减枝
// 思路与算法
// 我们也可以用递归的方法实现二进制枚举，像「方法一」那样枚举出所有的子序列，
// 然后判断是否合法。直接把方法一变成递归形式，我们可以得到这样的代码：
func leetcode2(nums []int) [][]int {
	ans = [][]int{}
	dfs(0, math.MinInt32, nums)
	return ans
}

func dfs(cur, last int, nums []int) {
	if cur == len(nums) {
		if len(temp) >= 2 {
			t := make([]int, len(temp))
			copy(t, temp)
			ans = append(ans, t)
		}
		return
	}
	if nums[cur] >= last {
		temp = append(temp, nums[cur])
		dfs(cur+1, nums[cur], nums)
		temp = temp[:len(temp)-1]
	}
	if nums[cur] != last {
		dfs(cur+1, last, nums)
	}
}

func leetcode3(nums []int) [][]int {
	var (
		ans  = make([][]int, 0, len(nums))
		path = make([]int, 0, len(nums))
	)

	var dfs func(idx int)
	dfs = func(idx int) {
		if len(path) >= 2 {
			ans = append(ans, append([]int{}, path...))
		}

		used := make(map[int]bool, len(nums)) // 初始化used字典，用以对同层元素去重
		for i := idx; i < len(nums); i++ {
			if used[nums[i]] { // 去重
				continue
			}
			if len(path) == 0 || nums[i] >= path[len(path)-1] {
				path = append(path, nums[i])
				used[nums[i]] = true
				dfs(i + 1)
				path = path[:len(path)-1]
			}
		}
	}

	dfs(0)
	return ans
}

// 12ms代码示例
func leetcode4(nums []int) [][]int {
	res := make([][]int, 0)

	bT491(&res, []int{}, nums)
	return res
}

func bT491(res *[][]int, cur, nums []int) {
	if len(nums) == 0 {
		return
	}

	m := make(map[int]bool, len(nums))
	for i := 0; i < len(nums); i++ {
		if m[nums[i]] {
			continue
		}
		m[nums[i]] = true

		if len(cur) == 0 {

			cur = append(cur, nums[i])

			bT491(res, cur, nums[i+1:])
			cur = cur[:len(cur)-1]
		} else {
			if nums[i] >= cur[len(cur)-1] {
				cur = append(cur, nums[i])
				tmp := make([]int, len(cur))
				copy(tmp, cur)
				*res = append(*res, tmp)
				bT491(res, cur, nums[i+1:])
				cur = cur[:len(cur)-1]
			}
		}
	}

}
