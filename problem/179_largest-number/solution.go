package main

import (
	"sort"
	"strconv"
	"strings"
)

func largestNumber(nums []int) string {
	var list = make([]string, 0, len(nums))
	for _, num := range nums {
		list = append(list, strconv.Itoa(num))
	}

	sort.Slice(list, func(i, j int) bool {
		vi, vj := list[i], list[j]
		return vi+vj > vj+vi
	})

	if list[0] == "0" {
		return "0"
	}

	return strings.Join(list, "")
}

func leetcode(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := nums[i], nums[j]
		sx, sy := 10, 10
		for sx <= x {
			sx *= 10
		}
		for sy <= y {
			sy *= 10
		}
		return sy*x+y > sx*y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	ans := []byte{}
	for _, x := range nums {
		ans = append(ans, strconv.Itoa(x)...)
	}
	return string(ans)
}

// leetcode 2 0ms示例
func leetcode2(nums []int) string {
	arr := []string{}
	for _, v := range nums {
		arr = append(arr, strconv.Itoa(v))
	}
	sort.Slice(arr, func(i, j int) bool {
		// 这种排序方式对于输入数组 没有相同数字开头 的时候是有效的
		if arr[i][0] > arr[j][0] {
			return true
		} else {
			return arr[i]+arr[j] > arr[j]+arr[i]
		}
	})
	if arr[0] == "0" {
		return "0"
	}
	return strings.Join(arr, "")
}
