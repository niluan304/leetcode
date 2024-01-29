package main

func maxNonDecreasingLength(nums1 []int, nums2 []int) int {

	var n = len(nums1)
	var ans = 1 // math.MaxInt32 // math.MinInt32

	var dp1, dp2 = 1, 1
	for i := 1; i < n; i++ {
		x1, x2 := 1, 1
		if nums1[i] >= nums1[i-1] {
			x1 = max(x1, dp1+1)
		}
		if nums1[i] >= nums2[i-1] {
			x1 = max(x1, dp2+1)
		}

		if nums2[i] >= nums1[i-1] {
			x2 = max(x2, dp1+1)
		}
		if nums2[i] >= nums2[i-1] {
			x2 = max(x2, dp2+1)
		}

		dp1, dp2 = x1, x2
		ans = max(ans, dp1, dp2)
	}

	return ans
}
