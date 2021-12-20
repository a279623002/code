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
	tmp := &ListNode{Val:0}
	cur := head
	for cur != nil {
		if cur.Val >= x {
			tmp.Next = cur
			tmp = tmp.Next
		}
		cur = cur.Next
	}

	return head
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
