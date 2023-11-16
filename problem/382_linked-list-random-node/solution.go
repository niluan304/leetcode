package main

import (
	"math/rand"
	"time"

	. "github.com/EndlessCheng/codeforces-go/leetcode/testutil"
)

var seed = rand.New(rand.NewSource(time.Now().UnixNano()))

type Solution struct {
	head *ListNode
}

func Constructor(head *ListNode) Solution {
	return Solution{head: head}
}

func (s *Solution) GetRandom() (node int) {
	cur := s.head
	for i := 1; cur.Next != nil; i++ {
		if seed.Intn(i) == 1 {
			node = cur.Val
		}
		cur = cur.Next
	}
	return node
}

type Solution2 []int

func Constructor2(head *ListNode) Solution2 {
	var s Solution2
	for node := head; node != nil; node = node.Next {
		s = append(s, node.Val)
	}
	return s
}

func (s Solution2) GetRandom() int {
	return s[seed.Intn(len(s))]
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */
