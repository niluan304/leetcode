
| English | [简体中文](problem_zh.md) |

# [493. Reverse Pairs](https://leetcode.com/problems/reverse-pairs/)
Difficulty:Hard, Likes: 439

## Description

<p>Given an integer array <code>nums</code>, return <em>the number of <strong>reverse pairs</strong> in the array</em>.</p>

<p>A <strong>reverse pair</strong> is a pair <code>(i, j)</code> where:</p>

<ul>
	<li><code>0 &lt;= i &lt; j &lt; nums.length</code> and</li>
	<li><code>nums[i] &gt; 2 * nums[j]</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,2,3,1]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The reverse pairs are:
(1, 4) --&gt; nums[1] = 3, nums[4] = 1, 3 &gt; 2 * 1
(3, 4) --&gt; nums[3] = 3, nums[4] = 1, 3 &gt; 2 * 1
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [2,4,3,5,1]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The reverse pairs are:
(1, 4) --&gt; nums[1] = 4, nums[4] = 1, 4 &gt; 2 * 1
(2, 4) --&gt; nums[2] = 3, nums[4] = 1, 3 &gt; 2 * 1
(3, 4) --&gt; nums[3] = 5, nums[4] = 1, 5 &gt; 2 * 1
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 5 * 10<sup>4</sup></code></li>
	<li><code>-2<sup>31</sup> &lt;= nums[i] &lt;= 2<sup>31</sup> - 1</code></li>
</ul>


## Related Topics

- [Binary Indexed Tree](https://leetcode.com/tag/binary-indexed-tree/)
- [Segment Tree](https://leetcode.com/tag/segment-tree/)
- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)
- [Divide and Conquer](https://leetcode.com/tag/divide-and-conquer/)
- [Ordered Set](https://leetcode.com/tag/ordered-set/)
- [Merge Sort](https://leetcode.com/tag/merge-sort/)

## Similar Questions

- [Count of Smaller Numbers After Self](../count-of-smaller-numbers-after-self/README_EN.md) Hard 
- [Count of Range Sum](../count-of-range-sum/README_EN.md) Hard 
