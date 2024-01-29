package main

// 哈希表
//
// 类似 # [128. 最长连续序列](https://leetcode.cn/problems/longest-consecutive-sequence/)
//
// - 时间复杂度：$\mathcal{O}(n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func maximumLength(nums []int) int {
	var count = map[int]int{}
	for _, num := range nums {
		count[num]++
	}

	var ans = (count[1]-1)/2*2 + 1
	delete(count, 1) // 特判 1

	var visit = map[int]bool{}
	for num, _ := range count {
		if visit[num] {
			continue
		}

		for i := 0; ; i++ {
			if x := count[num]; x <= 1 {
				// if x == 0: i*2-1;
				// if x == 1: i*2+1;
				ans = max(ans, i*2+x*2-1)
				break
			}
			visit[num] = true
			num = num * num
		}
	}
	return ans
}
