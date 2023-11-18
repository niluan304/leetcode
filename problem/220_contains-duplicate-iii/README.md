
| [English](README_EN.md) | 简体中文 |

# [220. 存在重复元素 III](https://leetcode.cn/problems/contains-duplicate-iii/)
Difficulty:Hard, Likes: 717

## 题目描述

<p>给你一个整数数组 <code>nums</code> 和两个整数 <code>indexDiff</code> 和 <code>valueDiff</code> 。</p>

<p>找出满足下述条件的下标对 <code>(i, j)</code>：</p>

<ul>
	<li><code>i != j</code>,</li>
	<li><code>abs(i - j) &lt;= indexDiff</code></li>
	<li><code>abs(nums[i] - nums[j]) &lt;= valueDiff</code></li>
</ul>

<p>如果存在，返回 <code>true</code><em> ；</em>否则，返回<em> </em><code>false</code><em> </em>。</p>

<p>&nbsp;</p>

<p><strong class="example">示例 1：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,2,3,1], indexDiff = 3, valueDiff = 0
<strong>输出：</strong>true
<strong>解释：</strong>可以找出 (i, j) = (0, 3) 。
满足下述 3 个条件：
i != j --&gt; 0 != 3
abs(i - j) &lt;= indexDiff --&gt; abs(0 - 3) &lt;= 3
abs(nums[i] - nums[j]) &lt;= valueDiff --&gt; abs(1 - 1) &lt;= 0
</pre>

<p><strong class="example">示例 2：</strong></p>

<pre>
<strong>输入：</strong>nums = [1,5,9,1,5,9], indexDiff = 2, valueDiff = 3
<strong>输出：</strong>false
<strong>解释：</strong>尝试所有可能的下标对 (i, j) ，均无法满足这 3 个条件，因此返回 false 。
</pre>

<p>&nbsp;</p>

<p><strong>提示：</strong></p>

<ul>
	<li><code>2 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>9</sup> &lt;= nums[i] &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= indexDiff &lt;= nums.length</code></li>
	<li><code>0 &lt;= valueDiff &lt;= 10<sup>9</sup></code></li>
</ul>


## 相关话题

- [数组](https://leetcode.cn/tag/array/)
- [桶排序](https://leetcode.cn/tag/bucket-sort/)
- [有序集合](https://leetcode.cn/tag/ordered-set/)
- [排序](https://leetcode.cn/tag/sorting/)
- [滑动窗口](https://leetcode.cn/tag/sliding-window/)

## 相似题目

- [存在重复元素](../contains-duplicate/README.md) Easy 
- [存在重复元素 II](../contains-duplicate-ii/README.md) Easy 
