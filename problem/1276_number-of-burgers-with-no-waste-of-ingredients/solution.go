package main

// # 思路
// 【鸡兔同笼原题】 **砍腿法**
// 鸡兔同笼，共15个头、40只脚，问鸡兔各几只？
// 【答】假设每个动物砍去两只脚，共砍去15×2=30只脚，这时鸡已经没脚了，剩下的 40-30=10只脚都是兔子砍剩下的脚，10÷2=5(兔)，15-5=10(鸡)。
//
// 【鸡兔同笼】的升级版，本题着重考察的是边界问题。
// **不成立**条件有三种情况。
// 1. tomatoSlices(番茄片)的数量不为偶数是不成立。
// 2. tomatoSlices(番茄片)的数量太多，全部用于制作**巨无霸汉堡**也用不完。
// 3. tomatoSlices(番茄片)的数量太少，全部用于制作**小皇堡**也不够。
//
// **成立**条件有两种情况;
// 1. 番茄片和奶酪片的数量时为0；
// 2. 番茄片和奶酪片的数量制作汉堡是刚好用完。
//
// # 复杂度
//
// 时间复杂度: $O(1)$
//
// 空间复杂度: $O(1)$
func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	x := tomatoSlices - 2*cheeseSlices
	if x < 0 || x%2 != 0 || (x/2 > cheeseSlices) {
		return nil
	}

	x /= 2
	return []int{x, cheeseSlices - x}
}

// 鸡兔同笼原题
func numOfBurgers2(tomatoSlices int, cheeseSlices int) []int {
	x := tomatoSlices - 2*cheeseSlices
	a, b := x/2, cheeseSlices-x/2
	if x%2 != 0 || a < 0 || b < 0 {
		return nil
	}
	return []int{a, b}
}
