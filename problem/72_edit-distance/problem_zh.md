
| [English](problem_en.md) | 简体中文 |

# [72. 编辑距离](https://leetcode.cn/problems/edit-distance/)
Difficulty:Medium, Likes: 3205

## 题目描述

<p>给你两个单词&nbsp;<code>word1</code> 和&nbsp;<code>word2</code>， <em>请返回将&nbsp;<code>word1</code>&nbsp;转换成&nbsp;<code>word2</code> 所使用的最少操作数</em> &nbsp;。</p>

<p>你可以对一个单词进行如下三种操作：</p>

<ul>
	<li>插入一个字符</li>
	<li>删除一个字符</li>
	<li>替换一个字符</li>
</ul>

<p>&nbsp;</p>

<p><strong>示例&nbsp;1：</strong></p>

<pre>
<strong>输入：</strong>word1 = "horse", word2 = "ros"
<strong>输出：</strong>3
<strong>解释：</strong>
horse -&gt; rorse (将 'h' 替换为 'r')
rorse -&gt; rose (删除 'r')
rose -&gt; ros (删除 'e')
</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre>
<strong>输入：</strong>word1 = "intention", word2 = "execution"
<strong>输出：</strong>5
<strong>解释：</strong>
intention -&gt; inention (删除 't')
inention -&gt; enention (将 'i' 替换为 'e')
enention -&gt; exention (将 'n' 替换为 'x')
exention -&gt; exection (将 'n' 替换为 'c')
exection -&gt; execution (插入 'u')
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>0 &lt;= word1.length, word2.length &lt;= 500</code></li>
	<li><code>word1</code> 和 <code>word2</code> 由小写英文字母组成</li>
</ul>


## 相关话题

- [字符串](https://leetcode.cn/tag/string/)
- [动态规划](https://leetcode.cn/tag/dynamic-programming/)

## 相似题目

- [相隔为 1 的编辑距离](../one-edit-distance/README.md) Medium 🔒
- [两个字符串的删除操作](../delete-operation-for-two-strings/README.md) Medium 
- [两个字符串的最小ASCII删除和](../minimum-ascii-delete-sum-for-two-strings/README.md) Medium 
- [不相交的线](../uncrossed-lines/README.md) Medium 
