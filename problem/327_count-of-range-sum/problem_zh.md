
| [English](problem_en.md) | 简体中文 |

# [327. 区间和的个数](https://leetcode.cn/problems/count-of-range-sum/)
Difficulty:Hard, Likes: 581

## 题目描述

<p>给你一个整数数组 <code>nums</code> 以及两个整数 <code>lower</code> 和 <code>upper</code> 。求数组中，值位于范围 <code>[lower, upper]</code> （包含 <code>lower</code> 和 <code>upper</code>）之内的 <strong>区间和的个数</strong> 。</p>

<p><strong>区间和</strong> <code>S(i, j)</code> 表示在 <code>nums</code> 中，位置从 <code>i</code> 到 <code>j</code> 的元素之和，包含 <code>i</code> 和 <code>j</code> (<code>i</code> ≤ <code>j</code>)。</p>

<p> </p>
<strong>示例 1：</strong>

<pre>
<strong>输入：</strong>nums = [-2,5,-1], lower = -2, upper = 2
<strong>输出：</strong>3
<strong>解释：</strong>存在三个区间：[0,0]、[2,2] 和 [0,2] ，对应的区间和分别是：-2 、-1 、2 。
</pre>

<p><strong>示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [0], lower = 0, upper = 0
<strong>输出：</strong>1
</pre>

<p> </p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>1 <= nums.length <= 10<sup>5</sup></code></li>
	<li><code>-2<sup>31</sup> <= nums[i] <= 2<sup>31</sup> - 1</code></li>
	<li><code>-10<sup>5</sup> <= lower <= upper <= 10<sup>5</sup></code></li>
	<li>题目数据保证答案是一个 <strong>32 位</strong> 的整数</li>
</ul>


## 相关话题

    - [树状数组](https://leetcode.cn/tag/binary-indexed-tree/)
    - [线段树](https://leetcode.cn/tag/segment-tree/)
    - [数组](https://leetcode.cn/tag/array/)
    - [二分查找](https://leetcode.cn/tag/binary-search/)
    - [分治](https://leetcode.cn/tag/divide-and-conquer/)
    - [有序集合](https://leetcode.cn/tag/ordered-set/)
    - [归并排序](https://leetcode.cn/tag/merge-sort/)

## 相似题目

    - [计算右侧小于当前元素的个数](../count-of-smaller-numbers-after-self/README.md) Hard 
    - [翻转对](../reverse-pairs/README.md) Hard 
