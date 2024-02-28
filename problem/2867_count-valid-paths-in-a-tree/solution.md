树形DP + 分类讨论

## 树形DP
题面要求统计只包含一个质数的路径数量，可以尝试枚举每个节点，去统计经过这个节点的路径数量，那么该怎么做？





首先，需要分类讨论，

- 如果一个节点是质数，那么其他点就不能是质数；
- 如果一个节点是非质数，那么就还需要一个质数节点



### 根节点为非质数

先假设作为拐点的根节点是非质数的，我们可以找出根节点的所有子树，各有多少个「质数节点数」和「非质数节点数」路径，这个可以用数组来保存

```go
var pairs []struct{ prime, notPrime int }
```
此时根节点为非质数，那么还需要一个质数，可以通过排列组合，将子树的 「只含$1$质数路径数量」 *「 其他子树非质数路径数量」，就能得到非质数根节点时的合法路径数量。



以下面这颗树举例，：

![image-20240228003951202.png](https://pic.leetcode.cn/1709089790-ADbPxY-image-20240228003951202.png)

- 通过子树 $3$ 有 3条只含 1质数的路径，0条不含质数的路径
- 通过子树 $4$ 有 2条只含 1质数的路径，2条不含质数的路径
- 通过子树 $6$ 有 1条只含 1质数的路径，2条不含质数的路径

对应到这棵树：

- 从子树 $3$ 只含 $1$ 质数的路径出发，经过根节点，到子树 $4$, $6$ 和根节点：$3 * (2 + 2 +1)$
- 从子树 $4$ 只含 $1$ 质数的路径出发，经过根节点，到子树 $3$, $6$ 和根节点：$2 * (0 + 2 + 1)$
- 从子树 $6$ 只含 $1$ 质数的路径出发，经过根节点，到子树 $3$, $4$ 和根节点：$1 * (0 + 2 + 1)$

相加可得，新增的合法路径为 $24$ 条。

分析可知，这其实是一个组合问题，从子树A出发，经过非质数的根节点，到其他子树的合法路径为：子树 A 只含 $1$ 质数的路径数量 * （其他子树不含质数的路径数量之和 + 1）。

也就是 ：

```go
A.prime * sum(1 + B.notPrime + C.notPrime + D.notPrime + ...)
```

同理，从子树B出发，合法路径为：

```go
B.prime * sum(1 + A.notPrime + C.notPrime + D.notPrime + ...)
```



相应的代码逻辑：

```go
// var pairs []struct{ prime, notPrime int } 
// pairs 保存每个子树中有多少个「1 质数路径数」和「非质数路径数」

// 以非质数的根节点为拐点，新增的合法路径的数量
for _, pair := range pairs { // notPrime 初始化为子树的 notPrime 之和
    ans += pair.prime * (notPrime - pair.notPrime + 1) // +1 表示以根节点为终点的情况
}
```



### 根节点为质数

仍然以上面那颗树举例，如果根节点为质数，那么子树上的路径，必须全部为非质数。

为避免统计重复路径 $(a, b)$ 与 $(b, a)$，可以规定一个方向：只从左边出发，以右边为终点：

- 从子树 $3$ 不含质数的路径出发，经过根节点，到子树 $4$, $6$ 和根节点：$0 * (2 + 2 + 1)$
- 从子树 $4$ 不含质数的路径出发，经过根节点，到子树 $6$ 和根节点：$2 * (2 + 1)$
- 从子树 $6$ 不含质数的路径出发，只能到根节点了：$2 * (1)$

相加可得，新增的合法路径为 $8$ 条。



分析可知，从子树A出发，经过质数根节点，以右边子树为终点的合法路径有：子树A的非质数路径数量 *（右边子树的非质数路径数量之和 +1）。

也就是 ：

```go
A.notPrime * sum(1 + B.notPrime + C.notPrime + D.notPrime + ...)
```

同理，从子树B出发，合法路径为：

```go
B.notPrime * sum(1 + C.notPrime + D.notPrime + ...)
```



相应的代码逻辑：

```go
// var pairs []struct{ prime, notPrime int } 
// pairs 保存每个子树中有多少个「1 质数路径数」和「非质数路径数」

// 以质数的根节点为拐点，新增的合法路径的数量
for _, pair := range pairs {
    notPrime -= pair.notPrime // notPrime 初始化为子树的 notPrime 之和
    ans += pair.notPrime * (notPrime + 1) // +1 表示以根节点为终点的情况
}
```



### 确定 dfs 返回值

经过上面的分类讨论后，我们已经可以求出，每个节点作为拐点时的合法路径数量，但还有一个问题需要解决，那就是确定子树的「$1$ 质数路径数」和「非质数路径数」。



可以先假设子树A的根节点为**质数**：

那么通过节点A的「非质数路径数」，很明显为 $0$；

求通过节点A的「$1$ 质数路径数」，由于节点A 是质数，那么后续有效的路径只能是非质数的，也就是子树的「非质数路径」之和以及自身。



如果子树A的根节点为**非质数**：

那么通过节点A的「$1$ 质数路径数」，则只能从 A的子树中寻找，也就是 A的子树的「$1$ 质数路径数」之和；

求通过节点A的「非质数路径数」，那么后续有效的路径只能是非质数的，也就是子树的「非质数路径」之和以及自身。

### 代码实现

```go []
const mx = 1e5

// NotPrimes 质数 = false，非质数 = true
var NotPrimes = notPrimes(mx) // 预处理质数

func notPrimes(n int) (notPrimes []bool) {
	notPrimes = make([]bool, max(n+1, 2))
	for i := 2; i <= n; i++ {
		if !notPrimes[i] {
			for j := i * i; j <= n; j += i {
				notPrimes[j] = true
			}
		}
	}
	notPrimes[0], notPrimes[1] = true, true // 0 和 1 不是质数，也不是合数
	return notPrimes
}

// 树形 dp + 分类讨论
func countPaths(n int, edges [][]int) int64 {
	graph := make([][]int, n+1)
	for _, edge := range edges {
		x, y := edge[0], edge[1]
		graph[x] = append(graph[x], y)
		graph[y] = append(graph[y], x)
	}

	type Pair struct{ prime, notPrime int }

	ans := 0

	var dfs func(i int, fa int) (prime int, notPrime int) // dfs(i, fa) 表示当前遍历到第 i 个节点，父节点为 fa
	dfs = func(i int, fa int) (prime int, notPrime int) {
		pairs := make([]Pair, 0, len(graph[i]))
		for _, j := range graph[i] { // 每个节点只会访问一次
			if j == fa {
				continue
			}
			p, np := dfs(j, i)
			pairs = append(pairs, Pair{prime: p, notPrime: np})
			prime += p
			notPrime += np
		}

		// 分类讨论
		if NotPrimes[i] {
			// 以非质数的根节点为拐点，新增的合法路径的数量
			for _, pair := range pairs {
				ans += pair.prime * (notPrime - pair.notPrime + 1)
			}
			return prime, notPrime + 1
		} else {
			// 以质数的根节点为拐点，新增的合法路径的数量
			np := notPrime
			for _, pair := range pairs {
				np -= pair.notPrime // 避免重复统计 a->b, b->a
				ans += pair.notPrime * (np + 1)
			}
			return notPrime + 1, 0 // 根节点为质数时的「1 质数路径数」 = 自身 + 子树非质数路径数之和
		}
	}
	dfs(1, -1)
	return int64(ans)
}

```

#### 复杂度分析

- 时间复杂度：$\mathcal{O}(n)$。不计入预处理的时间。
- 空间复杂度：$\mathcal{O}(n)$。