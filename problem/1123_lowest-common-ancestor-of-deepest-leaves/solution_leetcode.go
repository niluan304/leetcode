package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法一：递归递归，有递有归
//
// 看上图（示例 1），这棵树的节点 3,5,2 都是最深叶节点 7,4 的公共祖先，但只有节点 2 是最近的公共祖先。
//
// 上面视频中提到，如果我们要找的节点只在左子树中，那么最近公共祖先也必然只在左子树中。
// 对于本题，如果左子树的最大深度比右子树的大，那么最深叶结点就只在左子树中，所以最近公共祖先也只在左子树中。
//
// 如果左右子树的最大深度一样呢？当前节点一定是最近公共祖先吗？
//
// 不一定。比如节点 1 的左右子树最深叶节点 0,8 的深度都是 2，
// 但该深度并不是全局最大深度，所以节点 1 并不能是答案。
//
// 根据以上讨论，正确做法如下：
//
// 递归这棵二叉树，同时维护全局最大深度 maxDepth。
// 在「递」的时候往下传 depth，用来表示当前节点的深度。
// 在「归」的时候往上传当前子树最深叶节点的深度。
// 设左子树最深叶节点的深度为 leftMaxDepth，右子树最深叶节点的深度为 rightMaxDepth。
// 如果 leftMaxDepth=rightMaxDepth=maxDepth，那么更新答案为当前节点。
// 注意这并不代表我们找到了答案，如果后面发现了更深的叶节点，那么答案还会更新。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/solutions/2428724/liang-chong-di-gui-si-lu-pythonjavacgojs-xxnk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcodeEndlesscheng(root *TreeNode) (ans *TreeNode) {
	maxDepth := -1 // 全局最大深度
	var dfs func(*TreeNode, int) int
	dfs = func(node *TreeNode, depth int) int {
		if node == nil {
			maxDepth = max(maxDepth, depth) // 维护全局最大深度
			return depth
		}
		leftMaxDepth := dfs(node.Left, depth+1)   // 获取左子树最深叶节点的深度
		rightMaxDepth := dfs(node.Right, depth+1) // 获取右子树最深叶节点的深度
		if leftMaxDepth == rightMaxDepth && leftMaxDepth == maxDepth {
			ans = node
		}
		return max(leftMaxDepth, rightMaxDepth) // 当前子树最深叶节点的深度
	}
	dfs(root, 0)
	return
}

func max(a, b int) int {
	if b > a {
		return b
	}
	return a
}

// 方法二：自底向上
// 也可以不用全局变量，而是把每棵子树都看成是一个「子问题」，即对于每棵子树，我们需要知道：
//
// 这棵子树最深叶结点的深度。这里是指叶子在这棵子树内的深度，而不是在整棵二叉树的视角下的深度。相当于这棵子树的高度。
// 这棵子树的最深叶结点的最近公共祖先 lca。
// 分类讨论：
//
// 设子树的根节点为 node，node 的左子树的高度为 leftHeight，node 的右子树的高度为 rightHeight。
// 如果 leftHeight>rightHeight，那么子树的高度为 leftHeight+1，lca 是左子树的 lca。
// 如果 leftHeight<rightHeight，那么子树的高度为 rightHeight+1，lca 是右子树的 lca。
// 如果 leftHeight=rightHeight，那么子树的高度为 leftHeight+1，lca 就是 node。
// 反证法：如果 lca 在左子树中，那么 lca 不是右子树的最深叶结点的祖先，这不对；
// 如果 lca 在右子树中，那么 lca 不是左子树的最深叶结点的祖先，这也不对；
// 如果 lca 在 node 的上面，那就不符合「最近」的要求。所以 lca 只能是 node。
//
// 作者：灵茶山艾府
// 链接：https://leetcode.cn/problems/lowest-common-ancestor-of-deepest-leaves/solutions/2428724/liang-chong-di-gui-si-lu-pythonjavacgojs-xxnk/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcodeEndlesscheng2(root *TreeNode) *TreeNode {
	var dfs func(node *TreeNode) (int, *TreeNode)
	dfs = func(node *TreeNode) (int, *TreeNode) {
		if node == nil {
			return 0, nil
		}
		leftHeight, leftLCA := dfs(node.Left)
		rightHeight, rightLCA := dfs(node.Right)
		if leftHeight > rightHeight { // 左子树更高
			return leftHeight + 1, leftLCA
		}
		if leftHeight < rightHeight { // 右子树更高
			return rightHeight + 1, rightLCA
		}
		return leftHeight + 1, node // 一样高
	}

	_, lca := dfs(root)
	return lca
}
