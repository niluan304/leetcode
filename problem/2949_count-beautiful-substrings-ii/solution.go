package main

// 前后缀分解 + 哈希表
//
// 1. 前缀和
// 将元音字母视为 +1，辅音字母视为 -1，那么定义前缀和 preSum：
// preSum[0] = 0
// preSum[1] = s[0]
// preSum[i] = s[0] + ... + s[i-1]
// 区间 [j, i-1] 的和为：preSum[i] - preSum[j]
//
// 假设有子字符串 s[j:i-1] (不包含 s[i] 这个字符) 满足题意
// 要求元音字母和辅音字母的数量相等
// 那么有 preSum[i] - preSum[j] == 0
// 得到约束条件一：preSum[i] == preSum[j]
//
// 2. 找到真正的 k
// 条件二：元音字母和辅音字母的数量的乘积能被 k 整除
//
// 假设有子字符串 s[j:i-1] (不包含 s[i] 这个字符) 满足题意，元音字母的数量为 vowels，
// 由于条件 1：元音字母和辅音字母的数量相等，那么元音字母和辅音字母的数量都为 vowels，
// 因此有 (vowels^2) % k == 0，类似求解 vowels % sqrt(k) == 0，
// 由于 k <= 1000，可以使用穷举来求得 sqrt(k)，记为 k2。
// 问题就转换为 vowels 得是 k2 的倍数
// 为了求出区间 [j:i-1] 的元音字母数量，
// 同样可以使用前缀和求得，记为 (vowels[i]-vowels[j])% k2 == 0，
// 即 vowels[i]%k2 -vowels[j]%k2 == 0
// 得到约束条件二：vowels[i]%k2 == vowels[j]%k2
//
// 同类型题目：[2845. 统计趣味子数组的数目](https://leetcode.cn/problems/count-of-interesting-subarrays/description/)
//
// - 时间复杂度：$\mathcal{O}(n + k)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func beautifulSubstrings(s string, k int) int64 {
	for i := 2; i < k; i++ {
		if (i*i)%k == 0 {
			k = i
			break
		}
	}

	type Item struct {
		vowels int //
		sum    int // 元音字母视为 +1，辅音字母视为 -1，元音字母和辅音字母的数量相等。
	}
	var preSum = make([]Item, len(s)+1)
	for i, b := range s {
		preSum[i+1] = preSum[i]
		switch b {
		case 'a', 'e', 'i', 'o', 'u':
			preSum[i+1].vowels++
			preSum[i+1].sum++ // 元音字母转换为：+1
		default:
			preSum[i+1].sum-- // 辅音字母转换为：-1
		}
	}

	var count = make(map[Item]int, len(s))
	var ans = 0

	for _, item := range preSum {
		item.vowels = item.vowels % k
		ans += count[item] //非空美丽子字符串，代表不能与自身匹配，因此得先累加答案，再更新 count
		count[item]++
	}

	return int64(ans)
}

// beautifulSubstrings 的空间优化版本 + 一次遍历
//
// 前后缀分解 + 哈希表
//
// 同类型题目：[2845. 统计趣味子数组的数目](https://leetcode.cn/problems/count-of-interesting-subarrays/description/)
//
// - 时间复杂度：$\mathcal{O}(n + k)$。
// - 空间复杂度：$\mathcal{O}(n)$。
func beautifulSubstrings2(s string, k int) int64 {
	for i := 2; i < k; i++ {
		if i*i%k == 0 {
			k = i
			break
		}
	}

	type Item struct {
		vowels int //
		sum    int // 元音字母视为 +1，辅音字母视为 -1，元音字母和辅音字母的数量相等。
	}

	pre, ans := Item{}, 0
	var count = make(map[Item]int, len(s))
	count[pre] = 1 // 初始化 count[preSum[0]] = 1

	for _, b := range s {
		switch b {
		case 'a', 'e', 'i', 'o', 'u':
			pre.vowels = (pre.vowels + 1) % k
			pre.sum++ // 元音字母转换为：+1
		default:
			pre.sum-- // 辅音字母转换为：-1
		}
		ans += count[pre] //非空美丽子字符串，代表不能与自身匹配，因此得先累加答案，再更新 count
		count[pre]++
	}
	return int64(ans)
}

func beautifulSubstrings3(s string, k int) int64 {
	// 因子分解
	for i := 1; ; i++ {
		if i*i%(4*k) == 0 {
			k = i
			break
		}
	}

	// 前缀和
	sum := make([]int, len(s)+1)
	sum[0] = 0 // 初始化
	for i, b := range s {
		switch b {
		case 'a', 'e', 'i', 'o', 'u':
			sum[i+1] = sum[i] + 1
		default:
			sum[i+1] = sum[i] - 1
		}
	}

	// L = i - j, L % k == 0 <=> (i-j) % k == 0 <=> i % k == j % k
	// 哈希表统计 (j % k == i % k, sum[j] == sun[i]) 的出现次数
	var ans int64
	type pair struct{ idx, sum int }
	m := make(map[pair]int64)
	for i, t := range sum {
		p := pair{
			idx: i % k,
			sum: t,
		}
		ans += m[p]
		m[p]++
	}
	return ans
}
