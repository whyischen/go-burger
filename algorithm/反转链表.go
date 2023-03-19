package algorithm

// 反转链表
// 1. 维护 cur、pre 两个指针
// 2. 将 cur.Next 指向 pre
// 3. 向后移动 cur、pre 指针
func reverseList(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode = nil
	for cur != nil {
		pre, cur, cur.Next = cur, cur.Next, pre
	}
	return pre
}
