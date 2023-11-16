package main

func longestAlternatingSubarray(nums []int, threshold int) int {
	var ans = 0
	var n = len(nums)
	for i := 0; i < n; {
		j := i
		i++
		if nums[j]%2 != 0 || nums[j] > threshold {
			continue
		}
		for i < n && nums[i] <= threshold && nums[i]%2 != nums[i-1]%2 {
			i++
		}
		ans = max(ans, i-j)
	}
	return ans
}

func longestAlternatingSubarray2(nums []int, threshold int) int {
	var ans = 0

	for i, num := range nums {
		if num%2 != 0 || num > threshold {
			continue
		}

		var count = 1
		for j := i + 1; j < len(nums); j++ {
			if nums[j] > threshold || nums[j]%2 == nums[j-1]%2 {
				break
			}
			count++
		}
		ans = max(ans, count)
	}
	return ans
}

// Deprecated: Fail to pass
func longestAlternatingSubarray3(nums []int, threshold int) int {
	var ans = 0
	var left = 0
	var n = len(nums)

	for _, num := range nums {
		if num%2 == 0 && num <= threshold {
			ans = 1
			break
		}
	}

	for right := 1; right < n; right++ {
		for left <= right && nums[left]%2 != 0 {
			left++
		}
		if nums[right] > threshold {
			left = right + 1
		} else if nums[right]%2 != nums[right-1]%2 {
			ans = max(ans, right-left+1)
		}
	}
	return ans
}
