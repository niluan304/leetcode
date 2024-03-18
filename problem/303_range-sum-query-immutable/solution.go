package main

type NumArray struct {
	prefixSum []int
}

func Constructor(nums []int) NumArray {
	preSum := make([]int, len(nums)+1)
	for i, num := range nums {
		preSum[i+1] = preSum[i] + num
	}

	return NumArray{prefixSum: preSum}
}

func (n *NumArray) SumRange(left int, right int) int {
	return n.prefixSum[right+1] - n.prefixSum[left]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(left,right);
 */
