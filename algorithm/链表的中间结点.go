package algorithm

// 876. 链表的中间结点
// 给你单链表的头结点 head ，请你找出并返回链表的中间结点。
// 如果有两个中间结点，则返回第二个中间结点。

// 双指针思路，使用快慢指针
// 链表长度分奇偶两种情况考虑
func middleNode(head *ListNode) *ListNode {
	low, fast := head, head
	for fast.Next != nil {
		low = low.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil {
		low = low.Next
	}
	return low
}
