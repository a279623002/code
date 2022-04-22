package main

import "fmt"

//输入：head = [1,2,3,4]
//输出：[2,1,4,3]
type ListNode struct{
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 链表头节点指向的第二节点是新链表的头节点
	newHead := head.Next
	// 头节点指向递归第三节点
	head.Next = swapPairs(newHead.Next)
	// 新链表的第二节点指向头节点，此时头节点是指向了递归第三节点
	newHead.Next = head
	return newHead
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
	list1.Next.Next = &ListNode{Val:3}
	list1.Next.Next.Next = &ListNode{Val:4}
	printList(swapPairs(list1))
}
