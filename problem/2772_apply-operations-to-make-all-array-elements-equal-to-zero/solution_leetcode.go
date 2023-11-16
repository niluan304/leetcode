package main

// [视频讲解](https://b23.tv/8g0pnxs) 第四题。
//
// ## 提示 1
//
// 想一想，如果 $\\textit{nums}[0]>0$，我们必须要执行什么样的操作，才能让 $\\textit{nums}[0]=0$？
//
// ## 提示 2
//
// 对于 $\\textit{nums}[0]>0$ 的情况，必须把 $\\textit{nums}[0]$ 到 $\\textit{nums}[k-1]$ 都减去 $\\textit{nums}[0]$。
//
// 然后思考 $\\textit{nums}[1]$ 要怎么处理，依此类推。
//
// ## 提示 3
//
// 子数组同时加上/减去一个数，非常适合用**差分数组**来维护。
//
// 设差分数组为 $d$。那么把 $\\textit{nums}[i]$ 到 $\\textit{nums}[i+k-1]$ 同时减去 $1$，等价于把 $d[i]$ 减 $1$，$d[i+k]$ 加 $1$。
//
// 注意子数组长度必须恰好等于 $k$，所以当 $i+k\\le n$ 时，才能执行上述操作。
//
// 遍历数组的同时，用变量 $\\textit{sumd}$ 累加差分值。遍历到 $\\textit{nums}[i]$ 时，$\\textit{nums}[i]+\\textit{sumd}$ 就是 $\\textit{nums}[i]$ 的实际值了。
//
// 分类讨论：
//
// - 如果 $\\textit{nums}[i]<0$，由于无法让元素值增大，返回 `false`。
// - 如果 $\\textit{nums}[i]=0$，无需操作，遍历下一个数。
// - 如果 $\\textit{nums}[i]>0$：
//   - 如果 $i+k> n$，无法执行操作，所以 $\\textit{nums}[i]$ 无法变成 $0$，返回 `false`。
//   - 如果 $i+k\\le n$，按照上面说的执行操作，修改差分数组，遍历下一个数。
//
// 如果遍历中途没有返回 `false`，那么最后返回 `true`。
//
// #### 复杂度分析
//
// - 时间复杂度：$\\mathcal{o}(n)$，其中 $n$ 为 $\\textit{nums}$ 的长度。
// - 空间复杂度：$\\mathcal{o}(n)$。
//
// #### 相似题目
//
// - [2528. 最大化城市的最小供电站数目](https://leetcode.cn/problems/maximize-the-minimum-powered-city/)
func leetcodeEndlesscheng(nums []int, k int) bool {
	n := len(nums)
	d := make([]int, n+1)
	sumD := 0
	for i, x := range nums {
		sumD += d[i]
		x += sumD
		if x == 0 { // 无需操作
			continue
		}
		if x < 0 || i+k > n { // 无法操作
			return false
		}
		sumD -= x // 直接加到 sumD 中
		d[i+k] += x
	}
	return true
}
