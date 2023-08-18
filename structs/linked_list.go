package structs

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewListNode(nums []int) *ListNode {
	head := &ListNode{}
	cur := head

	for _, num := range nums {
		cur.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		cur = cur.Next
	}

	return head.Next
}

// ToSlice 链表转数组
func (h *ListNode) ToSlice() []int {
	if h.HasCycle() {
		panic("list has cycle")
	}

	var nums = make([]int, 0, 100)
	cur := h
	for cur != nil {
		nums = append(nums, cur.Val)
		cur = cur.Next
	}
	return nums
}

// HasCycle
// 快慢指针
// 时间复杂度：O(n)
// 空间复杂度：O(1)
func (h *ListNode) HasCycle() bool {
	var slow, fast = h, h

	for fast != nil && fast.Next != nil {
		// 兔子追上乌龟（套圈），说明有环
		if fast == slow || fast.Next == slow {
			return true
		}

		slow = slow.Next      // 乌龟走一步
		fast = fast.Next.Next // 兔子走两步
	}

	// 访问到了链表末尾，无环
	return false
}

// Reverse 反转链表
func (h *ListNode) Reverse() *ListNode {
	var cur = h
	var pre *ListNode

	for cur != nil {
		next := cur.Next
		cur.Next = pre
		pre = cur
		cur = next
	}
	return pre
}

func NewCycleListNode(nums []int, pos int) (*ListNode, *ListNode) {
	var (
		head  = &ListNode{}
		cur   = head
		cycle *ListNode
	)
	for i, num := range nums {
		cur.Next = &ListNode{
			Val:  num,
			Next: nil,
		}
		cur = cur.Next

		if i == pos {
			cycle = cur
		}
	}
	cur.Next = cycle
	return head.Next, cycle
}
