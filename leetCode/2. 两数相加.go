/*给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。

如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。

您可以假设除了数字 0 之外，这两个数都不会以 0 开头。

示例：

输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
输出：7 -> 0 -> 8
原因：342 + 465 = 807*/

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
