package main

import "fmt"

type ListNode struct{
	Val int
	Next *ListNode
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

//输入：head = [1,1,2,3,3]
//输出：[1,2,3]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	// 1->1->2->3->3	=> 1->2->3
	// 1->2->3->3		=> 1->2->3
	// 2->3->3			=> 2->3
	// 3->3				=> 3
	// 3
	head.Next = deleteDuplicates(head.Next)
	if head.Val == head.Next.Val {
		head = head.Next
	}
	return head
}

func deleteDuplicates1(head *ListNode) *ListNode {
	cur := head
	for cur != nil && cur.Next != nil {
		if cur.Val == cur.Next.Val {
			cur.Next = cur.Next.Next
		} else {
			cur = cur.Next
		}
	}
	return head
}

func main() {
	list := &ListNode{Val:1}
	list.Next = &ListNode{Val:1}
	list.Next.Next = &ListNode{Val:1}
	list.Next.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next.Next = &ListNode{Val:3}
	printList(deleteDuplicates(list))
}
