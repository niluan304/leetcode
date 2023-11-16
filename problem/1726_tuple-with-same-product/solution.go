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

// 哈希表
// 时间复杂度：O(n^2)
// 空间复杂度：O(n^2)
func tupleSameProduct(nums []int) int {
	var n = len(nums)
	var m = make(map[int]int, n*n/2)
	for i, a := range nums {
		for _, b := range nums[i+1:] {
			m[a*b]++
		}
	}

	var ans = 0
	for _, x := range m {
		ans += x * (x - 1) / 2 // 组合数公式: $C_{n}^{m}$
	}
	return ans * 8
}
