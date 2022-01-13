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
//输入：head = [1,2,3,4,5], left = 2, right = 4
//输出：[1,4,3,2,5]
func reverseBetween1(head *ListNode, left int, right int) *ListNode {
	run := head
	tmp := &ListNode{}
	tmp = tmp.Next

	for run != nil {
		next := run.Next
		run.Next = tmp
		tmp = run
		run = next
	}
	return tmp
}
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	run := head
	tmp := &ListNode{}
	start, end := &ListNode{}, &ListNode{} // 记录反转前头尾
	tmp = tmp.Next // 去掉0.使用nil
	count := 1
	for run != nil {
		next := run.Next
		if count >= left && count <= right {
			//开始反转
			if tmp == nil {
				end = run // 初始时记录反转链表的链尾
			}
			run.Next = tmp
			tmp = run
		}
		if count < left {
			start = run // 记录开始反转时的节点，记作指向反转链表的头的节点
		}
		if count > right {
			// 到达反转结束的节点退出，此时节点记作反转链表的尾的指向节点
			break
		}
		run = next
		count++
	}
	// 如果没有start,即left=1，使用tmp
	if start.Next != nil {
		start.Next = tmp
	} else {
		head = tmp
	}
	end.Next = run
	return head
}

func main() {
	list := &ListNode{Val:1}
	list.Next = &ListNode{Val:2}
	list.Next.Next = &ListNode{Val:3}
	list.Next.Next.Next = &ListNode{Val:4}
	list.Next.Next.Next.Next = &ListNode{Val:5}
	printList(reverseBetween(list, 1, 2))
}
