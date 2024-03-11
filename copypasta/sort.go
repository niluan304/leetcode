package copypasta

import (
	"cmp"
	"slices"
)

// BinarySearch
// 在区间 [left, right) 中二分查找最小的 left, 满足 check(left) == true，
// 如果未能找到满足条件的 left, 那么将会返回 right 而不是 -1。
//
// 定义：check(left-1) == false and check(right) == true.
func BinarySearch[Number interface {
	~int | ~int64 | ~uint | ~uint64 | ~float64
}](left, right Number, check func(i Number) bool) Number {
	// left + (right-left)/2 不支持负数
	// 因此将 [left, right) 偏移至 [0, right-left)
	offset := left
	left, right = left-left, right-left // TODO: math.MaxInt - (-1)会导致溢出

	// 循环不变量： check(left-1) == false, check(right) == true.
	for left < right {
		mid := left + (right-left)/2 // avoid overflow when computing mid
		// left ≤ mid < right
		if !check(mid + offset) {
			left = mid + 1 // preserves check(left-1) == false
		} else {
			right = mid // preserves check(right) == true
		}
	}
	// left == right, check(left-1) == false, and check(right) (= check(left)) == true  =>  answer is left.
	return left + offset
}

// MergeSortWith
// 根据 cmp 函数作归并排序
//
// 归并排序基于分治思想将数组分段排序后合并，也是一种稳定排序的算法
//
// 闭包可用于逆序对的拓展题目。
// 归并填入后才会调用闭包，后续不会再使用入参切片，因此闭包内可以修改入参切片
//
// 带索引排序：
// LC315. https://leetcode.cn/problems/count-of-smaller-numbers-after-self
func MergeSortWith[T any](s []T, cmp func(x, y T) int, f func(left, right []T)) {
	n := len(s)
	if n <= 1 {
		return
	}

	mid := n / 2
	left, right := slices.Clone(s[:mid]), slices.Clone(s[mid:]) // 拷贝后，闭包内也可以直接修改入参

	// 递归完毕后，left 和 right 均为有序
	MergeSortWith(left, cmp, f)
	MergeSortWith(right, cmp, f)

	// left 和 right 归并填入 s
	i, j := 0, 0
	for cur := range s {
		// left[i] <= right[j] => cmp(left[i], right[j]) <= 0
		if i < len(left) && (j == len(right) || cmp(left[i], right[j]) <= 0) {
			s[cur] = left[i]
			i++
		} else {
			s[cur] = right[j]
			j++
		}
	}

	// 对 left, right 统计
	f(left, right)
}

// MergeSort
// 归并排序
//
// 归并排序基于分治思想将数组分段排序后合并，也是一种稳定排序的算法
//
// 闭包可用于统计逆序对。
// 归并填入后才会调用闭包，后续不会再使用入参切片，因此闭包内可以修改入参切片
//
// 模板题：
// LCR170. https://leetcode.cn/problems/shu-zu-zhong-de-ni-xu-dui-lcof/
// LC493. https://leetcode.cn/problems/reverse-pairs/
//
// 简单变形：
// LC2426. https://leetcode.cn/problems/number-of-pairs-satisfying-inequality/
//
// 转换题：区间和 -> 前缀和之差 -> 逆序对：
// LC327. https://leetcode.cn/problems/count-of-range-sum/
func MergeSort[T cmp.Ordered](s []T, f func(left, right []T)) {
	n := len(s)
	if n <= 1 {
		return
	}

	mid := n / 2
	left, right := slices.Clone(s[:mid]), slices.Clone(s[mid:]) // 拷贝后，闭包内也可以直接修改入参

	// 递归完毕后，left 和 right 均为有序
	MergeSort(left, f)
	MergeSort(right, f)

	// left 和 right 归并填入 s
	i, j := 0, 0
	for cur := range s {
		if i < len(left) && (j == len(right) || left[i] <= right[j]) {
			s[cur] = left[i]
			i++
		} else {
			s[cur] = right[j]
			j++
		}
	}

	// 对 left, right 统计
	f(left, right)
}
