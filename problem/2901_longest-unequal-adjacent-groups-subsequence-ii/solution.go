package main

func getWordsInLongestSubsequence(_ int, words []string, groups []int) []string {
	// 按长度分组，可以剪枝
	var m = make(map[int][]Word, 10) // 1 <= words[i].length <= 10
	for i, word := range words {
		m[len(word)] = append(m[len(word)], Word{
			word:  word,
			index: i,
		})
	}

	check := func(x, y Word) bool {
		if groups[x.index] == groups[y.index] {
			return false
		}
		count := 0
		for i := 0; i < len(x.word); i++ {
			if x.word[i] != y.word[i] {
				count++
			}
		}
		return count == 1
	}

	var ans []string
	for _, pairs := range m {
		item := LIS(pairs, check)
		if len(item) > len(ans) {
			ans = item
		}
	}
	return ans
}

type Word struct {
	word  string
	index int
}

// LIS 最长递增子序列
//
// dp
// 要求返回选择的物品，增加 Last 字段来记录从哪里转移过来的
//
// - 时间复杂度：$\mathcal{O}(n^2 \codt L)$。其中 $L$ 为 $\textit{words}[i]$ 的长度，至多为 $10$。
// - 空间复杂度：$\mathcal{O}(n)$。
func LIS(pairs []Word, check func(x, y Word) bool) []string {
	n := len(pairs)

	type Item struct{ Value, Last int }
	var dp = make([]Item, n+1)

	maxIdx := 0
	for i := 1; i <= n; i++ {
		dp[i] = Item{Value: 1, Last: 0}
		for j := 1; j < i; j++ {
			// 遍历 [1, n] 而不是 [0, n-1]，对 pairs 取值需要 -1
			if check(pairs[j-1], pairs[i-1]) && dp[i].Value <= dp[j].Value {
				dp[i] = Item{
					Value: dp[j].Value + 1,
					Last:  j,
				}
			}
		}

		if dp[i].Value > dp[maxIdx].Value {
			maxIdx = i
		}
	}

	var ans = make([]string, dp[maxIdx].Value)
	for i := maxIdx; i > 0; i = dp[i].Last {
		ans[dp[i].Value-1] = pairs[i-1].word
	}
	return ans
}

// LIS 最长递增子序列
//
// dp
// 要求返回选择的物品列表，增加 Last 字段来记录从哪里转移过来的
//
// - 时间复杂度：$\mathcal{O}(n^2 \codt L)$。其中 $L$ 为 $\textit{words}[i]$ 的长度，至多为 $10$。
// - 空间复杂度：$\mathcal{O}(n)$。
func getWordsInLongestSubsequence2(n int, words []string, groups []int) []string {
	type Item struct{ Value, Last int }
	var dp = make([]Item, n+1)

	check := func(i, j int) bool {
		x, y := words[i], words[j]
		// 按长度分组，可以剪枝
		if len(x) != len(y) || groups[i] == groups[j] {
			return false
		}
		count := 0
		for k := 0; k < len(x); k++ {
			if x[k] != y[k] {
				count++
			}
		}
		return count == 1
	}

	maxIdx := 0
	for i := 1; i <= n; i++ {
		dp[i] = Item{Value: 1, Last: 0}
		for j := 1; j < i; j++ {
			// 遍历 [1, n] 而不是 [0, n-1]，因此索引需要向左偏移 1 位
			if check(i-1, j-1) && dp[i].Value <= dp[j].Value {
				dp[i] = Item{
					Value: dp[j].Value + 1,
					Last:  j,
				}
			}
		}
		if dp[i].Value > dp[maxIdx].Value {
			maxIdx = i
		}
	}

	var ans = make([]string, dp[maxIdx].Value)
	for i := maxIdx; i > 0; i = dp[i].Last {
		ans[dp[i].Value-1] = words[i-1]
	}
	return ans
}
