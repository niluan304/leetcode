package main

// 单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func nextGreaterElements(nums []int) []int {
	var n = len(nums)
	var ans = make([]int, n)

	var stack []int
	for i := 0; i < n*2; i++ {
		if i < n {
			ans[i] = -1
		}

		j := i % n
		num := nums[j]
		for m := len(stack); m > 0; m-- {
			k := stack[m-1]
			if num <= nums[k] {
				break
			}
			stack = stack[:m-1]
			if ans[k] == -1 {
				ans[k] = num
			}
		}
		stack = append(stack, j)
	}
	return ans
}
