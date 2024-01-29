package main

// 滑动窗口
// - 时间复杂度：$\mathcal{O}(n \cdot 26^2)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func countCompleteSubstrings(word string, k int) int {
	var ans = 0
	var n = len(word)
	for i := 0; i < n; {
		j := i
		for i++; i < n; i++ { // 可以用分组循环切割字符串
			if Abs(int(word[i])-int(word[i-1])) > 2 {
				break
			}
		}
		ans += countSubstr(word[j:i], k)
	}
	return ans
}

// s 中每个字符 恰好 出现 k 次。
// 滑动窗口
// 每个窗口大小是固定的：一定是 k 的倍数
func countSubstr(word string, k int) (ans int) {
	var count [26]int
	check := func() bool {
		for _, cnt := range count {
			if cnt > 0 && cnt != k {
				return false
			}
		}
		return true
	}

	for m := 1; m <= 26; m++ {
		size := m * k
		if size > len(word) {
			break
		}
		clear(count[:])

		for right, _ := range word {
			count[word[right]-'a']++
			left := right + 1 - size
			if left < 0 {
				continue
			}

			if check() {
				ans++
			}
			count[word[left]-'a']--
		}
	}

	return ans
}

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
