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

//输入：head = [1,2,3,4,5], k = 2
//输出：[4,5,1,2,3]
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil {
		return head
	}
	count := head
	ct := 1
	for count.Next != nil {
		ct++
		count = count.Next
	}
	mod := k%ct
	if mod == 0 {
		return head
	}
	// 最后节点指向链头
	count.Next = head

	for i := 0; i < ct-mod; i++ {
		count = count.Next
	}

	// 当前节点的下一节点就是头
	res := count.Next
	// 更新为空，断开循环
	count.Next = nil

	return res
}

func main() {
	list1 := &ListNode{Val:1}
	list1.Next = &ListNode{Val:2}
	list1.Next.Next = &ListNode{Val:3}
	list1.Next.Next.Next = &ListNode{Val:4}
	list1.Next.Next.Next.Next = &ListNode{Val:5}
	printList(rotateRight(list1, 2))
}

