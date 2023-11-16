package main

func twoSum(nums []int, target int) []int {
	var idx = map[int]int{}

	for i, num := range nums {
		if j, ok := idx[target-num]; ok {
			return []int{j, i}
		}

		idx[num] = i
	}

	return nil
}
