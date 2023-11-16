package main

func maxRepOpt1(text string) int {
	var (
		left = 0 // 滑动窗口的左边界
		most = 0 // 寻找连续字符的最大长度, 最多含有一个不同字符
		ans  = 0 // 记录最长

		count = [26]int{} // 记录当前子串中, 各字符的重复次数
		total = [26]int{} // 记录每个字符的总频数
	)

	for _, b := range text {
		total[b-'a']++
	}

	for right, b := range text {
		x := int(b - 'a')
		count[x]++

		// 更新当前子串中, 最大重复字符次数
		most = _max(most, count[x])

		// 当前子串中, 不同字符的个数 > 1, 需要缩小窗口
		for right-left > most {
			count[text[left]-'a']--
			left++
		}

		// 最长连续字符的总频数 > 连续字符的最大长度, 需要+1
		if most == count[x] && most < total[x] {
			most++
		}

		ans = _max(ans, most)
	}

	return ans
}

func _max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
