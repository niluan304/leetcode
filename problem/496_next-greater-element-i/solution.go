package main

// 单调栈
// 时间复杂度: O(n+m)
// 空间复杂度: O(n)
func nextGreaterElement(nums1 []int, nums2 []int) []int {
	var val = map[int]int{}
	var stack []int
	for i, num := range nums2 {
		val[num] = -1
		for m := len(stack); m > 0; m-- {
			v := nums2[stack[m-1]]
			if num < v {
				break
			}
			stack = stack[:m-1]
			val[v] = num
		}
		stack = append(stack, i)
	}

	var ans = make([]int, len(nums1))
	for i, num := range nums1 {
		ans[i] = val[num]
	}
	return ans
}
