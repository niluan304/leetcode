package main

func fib(n int) int {
	x, y := 0, 1

	for i := 0; i < n; i++ {
		x, y = y, x+y
	}

	return x
}

// 递归
func fib2(n int) int {
	if n < 2 {
		return n
	}
	return fib2(n-1) + fib2(n-2)
}

// 动态规划五部曲
func fib3(n int) int {
	if n < 2 {
		return n
	}

	// 1. dp[i]: 第i个斐波那契数
	var dp = make([]int, n+1)

	// 3. 初始化 dp数组
	dp[0] = 0
	dp[1] = 1

	// 4. 确定遍历顺序：从小到大
	for i := 2; i <= n; i++ {
		// 2. 确定递推公式
		dp[i] = dp[i-1] + dp[i-2]
	}

	// 5. debug: 打印 dp数组
	//fmt.Println("n", n, "dp", dp)
	return dp[n]
}

func fib4(n int) int {
	var answer = [31]int{
		0,
		1,
		1,
		2,
		3,
		5,
		8,
		13,
		21,
		34,
		55,
		89,
		144,
		233,
		377,
		610,
		987,
		1597,
		2584,
		4181,
		6765,
		10946,
		17711,
		28657,
		46368,
		75025,
		121393,
		196418,
		317811,
		514229,
		832040,
	}
	return answer[n]
}

func fib5(n int) int {
	switch n {
	case 2:
		return 1
	case 3:
		return 2
	case 4:
		return 3
	case 5:
		return 5
	case 6:
		return 8
	case 7:
		return 13
	case 8:
		return 21
	case 9:
		return 34
	case 10:
		return 55
	case 11:
		return 89
	case 12:
		return 144
	case 13:
		return 233
	case 14:
		return 377
	case 15:
		return 610
	case 16:
		return 987
	case 17:
		return 1597
	case 18:
		return 2584
	case 19:
		return 4181
	case 20:
		return 6765
	case 21:
		return 10946
	case 22:
		return 17711
	case 23:
		return 28657
	case 24:
		return 46368
	case 25:
		return 75025
	case 26:
		return 121393
	case 27:
		return 196418
	case 28:
		return 317811
	case 29:
		return 514229
	case 30:
		return 832040
	}

	return n
}
