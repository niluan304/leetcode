
| English | [简体中文](problem_zh.md) |

# [745.Prefix and Suffix Search](https://leetcode.com/problems/prefix-and-suffix-search/)
Difficulty:Hard, Likes: 198

## Description

<p>Design a special dictionary that searches the words in it by a prefix and a suffix.</p>

<p>Implement the <code>WordFilter</code> class:</p>

<ul>
	<li><code>WordFilter(string[] words)</code> Initializes the object with the <code>words</code> in the dictionary.</li>
	<li><code>f(string pref, string suff)</code> Returns <em>the index of the word in the dictionary,</em> which has the prefix <code>pref</code> and the suffix <code>suff</code>. If there is more than one valid index, return <strong>the largest</strong> of them. If there is no such word in the dictionary, return <code>-1</code>.</li>
</ul>

<p>&nbsp;</p>
<p><strong class="example">Example 1:</strong></p>

<pre>
<strong>Input</strong>
[&quot;WordFilter&quot;, &quot;f&quot;]
[[[&quot;apple&quot;]], [&quot;a&quot;, &quot;e&quot;]]
<strong>Output</strong>
[null, 0]
<strong>Explanation</strong>
WordFilter wordFilter = new WordFilter([&quot;apple&quot;]);
wordFilter.f(&quot;a&quot;, &quot;e&quot;); // return 0, because the word at index 0 has prefix = &quot;a&quot; and suffix = &quot;e&quot;.
</pre>

<p>&nbsp;</p>
<p><strong>Constraints:</strong></p>

<ul>
	<li><code>1 &lt;= words.length &lt;= 10<sup>4</sup></code></li>
	<li><code>1 &lt;= words[i].length &lt;= 7</code></li>
	<li><code>1 &lt;= pref.length, suff.length &lt;= 7</code></li>
	<li><code>words[i]</code>, <code>pref</code> and <code>suff</code> consist of lowercase English letters only.</li>
	<li>At most <code>10<sup>4</sup></code> calls will be made to the function <code>f</code>.</li>
</ul>


## Related Topics

- [Design](https://leetcode.com/tag/design/)
- [Trie](https://leetcode.com/tag/trie/)
- [Hash Table](https://leetcode.com/tag/hash-table/)
- [String](https://leetcode.com/tag/string/)

## Similar Questions

- [Design Add and Search Words Data Structure](../design-add-and-search-words-data-structure/README_EN.md) Medium 
