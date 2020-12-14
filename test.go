package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	a1 := ListNode{7, nil}
	a2 := ListNode{1, &a1}
	a3 := ListNode{6, &a2}

	b1 := ListNode{5, nil}
	b2 := ListNode{1, &b1}
	b3 := ListNode{7, &b2}

	res := addTwoNumbers(&a3, &b3)
	fmt.Println(res.Val, res.Next.Val, res.Next.Next.Val, res.Next.Next.Next.Val)
}

func addTwoNumbers(l1, l2 *ListNode) (head *ListNode) {
	var tail *ListNode
	var carry int

	for l1 != nil || l2 != nil {
		var a, b int

		if l1 != nil {
			a = l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			b = l2.Val
			l2 = l2.Next
		}

		sum := a + b + carry
		sum, carry = sum%10, sum/10

		if head == nil {
			head = &ListNode{Val: sum}
			tail = head
		} else {
			tail.Next = &ListNode{Val: sum}
			tail = tail.Next
		}
	}

	if carry > 0 {
		tail.Next = &ListNode{Val: carry}
	}

	return
}
