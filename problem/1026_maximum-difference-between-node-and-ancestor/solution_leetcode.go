package main

import (
	"math"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法二：「归」
// 方法一的思路是维护 B 的祖先节点中的最小值和最大值，我们还可以站在 A 的视角，维护 A 的子孙节点中的最小值 mn 和最大值 mx。
//
// 换句话说，最小值和最大值不再作为入参，而是作为返回值，意思是以 A 为根的子树中的最小值和最大值。
//
// 递归到节点 A 时，初始化 mn 和 mx 为 A.valA，然后递归左右子树，拿到左右子树的最小值和最大值，去更新 mn 和 mx，然后计算
// max(∣mn−A.val∣,∣mx−A.val∣)
// 并更新答案的最大值。
//
// 由于 mn≤A.val≤mx，上式可化简为
// max(A.val−mn,mx−A.val)
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/maximum-difference-between-node-and-ancestor/solutions/2232367/liang-chong-fang-fa-zi-ding-xiang-xia-zi-wj9v/
func leetcodEndlesscheng(root *TreeNode) (ans int) {
	var dfs func(*TreeNode) (int, int)
	dfs = func(node *TreeNode) (int, int) {
		if node == nil {
			return math.MaxInt, math.MinInt // 保证空节点不影响 mn 和 mx
		}
		mn, mx := node.Val, node.Val
		lMn, lMx := dfs(node.Left)
		rMn, rMx := dfs(node.Right)
		mn = min(mn, min(lMn, rMn))
		mx = max(mx, max(lMx, rMx))
		ans = max(ans, max(node.Val-mn, mx-node.Val))
		return mn, mx
	}
	dfs(root)
	return
}

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}
func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
