package main

import (
	"fmt"

	. "../leetcode/common"
)

func test1() {
	node1 := ListNode{Val: 4, Next: nil}
	node2 := ListNode{Val: 5, Next: nil}
	node3 := ListNode{}
	node3.Val = node1.Val + node2.Val
	fmt.Printf("Node 3 Val = %d\n", node3.Val)
}

func test2() {
	node1 := &ListNode{Val: 4, Next: nil}
	node2 := &ListNode{Val: 5, Next: nil}
	node3 := &ListNode{}
	node3.Val = node1.Val + node2.Val
	fmt.Printf("Node 3 Val = %d\n", node3.Val)
}

func test3() {
	var node1, node2, node3 *ListNode // simple node1,2,3 is nil
	// Wrong initial way:
	// node1.Val, node2.Val, node3.Val = 4, 5, 6
	// fmt.Printf("Node1 Value=%d, Node1 Next=%s", &node1.Val, node1.Next)
	node1 = &ListNode{Val: 4}
	node2 = &ListNode{Val: 5}
	node3 = &ListNode{}
	node3.Val = node1.Val + node2.Val
	fmt.Printf("Node 3 Val = %d, Next: %s\n", node3.Val, node3.Next)
}

func main() {
	fmt.Println("Test with direct using:")
	test1()
	fmt.Println("Test with pointer:")
	test2()
	fmt.Println("Wrong decrelation see commented code")
	test3()
}
