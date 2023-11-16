package main

import . "github.com/EndlessCheng/codeforces-go/leetcode/testutil"

// ## 方法一：DFS
//
// #### 前置知识
//
// [二叉树 DFS【基础算法精讲 09】](https://www.bilibili.com/video/BV1UD4y1Y769/)
//
// #### 思路
//
// ![lc117.png](https://pic.leetcode.cn/1698797451-YoHDtg-lc117.png)
//
// DFS 这棵树，从根节点 $1$ 出发，向左递归到 $2$，再向左递归到 $4$。
//
// 这三个节点正好是每一层的第一个节点（类似链表头），用一个数组 $\textit{pre}$ 记录，即 $\textit{pre}[0]$ 为节点 $1$，$\textit{pre}[1]$ 为节点 $2$，$\textit{pre}[2]$ 为节点 $4$。$\textit{pre}$ 的下标就是节点的**深度**。
//
// 继续递归到 $5$（深度为 $2$），从 $\textit{pre}[2]$ 中拿到节点 $4$，把 $4$ 的 $\textit{next}$ 指向 $5$。然后更新 $\textit{pre}[2]$ 为节点 $5$，这样在后面递归到节点 $7$ 时，就可以从 $\textit{pre}[2]$ 中拿到节点 $5$，把 $5$ 的 $\textit{next}$ 指向 $7$ 了。
//
// #### 算法
//
// 1. 创建一个空数组 $\textit{pre}$（因为一开始不知道二叉树有多深）。
// 2. DFS 这棵二叉树，递归参数为当前节点 $\textit{node}$，以及当前节点的深度 $\textit{depth}$。每往下递归一层，就把 $\textit{depth}$ 加一。
// 3. 如果 $\textit{depth}$ 等于 $\textit{pre}$ 数组的长度，说明 $\textit{node}$ 是这一层最左边的节点，把 $\textit{node}$ 添加到 $\textit{pre}$ 的末尾。
// 4. 否则，把 $\textit{pre}[\textit{depth}]$ 的 $\textit{next}$ 指向 $\textit{node}$，然后更新 $\textit{pre}[\textit{depth}]$ 为 $\textit{node}$。
// 5. 递归边界：如果 $\textit{node}$ 是空节点，直接返回。
// 6. 递归入口：$\textit{dfs}(\textit{root},0)$。
// 7. 最后返回 $\textit{root}$。
// ### 复杂度分析
//
// - 时间复杂度：$\mathcal{O}(n)$，其中 $n$ 为二叉树的节点个数。
// - 空间复杂度：$\mathcal{O}(1)$。只用到若干额外变量。
func endlesscheng1(root *Node) *Node {
	pre := []*Node{}
	var dfs func(*Node, int)
	dfs = func(node *Node, depth int) {
		if node == nil {
			return
		}
		if depth == len(pre) { // node 是这一层最左边的节点
			pre = append(pre, node)
		} else { // pre[depth] 是 node 左边的节点
			pre[depth].Next = node // node 左边的节点指向 node
			pre[depth] = node
		}
		dfs(node.Left, depth+1)
		dfs(node.Right, depth+1)
	}
	dfs(root, 0) // 根节点的深度为 0
	return root
}
