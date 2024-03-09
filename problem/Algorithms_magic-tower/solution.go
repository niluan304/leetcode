package main

import (
	. "github.com/niluan304/leetcode/container"
)

// 反悔贪心
// 遍历数组，能加血就尽管加血，要扣血就直接扣血，但如果血量小于 1，我们就「反悔」：
// 从前面的扣血中，拿出一个扣血量最大的数（最小的负数），移到数组的末尾，把之前扣掉的血重新加回来。
//
// - 时间复杂度：$\mathcal{O}(n\log n)$。
// - 空间复杂度：$\mathcal{O}(n)$。
//
// 相似题目：
// LC630 https://leetcode.cn/problems/course-schedule-iii
func magicTower(nums []int) int {
	ans := 0

	health, damage := 1, 0
	hp := NewMinIntHeap()

	for _, num := range nums {
		if num < 0 {
			hp.Push(num)
		}
		health += num
		if health <= 0 {
			// 血量不是正整数了，需要挪动怪物房间，也意味着 num <= 0，堆不为空
			// 这里不需要 for health <= 0 {}，因为堆顶的元素必定小于等于 num
			ans++         // 记录挪动次数
			x := hp.Pop() // 出队
			damage += x   // 记录挪动的伤害
			health -= x   // 恢复被伤害扣除的血量
		}
	}

	// 如果挪动的总伤害大于等于血量，那么血量无法维持为正整数
	if health+damage <= 0 {
		return -1
	}
	return ans
}
