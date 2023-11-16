package main

import (
	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

// 方法二：快慢指针
// 思路与算法
//
// 我们使用两个指针，fast 与 slow。它们起始都位于链表的头部。随后，slow 指针每次向后移动一个位置，而 fast 指针向后移动两个位置。
// 如果链表中存在环，则 fast 指针最终将再次与 slow 指针在环中相遇。
//
// 如下图所示，设链表中环外部分的长度为 a。slow 指针进入环后，又走了 b 的距离与 fast 相遇。
// 此时，fast 指针已经走完了环的 n 圈，因此它走过的总距离为 a+n(b+c)+b=a+(n+1)b+nc。
// ![](https://assets.leetcode-cn.com/solution-static/142/142_fig1.png)
// 根据题意，任意时刻，fast 指针走过的距离都为 slow 指针的 2 倍。因此，我们有
//
// a+(n+1)b+nc=2(a+b) => a=c+(n−1)(b+c)
// 有了 a=c+(n−1)(b+c) 的等量关系，我们会发现：从相遇点到入环点的距离加上 n−1 圈的环长，恰好等于从链表头部到入环点的距离。
//
// 因此，当发现 slow与 fast 相遇时，我们再额外使用一个指针 ptr。
// 起始，它指向链表头部；随后，它和 slow 每次向后移动一个位置。最终，它们会在入环点相遇。
//
// 作者：力扣官方题解
// 链接：https://leetcode.cn/problems/linked-list-cycle-ii/solutions/441131/huan-xing-lian-biao-ii-by-leetcode-solution/
func leetcode2(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil {
		slow = slow.Next
		if fast.Next == nil {
			return nil
		}
		fast = fast.Next.Next
		if fast == slow {
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
