### 思考题：如果操作改为可以加一或减一呢？

原题只能增加左右子树中的小值，修改后多了一种操作：减少左右子树中的大值。

记叶子节点的根节点为 $A$，假设 $A$ 的左右子树，也就是叶子节点的值分别为：$a_0$ 和 $a_1$，且有 $a_0 \geq a_1$，让这两个叶子节点相等的最小代价为：$a_0 - a_1$，此时这两个叶子节点可以在区间 $[a_0, a_1]$ 内取值。



#### 平衡左右子树

再站在 $A$ 的父节点 $P$ 的角度看，$P$ 到 $A$ 的叶子节点的路径值，则是在区间 $[a_0+A, a_1+A]$  内取值，记为 $[x_0, x_1]$。

同理，$P$ 到$A$ 的兄弟节点$B$ 的叶子节点的路径值，则可以在区间  $[y_0, y_1]$ 内，

和原题一样的思路，该如何花费最小的代价，来 「平衡」左右子树呢？

分析可知，区间 $[x_0, x_1]$ 和  $[y_0, y_1]$ 有相交和不相交两种情况：

- 不相交，必须调整左右子树，和分析节点 $A$ 时一致，此时的最小代价为：$y_0 - x_1$ ，假设  $ y_1 \geq x_1$，此时取值范围会变为：$[x_1, y_0]$

- 相交时，最小代价为 $0$，此时需要将取值范围限制到相交的区间：$[\max(x_0, y_0), \min(x_1, y_1)]$




#### 确定递归的返回值
新的左右子树取值范围记为：$[z_0, z_1]$，假设 $P$ 的根节点为 $PP$， $PP$ 到 叶子节点的路径值则在区间 $[z_0+P, z_1+P]$ 内取值。



#### 代码实现

```go []
// 思考题：如果操作改为可以加一或减一呢？
func minIncrements(n int, cost []int) int {
	ans := 0
	var dfs func(i int) (low, high int)
	dfs = func(i int) (low, high int) {
		if i > n {
			return 0, 0
		}

		leftLow, lowHigh := dfs(2 * i)      // 左儿子的取值范围
		rightLow, rightHigh := dfs(2*i + 1) // 右儿子的取值范围

		z0, z1, c := balance(leftLow, lowHigh, rightLow, rightHigh) // 让左右儿子平衡的取值范围，以及操作代价
		ans += c
		return cost[i-1] + z0, cost[i-1] + z1 // 当前节点的取值范围
	}
	dfs(1)
	return ans
}

func balance(x0, x1 int, y0, y1 int) (z0, z1 int, cost int) {
	// 不相交的两种情况
	if y0 > x1 {
		return x1, y0, y0 - x1
	}
	if x0 > y1 {
		return y1, x0, x0 - y1
	}

	// 相交的情况
	return max(x0, y0), min(x1, y1), 0
}

```

