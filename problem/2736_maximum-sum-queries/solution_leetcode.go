package main

import (
	"math/rand"
	"slices"
	"sort"
)

// 想了三天，还是没想到应该怎么去设计这个问题的解决过程，还是看官解了。
// 这道题编码不难，难的是如何去证明这样编码的正确性，我们一步一步来，不急 ~
// 问题有两个维度，每个维度的值分别列于位于 nums1 和 nums2 中。看上面的暴力
// 解法需要逐一枚举，那么如果我们不逐一枚举的话，我们在遍历到 queries 后面
// 的值时，最好是能利用到遍历前面时的一些信息，使得整体的复杂度降到 O(k * (< n))。
// 这里的< n 指的是远小于线性的复杂度，例如 logn, logk 等等。nums1 和 nums2
// 里的数是无序的，也就是我们肯定至少要两个数组都完整地遍历一遍，那么我们
// 如果能做到每个数尽可能只遍历一次 (或者常数次)，这就是一种优化。我们先不
// 考虑怎样的数据结构可以实现这样的过程 (其实先想想也可以，来来去去就是单调栈、
// 单调队列、双指针滑动窗口这种)，我们先想想有没有“吞并”的场景，也就是说
// 满足某些条件下，可以不考虑 nums1 和 nums2 中的某一些数? 首先会有这样一个结论:
// 对于索引 i 和 j (i,j大小任意)，如果 nums1[i] >= nums1[j] 且 nums2[i] >= nums2[j]
// 那么 nums1 和 nums2 中索引为 j 的数可以被忽略，因为对于 query = [x, y]，
// 如果满足 nums1[j] >= x 且 nums2[j] >= y，nums1[i] 和 nums2[i] 一定也会
// 满足且它俩可以组成更大的和，nums1[j] 和 nums2[j] 这一对不可能出现在最终
// answer 中。

// 这个结论能否利用呢? 如果 nums1 和 nums2 都是无序的，那需要双重循环遍历才
// 可以得到这个信息，还不如直接暴力求解。但如果先按其中一个维度排序，那就有办法
// 利用了。我们不妨先对 nums1 升序 (注意按题目意思，我们需要维护原始顺序中每个
// nums1 和 nums2 对应索引数据的关联，因此我们需要先创建一个新的列表，每一项
// 由 nums1 和 nums2 原本对应的数对组成，然后对新的列表按 nums1 维度升序)，
// 这样子我们在遍历排序后的列表时，对于索引 i ，所有满足 j < i 且 nums2[j] <=
// nums2[i] 的索引 j 都可以去除。至于如何实现 “去除” ，我们可以用一个单调栈来
// 维护，这样子单调栈中存放的就是 nums1 维度单调递增且 nums2 维度单调递减的序列。

// 我们现在可以通过这样的单调栈求答案了吗? 我们看下单调栈序列中，尽管 nums1 和
// nums2 各自维度独立有序，加和其实还是无序的。我们能否人为让加和也变得有序? 我们
// 回看一下单调栈的构建过程，我们在让新的索引加入栈之前，栈顶索引 j 满足，nums2[j]
// > nums2[i]，我们排除掉栈中 nums2[j] <= nums2[i] 是因为这些 j 会被索引 i “吞并”，
// 但是当我们排除完栈中的索引后，我们如果能确保我们的 i 真的有优势，也就是满足
// nums1[i] + nums2[i] >= nums1[j] + nums2[j]，那么单调栈中加和就有一个单调递增
// 的趋势，这个时候我们就可以用二分法加快检索速度。但是这样真的会没有问题吗?

// 我们假设当前栈顶 nums1[j] + nums2[j] > nums1[i] + nums2[i]，那么按我们的约定
// 索引 i 是要被剔除的，假设索引 i = n-1 ，而我们有一个 query 满足 nums1[j] < x <=
// nums1[i] 且 nums2[i] >= y，此时我们在单调栈中找不到答案，但是原数组中是有一个
// 满足条件的数对的，只是被我们删除了。为什么会有这样的现象呢? 因为我们并没有针对每个
// query 去在符合条件的数对中做我们的单调栈构建操作。对于一个特定 query = [x, y]，
// 我们应该只将满足 nums1[i] >= x 的这部分索引加入单调栈，这样子可以确保 nums1[i] < x
// 的部分不会影响到当前 query 的正确性，因为这部分 nums1[i] 有可能会导致多删除一些
// 在当前 query 时刻有用的数对，但是这部分 nums1[i] 对当前 query 又是毫无贡献的
// (可以再多理解下)。

// 因此我们需要对 queries 也按照第一个维度升序，并且要保留每个 query 原始的索引，
// 因为最终 answer 需要维持原来 queries 的顺序返回。然后我们倒着遍历 queries 和
// nums1，针对每个 query ，找到大于等于 x 的最小索引 i，然后将 [i, n) 的部分按
// 上述规则构建单调栈。构建好后，单调栈现在满足 nums1 维度降序，nums2 维度升序，
// 加和降序，因此我们对 nums2 维度进行二分搜索，找到 nums2 维度大于等于 y 的最小值，
// 其对应索引的加和就是当前 query 的结果；如果对应索引越界，当前 query 无解。

// 说句题外话，我觉得这个题好像还挺有应用价值的，比如说 nums1 和 nums2 对应学校里
// 每个班级的男女生数量，最后要求男生和女生数量分别大于 queries[i][0] 和 queries[i][1]
// 的班级人数的最大值，到时就可以用这一招了哈哈哈!!!

