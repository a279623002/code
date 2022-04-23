package main

import (
	"fmt"
)

//输入：l1 = [1,2,4], l2 = [1,3,4]
//输出：[1,1,2,3,4,4]
type ListNode struct{
	Val int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}
	if l1.Val > l2.Val {
		l2.Next = mergeTwoLists(l1, l2.Next)
		return l2
	} else {
		l1.Next = mergeTwoLists(l1.Next, l2)
		return l1
	}
}


func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	list1 := &ListNode{Val:1}
	list1.Next = &ListNode{Val:2}
	list1.Next.Next = &ListNode{Val:4}
	list2 := &ListNode{Val:1}
	list2.Next = &ListNode{Val:3}
	list2.Next.Next = &ListNode{Val:4}
	printList(mergeTwoLists(list1, list2))
}
