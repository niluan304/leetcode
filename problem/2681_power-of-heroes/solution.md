记一次解题过程：通过观察猜到了递推公式

# 读题

根据「这组英雄的力量」的定义，得知元素顺序不影响答案，那么可以先排序，然后固定最大值。

暴力解法：全排列小于当前最大值的数，再遍历最小值，但是这种做法时间复杂度过高，没办法优化。

> 优化暴力解法，也是一种解题方法。



读题时有两点必须要注意到：
- 数据范围： $1 \leq nums.length \leq 10^5$，表明本题时间复杂度至多为 $\mathcal{O}(n\log n)$;
- 由于答案可能非常大，请你将结果对 $10^9 + 7$  **取余。**



根据我以往的做题经验，猜测本题的解法是动态规划。


# 观察与猜测

没什么思路的时，我喜欢在草稿纸上写例子，比如有 $nums = [1,2,3,4,5]$，顺序从小到大。

> **以下思考过程是有问题的**，因为按子序列个数的划分方式并不方便找到规律。

设当前最大值为 $a_n$，$[a_1,a_2,...,a_n]$ 的全排列后的最小值之和为 $S_n$，那么答案为：$\sum{a_n^2 * S_n}$

$a_1 = 1, S_1 = [1]$

$a_2=2, S_2 = [ 2 + (1)]$

- 解释：只有 $1$ 个数时：$\sum\min{[2]} = 2$，$2$ 个数时：$\sum\min{[2,1]} = 1$

$a_3 = 3, S_3 = [3 + (2+1) + (1)]$

- 解释：只有 $1$ 个数时：$\sum\min{[3]} = 3$，$2$ 个数时：$\sum\min{[3,2]} + \sum\min{[3,1]} = 2 + 1$，$3$ 个数：$\sum\min{[3,2,1]} = 1$

$a_4 = 4, S_4 = [4 + (3+2+1) + (2+1+1) + (1)]$

$a_5 = 5, S_5 = [5 + (4+3+2+1) + ...]$

## 再观察

为了找到规律，我尝试 **在 $S_n$ 内逐个固定最大值，再分组** ：

$a_1 = 1, S_1 = 1$

$a_2=2, S_2 = 2 + (1)$

$a_3 = 3, S_3 = 3 + (2+1) + (1)$

$a_4 = 4, S_4 =  4 + [3+(2+1)+1] + [2+(1)] + (1)$

$a_5 = 5, S_5 = 5 + \{4 + [3+(2+1)+1] + [2+(1)] + (1)\} + ...$

我发现式子有以下特点：

$a_1 = 1, S_1 = a_1$

$a_2=2, S_2 = a_2 + S_1$

$a_3 = 3, S_3 = a_3 + S_2 + S_1$

$a_4 = 4, S_4 = a_4 + S_3 + S_2 + S_1$

$a_5 = 5, S_5 = a_5 + S_4 + S_3 + S_2 + S_1$

根据这样的规划，容易猜测出：$S_n = a_n + \sum{S_{i-1}}$

根据错位相减法或换元法，可以求得：$S_n = a_n - a_{n-1} + 2*S_{n-1}$



# 代码实现

下面就是编码验证环节：

```go []
func sumOfPower(nums []int) int {
    const mod = 1_000_000_007
	nums = append(nums, 0) // 哨兵，防止nums[i-1]时数组越界
	sort.Ints(nums)
	var n = len(nums)
	var dp, power = make([]int, n+1), 0

	for i := 1; i < n; i++ {
		num := nums[i]
		dp[i] = (num - nums[i-1] + 2*dp[i-1]%mod) % mod
		power = (power + (num*num%mod)*dp[i]) % mod
	}
	return power
}
```
```python []
class Solution(object):
    def sumOfPower(self, nums: List[int]) -> int:
        MOD = 10**9 + 7
        nums.append(0) # 哨兵，防止nums[i-1]时数组越界
        nums.sort()

        dp, power = [0] * len(nums), 0
        for i in range(1, len(nums)):
            num = nums[i]
            dp[i] = (num - nums[i-1] + 2 * dp[i-1] % MOD) % MOD
            power = (power + (num * num % MOD) * dp[i]) % MOD

        return power

```

```javascript []
var sumOfPower = function(nums) {
  const mod = BigInt(1e9 + 7);
  nums.push(0); // 哨兵，防止nums[i-1]时数组越界
  nums.sort((a, b) => a - b);

  const n = nums.length;
  const dp = Array(n).fill(BigInt(0));
  let power = BigInt(0);

  for (let i = 1; i < n; i++) {
    const num = BigInt(nums[i]);
    dp[i] = (num - BigInt(nums[i-1]) + BigInt(2) * dp[i-1]) % mod;
    power = (power + num * num * dp[i]) % mod;
  }

  return Number(power);
};
```



## 复杂度分析

- 时间复杂度：$\mathcal{O}(n\log n)$，其中 $n$ 为 $\textit{nums}$ 的长度。
- 空间复杂度：$\mathcal{O}(n)$。

# 感想

虽然过程崎岖，但胜在收获不错，动态规划的题目，如果没有思路，在草稿纸上列出具体的例子，或许能帮助你发现数据的规律，并找到解题方法。

非常感谢你的阅读，衷心祝愿你不会踩到取模陷阱。

我还是免不了因为取模，喜提 WA *2：
![wa1.png](https://pic.leetcode.cn/1704512340-GOtWxC-wa1.png)
![wa2.png](https://pic.leetcode.cn/1704512346-YLKJot-wa2.png)


警钟长鸣，谨记取模陷阱!
