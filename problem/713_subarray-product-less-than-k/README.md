
| [English](README_EN.md) | 简体中文 |

# [713. 乘积小于 K 的子数组](https://leetcode.cn/problems/subarray-product-less-than-k/)
Difficulty:Medium, Likes: 715

## 题目描述

给你一个整数数组 <code>nums</code> 和一个整数 <code>k</code> ，请你返回子数组内所有元素的乘积严格小于<em> </em><code>k</code> 的连续子数组的数目。
<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [10,5,2,6], k = 100
<strong>输出：</strong>8
<strong>解释：</strong>8 个乘积小于 100 的子数组分别为：[10]、[5]、[2],、[6]、[10,5]、[5,2]、[2,6]、[5,2,6]。
需要注意的是 [10,5,2] 并不是乘积小于 100 的子数组。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3], k = 0
<strong>输出：</strong>0</pre>

<p>&nbsp;</p>

<p><strong>提示:&nbsp;</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 3 * 10<sup>4</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 1000</code></li>
	<li><code>0 &lt;= k &lt;= 10<sup>6</sup></code></li>
</ul>


## 相关话题

- [数组](https://leetcode-cn.com/tag/array/)
- [滑动窗口](https://leetcode-cn.com/tag/sliding-window/)

## 相似题目

- [乘积最大子数组](../maximum-product-subarray/README.md) Medium 
- [和等于 k 的最长子数组长度](../maximum-size-subarray-sum-equals-k/README.md) Medium 🔒
- [和为 K 的子数组](../subarray-sum-equals-k/README.md) Medium 
- [小于 K 的两数之和](../two-sum-less-than-k/README.md) Easy 🔒
