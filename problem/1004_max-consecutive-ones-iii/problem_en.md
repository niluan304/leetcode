
| English | [简体中文](problem_zh.md) |

# [1004.Max Consecutive Ones III](https://leetcode.com/problems/max-consecutive-ones-iii/)
Difficulty:Medium, Likes: 655

## Description

<p>Given a binary array <code>nums</code> and an integer <code>k</code>, return <em>the maximum number of consecutive </em><code>1</code><em>&#39;s in the array if you can flip at most</em> <code>k</code> <code>0</code>&#39;s.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,1,1,0,0,0,1,1,1,1,0], k = 2
<strong>Output:</strong> 6
<strong>Explanation:</strong> [1,1,1,0,0,<u><strong>1</strong>,1,1,1,1,<strong>1</strong></u>]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [0,0,1,1,0,0,1,1,1,0,1,1,0,0,0,1,1,1,1], k = 3
<strong>Output:</strong> 10
<strong>Explanation:</strong> [0,0,<u>1,1,<strong>1</strong>,<strong>1</strong>,1,1,1,<strong>1</strong>,1,1</u>,0,0,0,1,1,1,1]
Bolded numbers were flipped from 0 to 1. The longest subarray is underlined.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>nums[i]</code> is either <code>0</code> or <code>1</code>.</li>
	<li><code>0 &lt;= k &lt;= nums.length</code></li>
</ul>


## Related Topics

- [Array](https://leetcode.com/tag/array/)
- [Binary Search](https://leetcode.com/tag/binary-search/)
- [Prefix Sum](https://leetcode.com/tag/prefix-sum/)
- [Sliding Window](https://leetcode.com/tag/sliding-window/)

## Similar Questions

- [Longest Substring with At Most K Distinct Characters](../longest-substring-with-at-most-k-distinct-characters/README_EN.md) Medium 🔒
- [Longest Repeating Character Replacement](../longest-repeating-character-replacement/README_EN.md) Medium 
- [Max Consecutive Ones](../max-consecutive-ones/README_EN.md) Easy 
- [Max Consecutive Ones II](../max-consecutive-ones-ii/README_EN.md) Medium 🔒
