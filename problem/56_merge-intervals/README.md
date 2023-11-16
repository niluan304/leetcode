
| [English](README_EN.md) | 简体中文 |

# [56. 合并区间](https://leetcode.cn/problems/merge-intervals/)
Difficulty:Medium, Likes: 2170

## 题目描述

<p>以数组 <code>intervals</code> 表示若干个区间的集合，其中单个区间为 <code>intervals[i] = [start<sub>i</sub>, end<sub>i</sub>]</code> 。请你合并所有重叠的区间，并返回&nbsp;<em>一个不重叠的区间数组，该数组需恰好覆盖输入中的所有区间</em>&nbsp;。</p>

<p>&nbsp;</p>

<p><strong>示例 1：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[1,3],[2,6],[8,10],[15,18]]
<strong>输出：</strong>[[1,6],[8,10],[15,18]]
<strong>解释：</strong>区间 [1,3] 和 [2,6] 重叠, 将它们合并为 [1,6].
</pre>

<p><strong>示例&nbsp;2：</strong></p>

<pre>
<strong>输入：</strong>intervals = [[1,4],[4,5]]
<strong>输出：</strong>[[1,5]]
<strong>解释：</strong>区间 [1,4] 和 [4,5] 可被视为重叠区间。</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 &lt;= intervals.length &lt;= 10<sup>4</sup></code></li>
	<li><code>intervals[i].length == 2</code></li>
	<li><code>0 &lt;= start<sub>i</sub> &lt;= end<sub>i</sub> &lt;= 10<sup>4</sup></code></li>
</ul>


## 相关话题

- [数组](https://leetcode.cn/tag/array/)
- [排序](https://leetcode.cn/tag/sorting/)

## 相似题目

- [插入区间](../insert-interval/README.md) Medium 
- [会议室](../meeting-rooms/README.md) Easy 🔒
- [会议室 II](../meeting-rooms-ii/README.md) Medium 🔒
- [提莫攻击](../teemo-attacking/README.md) Easy 
- [给字符串添加加粗标签](../add-bold-tag-in-string/README.md) Medium 🔒
- [Range 模块](../range-module/README.md) Hard 
- [员工空闲时间](../employee-free-time/README.md) Hard 🔒
- [划分字母区间](../partition-labels/README.md) Medium 
- [区间列表的交集](../interval-list-intersections/README.md) Medium 
