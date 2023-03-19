package algorithm

import (
	"fmt"
	"log"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

// 数组转链表
func genLinkList(arr []int) *ListNode {
	if len(arr) <= 0 {
		return nil
	}
	var head *ListNode
	var cur *ListNode
	for i := 0; i < len(arr); i++ {
		node := ListNode{
			Next: nil,
			Val:  arr[i],
		}
		if head == nil {
			head = &node
			cur = &node
			continue
		}
		cur.Next = &node
		cur = &node
	}
	return head
}

// 打印链表
func printLinkList(head *ListNode) {
	val := ""
	for {
		if head == nil {
			break
		}
		val = fmt.Sprintf("%v -> %v", val, head.Val)
		head = head.Next
	}
	log.Print(val)
}
