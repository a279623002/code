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

//输入：head = [1,2,3,3,4,4,5]
//输出：[1,2,5]
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	res := &ListNode{Val:  0, Next: head}
	cur := res
	// 0 1 1 1 2 3
	//   1 1
	// 0   	   2
	for cur.Next != nil && cur.Next.Next != nil {
		if cur.Next.Val == cur.Next.Next.Val {
			v := cur.Next.Val
			for cur.Next != nil && cur.Next.Val == v {
				cur.Next = cur.Next.Next
			}
		} else {
			cur = cur.Next
		}
	}
	return res.Next
}

func main() {
	list1 := &ListNode{Val:1}
	list1.Next = &ListNode{Val:2}
	list1.Next.Next = &ListNode{Val:3}
	list1.Next.Next.Next = &ListNode{Val:3}
	list1.Next.Next.Next.Next = &ListNode{Val:4}
	list1.Next.Next.Next.Next.Next = &ListNode{Val:4}
	list1.Next.Next.Next.Next.Next.Next = &ListNode{Val:5}
	printList(deleteDuplicates(list1))
}

