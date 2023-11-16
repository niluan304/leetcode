package main

// 1.73 MB 的代码示例
// 使用数组索引替代排序
func leetcodeMinMemory(num int) int {
	arr := make([]int, 10)
	for num > 0 {
		arr[num%10]++
		num /= 10
	}

	a, b := 0, 0
	for i := 0; i <= 9; i++ {
		for arr[i] > 0 {
			if a <= b {
				a = a*10 + i
			} else {
				b = b*10 + i
			}
			arr[i]--
		}
	}

	return a + b
}
