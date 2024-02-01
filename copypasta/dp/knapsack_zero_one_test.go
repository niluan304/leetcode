package dp

import (
	"testing"

	. "github.com/niluan304/leetcode/container"
)

func Test_zeroOneKnapsack(t *testing.T) {
	type args struct {
		values  []int
		weights []int
		maxW    int
	}

	lc1049Case1 := []int{2, 7, 4, 1, 8, 1}
	lc1049Case2 := []int{31, 26, 33, 21, 40}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC 1049 case1", // 需转换
			args: args{
				values:  lc1049Case1,
				weights: lc1049Case1,
				maxW:    Sum(lc1049Case1) / 2,
			},
			want: (Sum(lc1049Case1) - 1) / 2,
		},
		{
			name: "LC 1049 case2", // 需转换
			args: args{
				values:  lc1049Case2,
				weights: lc1049Case2,
				maxW:    Sum(lc1049Case2) / 2,
			},
			want: (Sum(lc1049Case2) - 5) / 2,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroOneKnapsack(tt.args.values, tt.args.weights, tt.args.maxW); got != tt.want {
				t.Errorf("ZeroOneKnapsack() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroOneKnapsackAtLeastFillUp(t *testing.T) {
	type args struct {
		values  []int
		weights []int
		maxW    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC 2742 case1", // 需转换
			args: args{
				values:  []int{1, 2, 3, 2},
				weights: []int{1 + 1, 2 + 1, 3 + 1, 2 + 1},
				maxW:    4,
			},
			want: 3,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroOneKnapsackAtLeastFillUp(tt.args.values, tt.args.weights, tt.args.maxW); got != tt.want {
				t.Errorf("ZeroOneKnapsackAtLeastFillUp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroOneKnapsackExactlyFull(t *testing.T) {
	type args struct {
		values  []int
		weights []int
		maxW    int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC 2915 case1",
			args: args{
				values:  []int{1, 1, 1, 1, 1},
				weights: []int{1, 2, 3, 4, 5},
				maxW:    9,
			},
			want: 3,
		},
		{
			name: "LC 2915 case2",
			args: args{
				values:  []int{1, 1, 1, 1, 1, 1},
				weights: []int{4, 1, 3, 2, 1, 5},
				maxW:    7,
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroOneKnapsackExactlyFull(tt.args.values, tt.args.weights, tt.args.maxW); got != tt.want {
				t.Errorf("ZeroOneKnapsackExactlyFull() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_zeroOneWaysToSum(t *testing.T) {
	type args struct {
		nums []int
		sum  int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC 2787 case1",
			args: args{
				nums: []int{1, 4, 9, 16, 25, 36, 49, 64, 81, 100},
				sum:  10,
			},
			want: 1,
		},
		{
			name: "LC 2787 case2",
			args: args{
				nums: []int{1, 2, 3, 4},
				sum:  4,
			},
			want: 2,
		},
		{
			name: "LC 494 case1", // 需要转换
			args: args{
				nums: []int{1, 1, 1, 1, 1},
				sum:  (Sum([]int{1, 1, 1, 1, 1}) + 3) / 2,
			},
			want: 5,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ZeroOneWaysToSum(tt.args.nums, tt.args.sum); got != tt.want {
				t.Errorf("ZeroOneWaysToSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
