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

//输入：head = [1,4,3,2,5,2], x = 3
//输出：[1,2,2,4,3,5]
// 1 	4
// 1 2	4 3
// 122 435
func partition(head *ListNode, x int) *ListNode {

	small := &ListNode{}
	smallCur := small
	large := &ListNode{}
	largeCur := large
	for head != nil {
		if head.Val < x {
			smallCur.Next = head
			smallCur = smallCur.Next
		} else {
			largeCur.Next = head
			largeCur = largeCur.Next
		}
		head = head.Next
	}
	largeCur.Next = nil
	smallCur.Next = large.Next
	return small.Next
}

func main() {
	list := &ListNode{Val:1}
	list.Next = &ListNode{Val:4}
	list.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next = &ListNode{Val:2}
	list.Next.Next.Next.Next = &ListNode{Val:5}
	list.Next.Next.Next.Next.Next = &ListNode{Val:2}
	printList(partition(list, 3))
}
