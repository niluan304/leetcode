
| [English](problem_en.md) | 简体中文 |

# [301. 删除无效的括号](https://leetcode.cn/problems/remove-invalid-parentheses/)
Difficulty:Hard, Likes: 910

## 题目描述

<p>给你一个由若干括号和字母组成的字符串 <code>s</code> ，删除最小数量的无效括号，使得输入的字符串有效。</p>

<p>返回所有可能的结果。答案可以按 <strong>任意顺序</strong> 返回。</p>

<p> </p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>s = "()())()"
<strong>输出：</strong>["(())()","()()()"]
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>s = "(a)())()"
<strong>输出：</strong>["(a())()","(a)()()"]
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>s = ")("
<strong>输出：</strong>[""]
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= s.length <= 25</code></li>
	<li><code>s</code> 由小写英文字母以及括号 <code>'('</code> 和 <code>')'</code> 组成</li>
	<li><code>s</code> 中至多含 <code>20</code> 个括号</li>
</ul>


## 相关话题

    - [广度优先搜索](https://leetcode.cn/tag/breadth-first-search/)
    - [字符串](https://leetcode.cn/tag/string/)
    - [回溯](https://leetcode.cn/tag/backtracking/)

## 相似题目

    - [有效的括号](../valid-parentheses/README.md) Easy 