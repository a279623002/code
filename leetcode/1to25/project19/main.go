package main

import (
	"fmt"
)

type ListNode struct{
	Val int
	Next *ListNode
}

//输入：head = [1,2,3,4,5], n = 2
//输出：[1,2,3,5]
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	// 很巧妙的定义，最后 next返回的就是结果，不用对n进行-1
	res := &ListNode{0, head}
	fast, slow := head, res
	for i := 0; i < n; i++ {
		// 到n停
		fast = fast.Next
	}
	// fast能继续的次数就是slow到n的位置 index = len - n
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}
	// 指向位置n的下一个
	slow.Next = slow.Next.Next
	return res.Next
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	list := &ListNode{Val:1}
	list.Next = &ListNode{Val:2}
	list.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next = &ListNode{Val:4}
	list.Next.Next.Next.Next = &ListNode{Val:5}
	printList(removeNthFromEnd(list, 2))
}