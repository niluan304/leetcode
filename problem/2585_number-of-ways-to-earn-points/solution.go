package main

import (
	"fmt"

	. "github.com/niluan304/leetcode/copypasta/dp"
)

const MOD = 1_000_000_007

// 分组背包
func waysToReachTarget(target int, kinds [][]int) int {
	var n = len(kinds)
	var dp = make([][]int, n+1)
	for i, _ := range dp {
		dp[i] = make([]int, target+1)
	}
	dp[0][0] = 1

	for i, kind := range kinds {
		count, mark := kind[0], kind[1]
		for j := 0; j <= target; j++ {
			dp[i+1][j] = dp[i][j]

			for k := 1; k <= min(count, j/mark); k++ { // k == 0，等价于 dp[i+1][j] = dp[i][j]
				dp[i+1][j] += dp[i][j-k*mark]
			}

			dp[i+1][j] %= MOD
		}
	}
	return dp[n][target]
}

func waysToReachTarget2(target int, kinds [][]int) int {
	var stocks, weights []int
	for _, kind := range kinds {
		count, mark := kind[0], kind[1]
		stocks = append(stocks, count)
		weights = append(weights, mark)
	}
	return GroupKnapsackWaysToSum(stocks, weights, target) // need % MOD
}

func waysToReachTarget3(target int, types [][]int) int {
	// counts := make([]int, len(types))
	// for i := range types {
	// 	counts[i] = types[i][0]
	// }
	// values := make([]int, len(types))
	// for i := range types {
	// 	values[i] = types[i][1]
	// }

	// dp := BoundedKnapsackDPCountWays(values, counts)
	// return dp[target]
	K := NewBoundedKnapsackWays(target, 1e9+7)
	for _, t := range types {
		K.Add(t[1], t[0])
	}
	return K.Query(target)
}

// NewBoundedKnapsackWays
// maxWeight: 需要的价值上限.
// mod: 模数，传入-1表示不需要取模.
func NewBoundedKnapsackWays(maxValue int, mod int) *BoundedKnapsackWays {
	dp := make([]int, maxValue+1)
	dp[0] = 1
	return &BoundedKnapsackWays{
		dp:       dp,
		mod:      mod,
		maxValue: maxValue,
	}
}

// BoundedKnapsackWays 多重背包求方案数.
type BoundedKnapsackWays struct {
	dp       []int
	mod      int
	maxValue int
	maxJ     int
}

// Add
// 加入一个价值为value(value>0)的物品，数量为count.
// O(maxValue).
func (ks *BoundedKnapsackWays) Add(value, count int) {
	if value <= 0 {
		panic(fmt.Sprintf("value must be positive, but got %d", value))
	}
	ks.maxJ = min(ks.maxJ+count*value, ks.maxValue)
	if ks.mod == -1 {
		for j := value; j <= ks.maxJ; j++ {
			ks.dp[j] += ks.dp[j-value]
		}
		for j := ks.maxJ; j >= value*(count+1); j-- {
			ks.dp[j] -= ks.dp[j-value*(count+1)]
		}
	} else {
		for j := value; j <= ks.maxJ; j++ {
			ks.dp[j] = (ks.dp[j] + ks.dp[j-value]) % ks.mod
		}
		for j := ks.maxJ; j >= value*(count+1); j-- {
			ks.dp[j] = (ks.dp[j] - ks.dp[j-value*(count+1)]) % ks.mod
		}
	}
}

func (ks *BoundedKnapsackWays) Query(value int) int {
	if value < 0 || value > ks.maxValue {
		return 0
	}
	if ks.mod == -1 {
		return ks.dp[value]
	}
	if ks.dp[value] < 0 {
		ks.dp[value] += ks.mod
	}
	return ks.dp[value]
}

func (ks *BoundedKnapsackWays) Copy() *BoundedKnapsackWays {
	dp := append(ks.dp[:0:0], ks.dp...)
	return &BoundedKnapsackWays{
		dp:       dp,
		mod:      ks.mod,
		maxValue: ks.maxValue,
		maxJ:     ks.maxJ,
	}
}
