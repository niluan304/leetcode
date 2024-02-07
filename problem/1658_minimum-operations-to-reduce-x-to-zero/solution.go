package main

func minOperations(nums []int, x int) int {
	var (
		n   = len(nums)
		ans = n + 1
	)

	var dfs func(left, right, sum int)
	dfs = func(left, right, sum int) {
		if left == right {
			return
		}

		count := left + n - right
		xl := sum + nums[left]
		if xl < x {
			dfs(left+1, right, xl)
		} else if xl == x {
			ans = min(ans, count)
			return
		}

		xr := sum + nums[right]
		if xr < x {
			dfs(left, right-1, xr)
		} else if xr == x {
			ans = min(ans, count)
			return
		}
	}

	dfs(0, n-1, 0)

	if ans == n+1 {
		ans = -1
	}

	return ans
}

func minOperations2(nums []int, x int) int {
	n := len(nums)
	target := -x
	for _, num := range nums {
		target += num
	}

	if target < 0 {
		return -1
	}
	if target == 0 {
		return n
	}

	var (
		ans  = -1
		sum  = 0
		left = 0
	)
	for right, num := range nums {
		sum += num

		for sum > target {
			sum -= nums[left]
			left++
		}

		if sum == target {
			// 应当是 ans =max(ans, right-left +1), 但可以将 +1 改为  n - ans - 1
			ans = max(ans, right-left)
		}
	}

	if ans == -1 {
		return -1
	}

	return n - ans - 1
}
