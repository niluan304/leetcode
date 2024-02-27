
| [English](problem_en.md) | 简体中文 |

# [105. 从前序与中序遍历序列构造二叉树](https://leetcode.cn/problems/construct-binary-tree-from-preorder-and-inorder-traversal/)
Difficulty:Medium, Likes: 2210

## 题目描述

<p>给定两个整数数组&nbsp;<code>preorder</code> 和 <code>inorder</code>&nbsp;，其中&nbsp;<code>preorder</code> 是二叉树的<strong>先序遍历</strong>， <code>inorder</code>&nbsp;是同一棵树的<strong>中序遍历</strong>，请构造二叉树并返回其根节点。</p>

<p>&nbsp;</p>

<p><strong>示例 1:</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2021/02/19/tree.jpg" style="height: 302px; width: 277px;" />
<pre>
<strong>输入</strong><strong>:</strong> preorder = [3,9,20,15,7], inorder = [9,3,15,20,7]
<strong>输出:</strong> [3,9,20,null,null,15,7]
</pre>

<p><strong>示例 2:</strong></p>

<pre>
<strong>输入:</strong> preorder = [-1], inorder = [-1]
<strong>输出:</strong> [-1]
</pre>

<p>&nbsp;</p>

<p><strong>提示:</strong></p>

<ul>
	<li><code>1 &lt;= preorder.length &lt;= 3000</code></li>
	<li><code>inorder.length == preorder.length</code></li>
	<li><code>-3000 &lt;= preorder[i], inorder[i] &lt;= 3000</code></li>
	<li><code>preorder</code>&nbsp;和&nbsp;<code>inorder</code>&nbsp;均 <strong>无重复</strong> 元素</li>
	<li><code>inorder</code>&nbsp;均出现在&nbsp;<code>preorder</code></li>
	<li><code>preorder</code>&nbsp;<strong>保证</strong> 为二叉树的前序遍历序列</li>
	<li><code>inorder</code>&nbsp;<strong>保证</strong> 为二叉树的中序遍历序列</li>
</ul>


## 相关话题

    - [树](https://leetcode.cn/tag/tree/)
    - [数组](https://leetcode.cn/tag/array/)
    - [哈希表](https://leetcode.cn/tag/hash-table/)
    - [分治](https://leetcode.cn/tag/divide-and-conquer/)
    - [二叉树](https://leetcode.cn/tag/binary-tree/)

## 相似题目

    - [从中序与后序遍历序列构造二叉树](../construct-binary-tree-from-inorder-and-postorder-traversal/README.md) Medium 
