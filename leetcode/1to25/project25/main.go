package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

//输入：head = [1,2,3,4,5], k = 2
//输出：[2,1,4,3,5]
func reverseKGroup(head *ListNode, k int) *ListNode {
	tail := head
	for i:=0; i<k; i++ {
		// 后续节点不足k返回head
		if tail == nil {
			return head
		}
		tail = tail.Next
	}

	// 反转这一截链表, prev第一个节点将会是最后一个节点 即head.Next要指向递归节点
	prev := head
	curr := head.Next
	for curr != tail {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}

	// 反转后的链表指向递归剩余一截链表
	head.Next = reverseKGroup(tail, k)

	// 返回反转后的头节点
	return prev
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
	list1.Next.Next.Next.Next = &ListNode{Val:5}
	printList(reverseKGroup(list1, 2))
}