// 时间复杂度 O((n + k) * logn)，对原数组排序 O(nlogn)，构建单调栈 O(n)，单调栈上二分
// O(klogn)。空间复杂度 O(n + k)，新建的 queries 以及单调栈使用
// (排序的 logn 栈空间在渐进意义下可忽略)。
func leetcodeMinTime(nums1, nums2 []int, queries [][]int) []int {
	n, k := len(nums1), len(queries)

	Sort(
		nums1,
		// 按 nums1 升序排列
		func(a, b int) int { return a - b },
		// nums2 跟随 nums1 排序，这样子可以不用新建列表，节省空间
		func(i, j int) { nums2[i], nums2[j] = nums2[j], nums2[i] },
	)

	// 构建新的 queries，补充原始索引信息
	sortedQueries := make([][]int, k)
	for i := range sortedQueries {
		sortedQueries[i] = []int{i, queries[i][0], queries[i][1]}
	}

	// 对 queries 也按照 nums1 维度升序
	Sort(
		sortedQueries,
		func(a, b []int) int { return a[1] - b[1] },
		func(i, j int) {},
	)

	// 按 nums1[i] + nums2[i] 单调递减， nums2[i] 单调递增的单调栈，其中存的是索引
	stk := make([]int, n)
	stkLen := 0

	i := n - 1
	answer := make([]int, k)

	// 倒着遍历 queries
	for q := k - 1; q >= 0; q-- {
		index, x, y := sortedQueries[q][0], sortedQueries[q][1], sortedQueries[q][2]

		// 将所有 nums1 大于等于 x 的数对入栈，确保 nums1 小于 x 的数对不会对当前
		// query 产生影响
		for ; i >= 0 && nums1[i] >= x; i-- {
			for stkLen > 0 {
				j := stk[stkLen-1]

				// 确保栈顶的加和大于准备加入的加和，剔除栈顶不满足条件的数对
				if nums1[j]+nums2[j] > nums1[i]+nums2[i] {
					break
				}

				stkLen--
			}

			// 当前数对还需要满足 nums2 维度大于栈顶才可以加入。
			// 这里我们考虑下，因为 num1 维度在单调栈中是降序的，如果 nums2 维度不大于栈顶，
			// 上面也不会出现弹出栈顶的情况
			if stkLen == 0 || nums2[i] > nums2[stk[stkLen-1]] {
				stk[stkLen] = i
				stkLen++
			}
		}

		// 满足当前 query 条件的单调栈构建完成，可以开始在单调栈上二分搜索了
		left, right := 0, stkLen
		for left < right {
			middle := left + (right-left)/2
			if nums2[stk[middle]] >= y {
				right = middle
			} else {
				left = middle + 1
			}
		}

		if left < stkLen {
			j := stk[left]
			answer[index] = nums1[j] + nums2[j]
		} else {
			answer[index] = -1
		}
	}

	return answer
}

// 复习写快速排序

type Compare[T any] func(a, b T) int

func Sort[T any](arr []T, compare Compare[T], onSwap func(a, b int)) {
	innerSort(arr, 0, len(arr)-1, compare, onSwap)
}

func innerSort[T any](
	arr []T,
	left, right int,
	compare Compare[T],
	onSwap func(a, b int),
) {
	if left >= right {
		return
	}

	p := partition(arr, left, right, compare, onSwap)
	innerSort(arr, left, p-1, compare, onSwap)
	innerSort(arr, p+1, right, compare, onSwap)
}

func partition[T any](
	arr []T, left, right int,
	compare Compare[T],
	onSwap func(a, b int),
) int {
	pivot := left + rand.Intn(right-left+1)
	swap(arr, pivot, right, onSwap)

	i, j := left, right-1
	for {
		for i <= j && compare(arr[i], arr[right]) < 0 {
			i++
		}

		for i <= j && compare(arr[j], arr[right]) > 0 {
			j--
		}

		if i > j {
			break
		}

		swap(arr, i, j, onSwap)
		i++
		j--
	}

	swap(arr, i, right, onSwap)
	return i
}

func swap[T any](arr []T, i, j int, onSwap func(a, b int)) {
	arr[i], arr[j] = arr[j], arr[i]
	onSwap(i, j)
}

func endlessCheng(nums1, nums2 []int, queries [][]int) []int {
	type Pair struct{ x, y int }
	pairs := make([]Pair, len(nums1))
	for i, x := range nums1 {
		pairs[i] = Pair{x, nums2[i]}
	}
	slices.SortFunc(pairs, func(a, b Pair) int { return b.x - a.x })
	qid := make([]int, len(queries))
	for i := range qid {
		qid[i] = i
	}

	for i := range queries {
		queries[i] = append(queries[i], i)
	}
	slices.SortFunc(queries, func(a, b []int) int {
		return b[0] - a[0] // 降序排列
	})

	slices.SortFunc(qid, func(i, j int) int { return queries[j][0] - queries[i][0] })

	ans := make([]int, len(queries))
	type Data struct{ y, sum int }
	stack := []Data{}
	j := 0

	for _, query := range queries {
		x, y, i := query[0], query[1], query[2]
		for ; j < len(pairs); j++ { // 下面只需关心 pairs[j].y
			px, py := pairs[j].x, pairs[j].y
			if px < x {
				break
			}

			sum := px + py
			m := len(stack)
			for ; m > 0; m-- { // pairs[j].y >= stack[len(stack)-1].y
				if stack[m-1].sum > sum {
					break
				}
				stack = stack[:m-1]
			}

			if m == 0 || stack[m-1].y < py {
				stack = append(stack, Data{y: py, sum: sum})
			}
		}

		k := sort.Search(len(stack), func(i int) bool {
			return stack[i].y >= y
		})
		if k < len(stack) {
			ans[i] = stack[k].sum
		} else {
			ans[i] = -1
		}
	}
	return ans
}
