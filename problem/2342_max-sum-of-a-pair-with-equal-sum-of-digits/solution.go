package main

func maximumSum(nums []int) int {
	var cache [100]int
	var ans = -1

	for _, num := range nums {
		x := 0
		for y := num; y != 0; y /= 10 {
			x += y % 10
		}

		v := cache[x]
		if v > 0 {
			ans = max(ans, v+num)
		}
		cache[x] = max(v, num)
	}
	return ans
}
