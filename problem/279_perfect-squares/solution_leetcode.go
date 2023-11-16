package main

import "math"

// 方法二：数学
// 思路及算法
//
// 一个数学定理可以帮助解决本题：「四平方和定理」。
//
// 四平方和定理证明了任意一个正整数都可以被表示为至多四个正整数的平方和。这给出了本题的答案的上界。
//
// 同时四平方和定理包含了一个更强的结论：当且仅当 n≠4k×(8m+7) 时，n 可以被表示为至多三个正整数的平方和。
// 因此，当 n=4k×(8m+7)n = 4^k 时，n 只能被表示为四个正整数的平方和。此时我们可以直接返回 4。
//
// 当 n≠4k×(8m+7)n 时，我们需要判断到底多少个完全平方数能够表示 n，我们知道答案只会是 1,2,3 中的一个：
//
// 答案为 1 时，则必有 n 为完全平方数，这很好判断；
//
// 答案为 2 时，则有 n=a，我们只需要枚举所有的 a(1≤a≤n)，判断n-a^2 是否为完全平方数即可；
//
// 答案为 3 时，我们很难在一个优秀的时间复杂度内解决它，但我们只需要检查答案为 1 或 2 的两种情况，即可利用排除法确定答案。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/perfect-squares/solutions/822940/wan-quan-ping-fang-shu-by-leetcode-solut-t99c/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode2(n int) int {
	if isPerfectSquare(n) {
		return 1
	}
	if checkAnswer4(n) {
		return 4
	}
	for i := 1; i*i <= n; i++ {
		j := n - i*i
		if isPerfectSquare(j) {
			return 2
		}
	}
	return 3
}

// 判断是否为完全平方数
func isPerfectSquare(x int) bool {
	y := int(math.Sqrt(float64(x)))
	return y*y == x
}

// 判断是否能表示为 4^k*(8m+7)
func checkAnswer4(x int) bool {
	for x%4 == 0 {
		x /= 4
	}
	return x%8 == 7
}
