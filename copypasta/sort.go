package copypasta

import (
	"cmp"
	"slices"
)

// MergeSortWith
// 根据 cmp 函数作归并排序
//
// 归并排序基于分治思想将数组分段排序后合并，也是一种稳定排序的算法
//
// 闭包可用于逆序对的拓展题目。
// 归并填入后才会调用闭包，后续不会再使用入参切片，因此闭包内可以修改入参切片
//
// 带索引排序：
// LC315. https://leetcode.cn/problems/count-of-smaller-numbers-after-self 带索引的排序
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
