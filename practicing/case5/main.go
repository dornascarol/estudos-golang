package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummy := &ListNode{}
	current := dummy
	carry := 0

	for l1 != nil || l2 != nil || carry != 0 {
		sum := carry

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}

		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}

		carry = sum / 10
		current.Next = &ListNode{Val: sum % 10}
		current = current.Next
	}

	return dummy.Next
}

func buildList(nums []int) *ListNode {
	dummy := &ListNode{}
	current := dummy

	for _, v := range nums {
		current.Next = &ListNode{Val: v}
		current = current.Next
	}

	return dummy.Next
}

func printList(head *ListNode) {
	for head != nil {
		fmt.Print(head.Val)
		if head.Next != nil {
			fmt.Print(" -> ")
		}
		head = head.Next
	}
	fmt.Println()
}

func main() {
	l1 := buildList([]int{2, 4, 3})
	l2 := buildList([]int{5, 6, 4})

	fmt.Print("Lista 1: ")
	printList(l1)

	fmt.Print("Lista 2: ")
	printList(l2)

	result := addTwoNumbers(l1, l2)

	fmt.Print("Resultado: ")
	printList(result)
}
