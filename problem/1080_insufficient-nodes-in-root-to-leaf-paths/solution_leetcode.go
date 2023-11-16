package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 一、思考
// 对于一个叶子节点，要想删除它，需要满足什么条件？
//
// 对于一个非叶节点，如果它有一个儿子没被删除，那么它能被删除吗？如果它的儿子都被删除，意味着什么？
//
// 二、解惑
// 对于一个叶子节点 leaf，由于根到 leaf 的路径仅有一条，所以如果这条路径的元素和小于 limit，就删除 leaf。
//
// 对于一个非叶节点 node，如果 node 有一个儿子没被删除，那么 node 就不能被删除。这可以用反证法证明：假设可以把 node 删除，那么经过 node 的所有路径和都小于 limit，也就意味着经过 node 的儿子的路径和也小于 limit，说明 node 的儿子需要被删除，矛盾，所以 node 不能被删除。
//
// 如果 node 的儿子都被删除，说明经过 node 的所有儿子的路径和都小于 limit，这等价于经过 node 的所有路径和都小于 limit，所以 node 需要被删除。
//
// 因此，要删除非叶节点 node，当且仅当 node 的所有儿子都被删除。
//
// 三、算法
// 一个直接的想法是，添加一个递归参数 sumPath，表示从根到当前节点的路径和。
//
// 但为了能直接调用 sufficientSubset，还可以从 limit 中减去当前节点值。
//
// 如果当前节点是叶子，且此时 limit>0，说明从根到这个叶子的路径和小于 limit，那么删除这个叶子。
//
// 如果当前节点不是叶子，那么往下递归，更新它的左儿子为对左儿子调用 sufficientSubset 的结果，更新它的右儿子为对右儿子调用 sufficientSubset 的结果。
//
// 如果左右儿子都为空，那么就删除当前节点，返回空；否则不删，返回当前节点。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/insufficient-nodes-in-root-to-leaf-paths/solutions/2278769/jian-ji-xie-fa-diao-yong-zi-shen-pythonj-64lf/
func leetcodEndlesscheng(root *TreeNode, limit int) *TreeNode {
	if root == nil {
		return nil
	}
	limit -= root.Val
	if root.Left == root.Right { // root 是叶子
		if limit > 0 { // 从根到叶子的路径和小于 limit，删除叶子
			return nil
		}
		return root // 否则不删除
	}
	root.Left = leetcodEndlesscheng(root.Left, limit)
	root.Right = leetcodEndlesscheng(root.Right, limit)
	if root.Left == nil && root.Right == nil { // 如果儿子都被删除，就删 root
		return nil
	}
	return root // 否则不删 root
}
