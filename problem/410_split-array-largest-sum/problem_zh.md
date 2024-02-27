
| [English](problem_en.md) | 简体中文 |

# [410. 分割数组的最大值](https://leetcode.cn/problems/split-array-largest-sum/)
Difficulty:Hard, Likes: 871

## 题目描述

<p>给定一个非负整数数组 <code>nums</code> 和一个整数&nbsp;<code>k</code> ，你需要将这个数组分成&nbsp;<code>k</code><em>&nbsp;</em>个非空的连续子数组。</p>

<p>设计一个算法使得这&nbsp;<code>k</code><em>&nbsp;</em>个子数组各自和的最大值最小。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [7,2,5,10,8], k = 2
<strong>输出：</strong>18
<strong>解释：</strong>
一共有四种方法将 nums 分割为 2 个子数组。 
其中最好的方式是将其分为 [7,2,5] 和 [10,8] 。
因为此时这两个子数组各自的和的最大值为18，在所有情况中最小。</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3,4,5], k = 2
<strong>输出：</strong>9
</pre>

<p><strong>示例 3：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,4,4], k = 3
<strong>输出：</strong>4
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 1000</code></li>
	<li><code>0 &lt;= nums[i] &lt;= 10<sup>6</sup></code></li>
	<li><code>1 &lt;= k &lt;= min(50, nums.length)</code></li>
</ul>


## 相关话题

    - [贪心](https://leetcode.cn/tag/greedy/)
    - [数组](https://leetcode.cn/tag/array/)
    - [二分查找](https://leetcode.cn/tag/binary-search/)
    - [动态规划](https://leetcode.cn/tag/dynamic-programming/)
    - [前缀和](https://leetcode.cn/tag/prefix-sum/)

## 相似题目

