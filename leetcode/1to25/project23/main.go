package main

import (
	"fmt"
)
//输入：lists = [[1,4,5],[1,3,4],[2,6]]
//输出：[1,1,2,3,4,4,5,6]
//解释：链表数组如下：
//[
//1->4->5,
//1->3->4,
//2->6
//]
//将它们合并到一个有序链表中得到。
//1->1->2->3->4->4->5->6
type ListNode struct{
	Val int
	Next *ListNode
}

func mergeKLists(lists []*ListNode) *ListNode {
	return merge(lists, 0, len(lists)-1)
}

func merge(lists []*ListNode, l, r int) *ListNode {

	if l == r {
		return lists[l]
	}

	if l > r {
		return nil
	}

	middle := (l + r)/2
	return mergeTwoLists(merge(lists, l, middle), merge(lists, middle+1, r))
}

//mergeTwoLists
func mergeTwoLists(l, r *ListNode) *ListNode {
	if l == nil {
		return r
	}
	if r == nil {
		return l
	}
	if l.Val > r.Val {
		r.Next = mergeTwoLists(l, r.Next)
		return r
	} else {
		l.Next = mergeTwoLists(l.Next, r)
		return l
	}
}

func printList(node *ListNode) {
	for node != nil {
		fmt.Println(node.Val)
		node = node.Next
	}
}

func main() {
	list1 := &ListNode{Val:1}
	list1.Next = &ListNode{Val:4}
	list1.Next.Next = &ListNode{Val:5}
	list2 := &ListNode{Val:1}
	list2.Next = &ListNode{Val:3}
	list2.Next.Next = &ListNode{Val:4}
	list3 := &ListNode{Val:2}
	list3.Next = &ListNode{Val:6}
	printList(mergeKLists([]*ListNode{nil, list1, list2, list3}))
}
