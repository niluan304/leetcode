package main

func numSubarrayProductLessThanK(nums []int, k int) int {
	var (
		count = 0
	)

	n := len(nums)
	for i := range nums {
		sum := 1
		for j := i; j < n; j++ {
			sum *= nums[j]
			if sum >= k {
				break
			}
			count++
		}
	}

	return count
}

func numSubarrayProductLessThanK2(nums []int, k int) int {
	var (
		count = 0
		left  = 0
		right = 0
		sum   = 1
		n     = len(nums)
	)

	for i, num := range nums {
		sum *= num
		if sum < k {
			continue
		}
		count += Sum(i - left)
		for sum >= k && left <= i {
			sum /= nums[left]
			left++
		}

		count -= Sum(right - left)
		right = i
	}

	count += Sum(n - left)
	count -= Sum(right - left)

	return count
}

// Sum 求自然数数列[1,n]的和
func Sum(n int) int {
	if n <= 0 {
		return 0
	}
	return n * (n + 1) / 2
}

func numSubarrayProductLessThanK3(nums []int, k int) int {
	var (
		left = 0
		prod = 1
		ans  = 0
	)

	for right, num := range nums {
		prod *= num
		for prod >= k && left <= right {
			prod /= nums[left]
			left++
		}
		ans += right - left + 1
	}
	return ans
}
