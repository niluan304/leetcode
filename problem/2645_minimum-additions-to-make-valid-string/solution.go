package main

func addMinimum(word string) int {
	word = "abc" + word + "abc" // 可以 头部只补 'c'，尾部只补 'a'。
	var ans int
	for i := 3; i < len(word); i++ {
		switch word[i-1 : i+1] {
		case "ab", "bc", "ca":
			ans += 0 // 正常的顺序，不需要修改字符，可注释
		case "ac", "ba", "cb":
			ans += 1 // 差一个字符
		case "aa", "bb", "cc":
			ans += 2 // 差两个字符
		}
	}
	return ans
}

func addMinimum2(word string) int {
	word = "abc" + word + "abc"
	var ans = 0
	for i := 3; i < len(word); i++ {
		ans += int(word[i]-word[i-1]+2) % 3
	}
	return ans
}
