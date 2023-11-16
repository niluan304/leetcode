
| [English](README_EN.md) | 简体中文 |

# [124. 二叉树中的最大路径和](https://leetcode.cn/problems/binary-tree-maximum-path-sum/)
Difficulty:Hard, Likes: 2100

## 题目描述

<p>二叉树中的<strong> 路径</strong> 被定义为一条节点序列，序列中每对相邻节点之间都存在一条边。同一个节点在一条路径序列中 <strong>至多出现一次</strong> 。该路径<strong> 至少包含一个 </strong>节点，且不一定经过根节点。</p>

<p><strong>路径和</strong> 是路径中各节点值的总和。</p>

<p>给你一个二叉树的根节点 <code>root</code> ，返回其 <strong>最大路径和</strong> 。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/exx1.jpg" style="width: 322px; height: 182px;" />
<pre>
<strong>输入：</strong>root = [1,2,3]
<strong>输出：</strong>6
<strong>解释：</strong>最优路径是 2 -&gt; 1 -&gt; 3 ，路径和为 2 + 1 + 3 = 6</pre>

<p><strong>示例 2：</strong></p>
<img alt="" src="https://assets.leetcode.com/uploads/2020/10/13/exx2.jpg" />
<pre>
<strong>输入：</strong>root = [-10,9,20,null,null,15,7]
<strong>输出：</strong>42
<strong>解释：</strong>最优路径是 15 -&gt; 20 -&gt; 7 ，路径和为 15 + 20 + 7 = 42
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li>树中节点数目范围是 <code>[1, 3 * 10<sup>4</sup>]</code></li>
	<li><code>-1000 &lt;= Node.val &lt;= 1000</code></li>
</ul>


## 相关话题

- [树](https://leetcode.cn/tag/tree/)
- [深度优先搜索](https://leetcode.cn/tag/depth-first-search/)
- [动态规划](https://leetcode.cn/tag/dynamic-programming/)
- [二叉树](https://leetcode.cn/tag/binary-tree/)

## 相似题目

- [路径总和](../path-sum/README.md) Easy 
- [求根节点到叶节点数字之和](../sum-root-to-leaf-numbers/README.md) Medium 
- [路径总和 IV](../path-sum-iv/README.md) Medium 🔒
- [最长同值路径](../longest-univalue-path/README.md) Medium 
