package main

func waysToBuyPensPencils(total int, cost1 int, cost2 int) int64 {
	var ans int64 = 0

	// 能买 [0, total/cost1] 支钢笔
	for i := 0; i <= total/cost1; i++ {
		// 买 i支钢笔前提下，最多能买多少支铅笔
		n := (total - cost1*i) / cost2
		ans += int64(n + 1) // +1表示买0支铅笔的情况
	}

	return ans
}

// error
// 整数相除导致无法合并通项式
func waysToBuyPensPencils2(total int, cost1 int, cost2 int) int64 {
	if cost1 < cost2 {
		cost1, cost2 = cost2, cost1
	}
	// 能买 [0, total/cost1] 支钢笔
	n := int64(total/cost1) + 1
	sum := int64(total)*n - int64(cost1)*(n-1)*n/2
	//sum = n * (int64(total) - int64(cost1)*(n-1)/2)

	return sum/int64(cost2) + n
}
