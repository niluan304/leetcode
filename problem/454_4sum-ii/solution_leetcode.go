package main

// 方法一：分组 + 哈希表
// 思路与算法
//
// 我们可以将四个数组分成两部分，A 和 B 为一组，C 和 D 为另外一组。
//
// 对于 A 和 B，我们使用二重循环对它们进行遍历，得到所有 A[i]+B[j] 的值并存入哈希映射中。
// 对于哈希映射中的每个键值对，每个键表示一种 A[i]+B[j]，对应的值为 A[i]+B[j] 出现的次数。
//
// 对于 C 和 D，我们同样使用二重循环对它们进行遍历。当遍历到 C[k]+D[l]] 时，
// 如果 −(C[k]+D[l]) 出现在哈希映射中，那么将 −(C[k]+D[l]) 对应的值累加进答案中。
//
// 最终即可得到满足 A[i]+B[j]+C[k]+D[l]=0 的四元组数目。
func leetcode1(a, b, c, d []int) (ans int) {
	countAB := map[int]int{}
	for _, v := range a {
		for _, w := range b {
			countAB[v+w]++
		}
	}
	for _, v := range c {
		for _, w := range d {
			ans += countAB[-v-w]
		}
	}
	return
}
