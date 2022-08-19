package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	//listNode := []ListNode{}
	slow, fast := head, head
	for fast.Next != nil && fast.Next.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast.Next != nil && fast.Next.Next == nil {
		slow = slow.Next
	}
	return fast
}

func createList() *ListNode {
	head := new(ListNode)
	return head
}

func (listNode *ListNode) appendList(num int) {
	temp := ListNode{Val: num}
	listNode.Next = &temp
}

func main() {
	head := new(ListNode)
	head.appendList(1)
	fmt.Println(head.Next.Val)
}
