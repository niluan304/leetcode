package main

// 方法二：从皇后出发
// 思路与算法
//
// 我们枚举每个皇后，判断它是否在国王的八个方向上。如果在，说明皇后可以攻击到国王。
//
// 记国王的位置为 (kx,ky) ，皇后的位置为 (qx,qy)，那么皇后相对于国王的位置为 (x,y)=(qx−kx,qy−ky)，
// 显然当 x=0 或 y=0 或 ∣x∣=∣y∣ 时，皇后可以攻击到国王，方向为 (sgn(x),sgn(y))，
// 其中 sgn(x) 为符号函数，当 x>0 时为 1，x<0 时为 −1，x=0 时为 0。
//
// 同一个方向的皇后可能有多个，我们需要选择距离国王最近的那一个，因此可以使用一个哈希映射，
// 它的键表示某一个方向，值是一个二元组，分别表示当前距离最近的皇后以及对应的距离。
// 当我们枚举到一个新的皇后时，如果它在国王的八个方向上，就与哈希映射中对应的值比较一下大小关系即可。
//
// 当枚举完所有皇后，我们就可以从哈希映射值的部分中得到答案。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/queens-that-can-attack-the-king/
// 来源：力扣（LeetCode）
// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。
func leetcode2(queens [][]int, king []int) [][]int {
	var sgn = func(x int) int {
		if x > 0 {
			return 1
		} else if x == 0 {
			return 0
		} else {
			return -1
		}
	}

	var abs = func(x int) int {
		if x < 0 {
			return -x
		}
		return x
	}

	candidates := make(map[int][]int)
	kx, ky := king[0], king[1]
	for _, queen := range queens {
		qx, qy := queen[0], queen[1]
		x, y := qx-kx, qy-ky
		if x == 0 || y == 0 || abs(x) == abs(y) {
			dx, dy := sgn(x), sgn(y)
			key := dx*10 + dy
			if val, ok := candidates[key]; !ok || val[2] > abs(x)+abs(y) {
				candidates[key] = []int{qx, qy, abs(x) + abs(y)}
			}
		}
	}

	ans := [][]int{}
	for _, value := range candidates {
		ans = append(ans, []int{value[0], value[1]})
	}
	return ans
}
