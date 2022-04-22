package main

import "fmt"

type ListNode struct {
	Val int
	Next *ListNode
}

//输入：l1 = [2,4,3], l2 = [5,6,4]
//输出：[7,0,8]
//解释：342 + 465 = 807.
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var head *ListNode
	var tail *ListNode
	carry := 0
	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		sum := n1 + n2 + carry
		sum, carry = sum%10, sum/10
		if head == nil {
			head = &ListNode{Val:sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val:sum}
			tail = tail.Next
		}
	}
	if carry > 0 {
		tail.Next = &ListNode{Val:carry}
	}
	return head
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main()  {
	l1 := &ListNode{
		Val:  2,
		Next: nil,
	}
	l1.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	l1.Next.Next = &ListNode{
		Val:  3,
		Next: nil,
	}
	l2 := &ListNode{
		Val:  5,
		Next: nil,
	}
	l2.Next = &ListNode{
		Val:  6,
		Next: nil,
	}
	l2.Next.Next = &ListNode{
		Val:  4,
		Next: nil,
	}
	res := addTwoNumbers(l1, l2)
	printList(res)
}