package main

func getLastMoment(n int, left []int, right []int) int {
	var ans = 0

	for _, v := range left {
		ans = max(ans, v)
	}

	for _, v := range right {
		ans = max(ans, n-v)
	}

	return ans
}
