package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法一：深度优先搜索
// 思路与算法
//
// 如果我们知道了左子树和右子树的最大深度 lll 和 rrr，那么该二叉树的最大深度即为
// max(l,r)+1
// 而左子树和右子树的最大深度又可以以同样的方式进行计算。
//
// 因此我们可以用「深度优先搜索」的方法来计算二叉树的最大深度。
// 具体而言，在计算当前二叉树的最大深度时，可以先递归计算出其左子树和右子树的最大深度，
// 然后在 O(1) 时间内计算出当前二叉树的最大深度。递归在访问到空节点时退出。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/solutions/349250/er-cha-shu-de-zui-da-shen-du-by-leetcode-solution/
func leetcode1(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return max(leetcode1(root.Left), leetcode1(root.Right)) + 1
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// 方法二：广度优先搜索
// 思路与算法
//
// 我们也可以用「广度优先搜索」的方法来解决这道题目，但我们需要对其进行一些修改，
// 此时我们广度优先搜索的队列里存放的是「当前层的所有节点」。
// 每次拓展下一层的时候，不同于广度优先搜索的每次只从队列里拿出一个节点，
// 我们需要将队列里的所有节点都拿出来进行拓展，这样能保证每次拓展完的时候队列里存放的是当前层的所有节点，
// 即我们是一层一层地进行拓展，最后我们用一个变量 ans 来维护拓展的次数，该二叉树的最大深度即为 ans。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/maximum-depth-of-binary-tree/solutions/349250/er-cha-shu-de-zui-da-shen-du-by-leetcode-solution/
func leetcode2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{}
	queue = append(queue, root)
	ans := 0
	for len(queue) > 0 {
		sz := len(queue)
		for sz > 0 {
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			sz--
		}
		ans++
	}
	return ans
}
