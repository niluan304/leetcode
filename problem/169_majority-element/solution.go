package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
)

func majorityElement(nums []int) int {
	var count = map[int]int{}
	for _, num := range nums {
		count[num]++
	}

	half := len(nums) / 2
	for num, n := range count {
		if n > half {
			return num
		}
	}

	return 0
}

func leetcode2(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

func leetcode3(nums []int) int {
	major, cnt := nums[0], 1

	for i := 1; i < len(nums); i++ {
		if cnt == 0 {
			major = nums[i]
		}

		if major == nums[i] {
			cnt += 1
		} else {
			cnt -= 1
		}
	}
	return major
}

func initLeetcode4() {
	in := bufio.NewScanner(os.Stdin)
	in.Buffer(nil, math.MaxInt32)
	_out, _ := os.Create("user.out")
	out := bufio.NewWriter(_out)
	for in.Scan() {
		s := in.Bytes()
		ans, c := 0, 0
		for _i := 1; _i < len(s); _i++ {
			_neg := false
			if s[_i] == '-' {
				_i++
				_neg = true
			}
			v := int(s[_i] & 15)
			for _i++; s[_i]&15 < 10; _i++ {
				v = v*10 + int(s[_i]&15)
			}
			if _neg {
				v = -v
			}
			if c == 0 {
				ans = v
			}
			if ans == v {
				c++
			} else {
				c--
			}
		}
		fmt.Fprintln(out, ans)
	}
	out.Flush()
	os.Exit(0)
}

func leetcode4([]int) (_ int) { return }

func leetcode5(nums []int) int {
	bitc := make([]int, 32)
	for i := 0; i < 32; i++ {
		for _, num := range nums {
			if (num>>i)&1 == 1 {
				bitc[i]++
			} else {
				bitc[i]--
			}
		}
	}
	var ans int32
	for i := 0; i < 32; i++ {
		if bitc[i] >= 0 {
			ans += int32(1) << i
		}
	}
	return int(ans)
}

func leetcode6(nums []int) int {
	return majorityElementRec(nums, 0, len(nums)-1)
}

func countInRange(nums []int, num, lo, hi int) (count int) {
	for i := lo; i < hi; i++ {
		if nums[i] == num {
			count++
		}
	}

	return count
}

func majorityElementRec(nums []int, lo, hi int) int {
	if lo == hi {
		return nums[lo]
	}

	mid := (hi-lo)/2 + lo
	left := majorityElementRec(nums, lo, mid)
	right := majorityElementRec(nums, mid+1, hi)

	if left == right {
		return left
	}

	leftCount := countInRange(nums, left, lo, hi)
	rightCount := countInRange(nums, right, lo, hi)

	if leftCount > rightCount {
		return left
	} else {
		return right
	}
}
