
| English | [简体中文](problem_zh.md) |

# [41. First Missing Positive](https://leetcode.com/problems/first-missing-positive/)
Difficulty:Hard, Likes: 2081

## Description

<p>Given an unsorted integer array <code>nums</code>. Return the <em>smallest positive integer</em> that is <em>not present</em> in <code>nums</code>.</p>

<p>You must implement an algorithm that runs in <code>O(n)</code> time and uses <code>O(1)</code> auxiliary space.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,2,0]
<strong>Output:</strong> 3
<strong>Explanation:</strong> The numbers in the range [1,2] are all in the array.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [3,4,-1,1]
<strong>Output:</strong> 2
<strong>Explanation:</strong> 1 is in the array but 2 is missing.
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> nums = [7,8,9,11,12]
<strong>Output:</strong> 1
<strong>Explanation:</strong> The smallest positive integer 1 is missing.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-2<sup>31</sup> &lt;= nums[i] &lt;= 2<sup>31</sup> - 1</code></li>
</ul>


## Related Topics

- [Array](https://leetcode.com/tag/array/)
- [Hash Table](https://leetcode.com/tag/hash-table/)

## Similar Questions

- [Missing Number](../missing-number/README_EN.md) Easy 
- [Find the Duplicate Number](../find-the-duplicate-number/README_EN.md) Medium 
- [Find All Numbers Disappeared in an Array](../find-all-numbers-disappeared-in-an-array/README_EN.md) Easy 
- [Couples Holding Hands](../couples-holding-hands/README_EN.md) Hard 
