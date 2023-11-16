
| English | [简体中文](README.md) |

# [209.Minimum Size Subarray Sum](https://leetcode.com/problems/minimum-size-subarray-sum/)
Difficulty:Medium, Likes: 1961

## Description

<p>Given an array of positive integers <code>nums</code> and a positive integer <code>target</code>, return <em>the <strong>minimal length</strong> of a </em><span data-keyword="subarray-nonempty"><em>subarray</em></span><em> whose sum is greater than or equal to</em> <code>target</code>. If there is no such subarray, return <code>0</code> instead.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> target = 7, nums = [2,3,1,2,4,3]
<strong>Output:</strong> 2
<strong>Explanation:</strong> The subarray [4,3] has the minimal length under the problem constraint.
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> target = 4, nums = [1,4,4]
<strong>Output:</strong> 1
</pre>

<p><strong class="example">Example 3:</strong></p>

<pre>
<strong>Input:</strong> target = 11, nums = [1,1,1,1,1,1,1,1]
<strong>Output:</strong> 0
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= target &lt;= 10<sup>9</sup></code></li>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>1 &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
</ul>

<p>&nbsp;</p>
<strong>Follow up:</strong> If you have figured out the <code>O(n)</code> solution, try coding another solution of which the time complexity is <code>O(n log(n))</code>.

## Related Topics

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)
- [Prefix Sum](https://leetcode.com/tag/prefix-sum/)
- [Sliding Window](https://leetcode.com/tag/sliding-window/)

## Similar Questions

- [Minimum Window Substring](../minimum-window-substring/README_EN.md) Hard 
- [Maximum Size Subarray Sum Equals k](../maximum-size-subarray-sum-equals-k/README_EN.md) Medium 🔒
- [Maximum Length of Repeated Subarray](../maximum-length-of-repeated-subarray/README_EN.md) Medium 
