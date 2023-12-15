package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// #### 方法二：深度优先搜索
//
// **思路与算法**
//
// 同样的方法我们还可以使用深度优先搜索来遍历该二叉树，对奇数层进行反转。遍历过程如下：
// + 由于该二叉树是**完美二叉树**，因此我们可以知道对于**根**节点来说，它的孩子节点为第一层节点，此时左孩子需要与右孩子需要进行反转；
// + 当遍历每一层时，由于 $\textit{root}_1,\textit{root}_2$ 分别指向该层两个可能需要进行值交换的节点。根据**完美二叉树**的层次反转规则，即左边排第一的元素与倒数第一元素进行交换，第二个元素与倒数二个元素交换，此时 $\textit{root}_1$ 的左孩子与 $\textit{root}_2$ 的右孩子可能需要进行交换，$\textit{root}_1$ 的右孩子与 $\textit{root}_2$ 的左孩子可能需要进行交换。在遍历的同时按照上述规则，将配对的节点进行递归传递到下一层；
// + 我们用 $\text{isOdd}$ 来标记当前层次是否为奇数层，由于偶数层不需要进行交换，当 $\text{isOdd}$ 为 $\text{true}$ 时，表明当前需要交换，我们直接交换两个节点 $\textit{root}_1,\textit{root}_2$ 的值；
// + 由于**完美二叉树**来说，第 $i$ 的节点数目要么为 $2^i$ 个，要么为 $0$ 个，因此如果最左边的节点 $\textit{root}_1$ 为空时，则可以直接返回。
func leetcode2(root *TreeNode) *TreeNode {
	if root == nil { // 树中的节点数目在范围 [1, 2^14] 内
		return root
	}

	var dfs func(root1 *TreeNode, root2 *TreeNode, isOdd bool)
	dfs = func(root1 *TreeNode, root2 *TreeNode, isOdd bool) {
		if root1 == nil {
			return
		}
		if isOdd {
			root1.Val, root2.Val = root2.Val, root1.Val
		}
		dfs(root1.Left, root2.Right, !isOdd)
		dfs(root1.Right, root2.Left, !isOdd)
	}

	dfs(root.Left, root.Right, true)
	return root
}
