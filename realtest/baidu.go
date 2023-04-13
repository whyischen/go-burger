package realtest

import "go-burger/algorithm"

// 乱序数组 to 二叉树
// 不能退化成链表

func insert(node *algorithm.TreeNode, val int) {
	var nextNode *algorithm.TreeNode
	if node.Val < val {
		nextNode = node.Right
	} else {
		nextNode = node.Left
	}

	if nextNode == nil {
		nextNode = &algorithm.TreeNode{
			Val: val,
		}
		return
	} else {
		insert(nextNode, val)
	}
}
