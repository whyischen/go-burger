package algorithm

import (
	"log"
	"testing"
)

func TestReverseList(t *testing.T) {
	head := genLinkList([]int{1, 2, 3, 4})
	printLinkList(head)

	head = reverseList(head)
	printLinkList(head)
}

func TestQuickSort(t *testing.T) {
	arr := []int{3, 2, 5, 8, 1}
	QuickSort(arr)
	log.Printf("TestQuickSort result: %v", arr)
}

func TestMergeTwoList(t *testing.T) {
	l1 := genLinkList([]int{1, 2})
	l2 := genLinkList([]int{1})
	printLinkList(mergeTwoList1(l1, l2))
}
