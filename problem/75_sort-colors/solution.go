package main

func sortColors(nums []int) {
	var count = [3]int{}

	for _, num := range nums {
		count[num]++
	}

	i := 0
	for num, cnt := range count {
		for j := 0; j < cnt; j++ {
			nums[i] = num
			i++
		}
	}
}
