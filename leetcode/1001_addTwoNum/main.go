package main

import (
	"fmt"

	. "../common"
)

/*
You are given two non-empty linked lists representing two non-negative integers.
The digits are stored in reverse order and each of their nodes contain a single digit.
Add the two numbers and return it as a linked list.

You may assume the two numbers do not contain any leading zero, except the number 0 itself.

Example

Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
Output: 7 -> 0 -> 8
Explanation: 342 + 465 = 807.
*/
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var p, p1, p2 *ListNode
	p1, p2 = l1, l2
	if p1 == nil && p2 == nil {
		return p
	}
	p = &ListNode{}
	pp := p
	var carry int = 0 // 方便进位
	for p1 != nil || p2 != nil || carry > 0 {

		if p1 != nil && p2 != nil {
			pp.Val = p1.Val + p2.Val
			p1, p2 = p1.Next, p2.Next
		} else if p1 == nil && p2 != nil {
			pp.Val = p2.Val
			p2 = p2.Next
		} else if p1 != nil && p2 == nil {
			pp.Val = p1.Val
			p1 = p1.Next
		}
		pp.Val += carry
		if pp.Val >= 10 {
			carry = 1
			pp.Val %= 10
		} else {
			carry = 0
		}
		if p1 != nil || p2 != nil || carry > 0 {
			pp.Next = &ListNode{}
			pp = pp.Next
		}
	}
	return p
}

func main() {
	/*



			a := &ListNode{
				Val: 5,
			}
			b := &ListNode{
				Val: 5,
			}
		a := &ListNode{
			Val: 9,
			Next: &ListNode{
				Val: 8,
			},
		}
		b := &ListNode{
			Val: 1,
		}
	*/
	a := &ListNode{
		Val: 2,
		Next: &ListNode{
			Val: 4,
			Next: &ListNode{
				Val:  3,
				Next: nil,
			},
		},
	}
	b := &ListNode{
		Val: 5,
		Next: &ListNode{
			Val: 6,
			Next: &ListNode{
				Val:  4,
				Next: nil,
			},
		},
	}
	c := addTwoNumbers(a, b)
	for c != nil {
		fmt.Printf("val = %d\n", c.Val)
		c = c.Next
	}

}
