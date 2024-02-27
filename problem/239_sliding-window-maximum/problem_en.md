
| English | [简体中文](problem_zh.md) |

# [239.Sliding Window Maximum](https://leetcode.com/problems/sliding-window-maximum/)
Difficulty:Hard, Likes: 2622

## Description

<p>You are given an array of integers&nbsp;<code>nums</code>, there is a sliding window of size <code>k</code> which is moving from the very left of the array to the very right. You can only see the <code>k</code> numbers in the window. Each time the sliding window moves right by one position.</p>

<p>Return <em>the max sliding window</em>.</p>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input:</strong> nums = [1,3,-1,-3,5,3,6,7], k = 3
<strong>Output:</strong> [3,3,5,5,6,7]
<strong>Explanation:</strong> 
Window position                Max
---------------               -----
[1  3  -1] -3  5  3  6  7       <strong>3</strong>
 1 [3  -1  -3] 5  3  6  7       <strong>3</strong>
 1  3 [-1  -3  5] 3  6  7      <strong> 5</strong>
 1  3  -1 [-3  5  3] 6  7       <strong>5</strong>
 1  3  -1  -3 [5  3  6] 7       <strong>6</strong>
 1  3  -1  -3  5 [3  6  7]      <strong>7</strong>
</pre>

<p><strong class="example">Example 2:</strong></p>

<pre>
<strong>Input:</strong> nums = [1], k = 1
<strong>Output:</strong> [1]
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= nums.length &lt;= 10<sup>5</sup></code></li>
	<li><code>-10<sup>4</sup> &lt;= nums[i] &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= k &lt;= nums.length</code></li>
</ul>


## Related Topics

- [Queue](https://leetcode.com/tag/queue/)
- [Array](https://leetcode.com/tag/array/)
- [Sliding Window](https://leetcode.com/tag/sliding-window/)
- [Monotonic Queue](https://leetcode.com/tag/monotonic-queue/)
- [Heap (Priority Queue)](https://leetcode.com/tag/heap-priority-queue/)

## Similar Questions

- [Minimum Window Substring](../minimum-window-substring/README_EN.md) Hard 
- [Min Stack](../min-stack/README_EN.md) Medium 
- [Longest Substring with At Most Two Distinct Characters](../longest-substring-with-at-most-two-distinct-characters/README_EN.md) Medium 🔒
- [Paint House II](../paint-house-ii/README_EN.md) Hard 🔒
