package main

// 双重单调栈
// 时间复杂度: O(n)
// 空间复杂度: O(n)
func secondGreaterElement(nums []int) []int {
	var s, t []int // s 第一重栈，t 第二重栈 表示从 s 出栈过的元素
	var ans = make([]int, len(nums))
	for i, num := range nums {
		ans[i] = -1

		for m := len(t); m > 0; m-- {
			j := t[m-1]
			if num <= nums[j] {
				break
			}
			t = t[:m-1]
			ans[j] = num
		}

		m := len(s)
		for m > 0 && num > nums[s[m-1]] {
			m--
		}
		t = append(t, s[m:]...) // 第二重栈按顺序入栈 第一重栈 出栈的元素
		s = append(s[:m], i)    // 第一重栈入栈新元素
	}
	return ans
}

// 暴力穷举
// 时间复杂度: O(n^2)
// 空间复杂度: O(1)
func secondGreaterElement2(nums []int) []int {
	var n = len(nums)
	var ans = make([]int, n)
	for i, num := range nums {
		ans[i] = -1
		var cnt = 0
		for j := i + 1; j < n; j++ {
			v := nums[j]
			if v > num {
				cnt++
			}
			if cnt == 2 {
				ans[i] = v
				break
			}
		}
	}

	return ans
}
