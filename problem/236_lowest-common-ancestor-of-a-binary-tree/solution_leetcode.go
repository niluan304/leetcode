package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法一：递归
// 思路和算法
//
// 我们递归遍历整棵二叉树，定义 fx 表示 x 节点的子树中是否包含 p 节点或 q 节点，如果包含为 true，否则为 false。那么符合条件的最近公共祖先 x 一定满足如下条件：
//
// (flson&&frson)∣∣((x=p∣∣x=q)&&(flson∣∣frson))
// 其中 lson 和 rson 分别代表 x 节点的左孩子和右孩子。
//
// 初看可能会感觉条件判断有点复杂，我们来一条条看，flson&&frson 说明左子树和右子树均包含 p 节点或 q 节点，
// 如果左子树包含的是 p 节点，那么右子树只能包含 q 节点，反之亦然，因为 p 节点和 q 节点都是不同且唯一的节点，
// 因此如果满足这个判断条件即可说明 x 就是我们要找的最近公共祖先。
//
// 再来看第二条判断条件，这个判断条件即是考虑了 x 恰好是 p 节点或 q 节点
// 且它的左子树或右子树有一个包含了另一个节点的情况，因此如果满足这个判断条件亦可说明 x 就是我们要找的最近公共祖先。
//
// 你可能会疑惑这样找出来的公共祖先深度是否是最大的。其实是最大的，因为我们是自底向上从叶子节点开始更新的，
// 所以在所有满足条件的公共祖先中一定是深度最大的祖先先被访问到，且由于 fx 本身的定义很巧妙，在找到最近公共祖先 x 以后，
// fx 按定义被设置为 true ，即假定了这个子树中只有一个 p 节点或 q 节点，因此其他公共祖先不会再被判断为符合条件。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/lowest-common-ancestor-of-a-binary-tree/solutions/238552/er-cha-shu-de-zui-jin-gong-gong-zu-xian-by-leetc-2/
func leetcode1(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	left := leetcode1(root.Left, p, q)
	right := leetcode1(root.Right, p, q)

	if left == nil {
		return right
	}
	if right == nil {
		return left
	}

	return root
}
