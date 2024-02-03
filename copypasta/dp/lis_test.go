package dp

import (
	"reflect"
	"testing"
)

func TestLIS(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "LC300",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: 4,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LIS(tt.args.nums); got != tt.want {
				t.Errorf("LIS() = %v, want %v", got, tt.want)
			}
		})

		t.Run(tt.name, func(t *testing.T) {
			if got := LISGreedy(tt.args.nums); got != tt.want {
				t.Errorf("LISGreedy() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestLISPath(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "LC300",
			args: args{
				nums: []int{10, 9, 2, 5, 3, 7, 101, 18},
			},
			want: []int{2, 3, 7, 101},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := LISPath(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("LISPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
