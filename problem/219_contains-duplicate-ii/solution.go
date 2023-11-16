package main

func containsNearbyDuplicate(nums []int, k int) bool {
	var (
		set = map[int]int{}
	)

	for i, num := range nums {
		j, ok := set[num]
		if ok && i-j <= k {
			return true
		}

		set[num] = i
	}

	return false
}
