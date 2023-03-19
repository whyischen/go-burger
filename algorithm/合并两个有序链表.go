package algorithm

// 21. 合并两个有序链表
// 将两个升序链表合并为一个新的 升序 链表并返回。新链表是通过拼接给定的两个链表的所有节点组成的。
// 示例1：
// 输入：l1 = [1,2,4], l2 = [1,3,4]
// 输出：[1,1,2,3,4,4]
// 迭代解法
func mergeTwoList1(list1, list2 *ListNode) *ListNode {
	root := &ListNode{}
	head := root
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			head.Next = list1
			list1 = list1.Next
		} else {
			head.Next = list2
			list2 = list2.Next
		}
		head = head.Next
	}

	if list1 != nil {
		head.Next = list1
	} else {
		head.Next = list2
	}

	return root.Next
}
