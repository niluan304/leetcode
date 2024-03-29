package main

import "sort"

func subsetsWithDup(nums []int) [][]int {
	sort.Ints(nums)

	type Value struct {
		Value int
		Count int
	}

	var values []Value
	for _, num := range nums {
		n := len(values)

		if n == 0 || values[n-1].Value != num {
			values = append(values, Value{
				Value: num,
				Count: 1,
			})
		} else {
			values[n-1].Count++
		}
	}

	var ans [][]int
	var path []int

	var dfs func(idx int)
	dfs = func(idx int) {
		ans = append(ans, append([]int{}, path...))

		for i := idx; i < len(values); i++ {
			v := &values[i]
			if v.Count == 0 {
				continue
			}

			path = append(path, v.Value)
			v.Count--

			dfs(i)

			path = path[:len(path)-1]
			v.Count++
		}
	}

	dfs(0)
	return ans
}

func subsetsWithDup2(nums []int) [][]int {
	sort.Ints(nums)

	var ans [][]int
	var path []int

	var dfs func(idx int)
	dfs = func(idx int) {
		ans = append(ans, append([]int{}, path...))

		for i := idx; i < len(nums); i++ {
			v := nums[i]

			if i != idx && v == nums[i-1] {
				continue
			}

			path = append(path, v)
			dfs(i + 1)
			path = path[:len(path)-1]
		}
	}

	dfs(0)
	return ans
}
