// Package provides singly Linked List implementation
package main

import "fmt"

type MyLinkedList struct {
	arr []int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}

func (this *MyLinkedList) Length() int {
	return len(this.arr)
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	n := this.Length()
	if index >= n || index < 0 {
		return -1
	} else {
		return this.arr[index]
	}
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	this.arr = append([]int{val}, this.arr...)
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
	this.arr = append(this.arr, val)
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	n := this.Length()
	if index > n || index < 0 {
		return
	} else {
		this.arr = append(this.arr[:index], append([]int{val}, this.arr[index:]...)...)
	}
}

func (this *MyLinkedList) Print() {
	fmt.Println(this.arr)
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	n := this.Length()
	if index >= n || index < 0 {
		return
	}
	this.arr = append(this.arr[:index], this.arr[index+1:]...)
}


/**
 * Your MyLinkedList object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Get(index);
 * obj.AddAtHead(val);
 * obj.AddAtTail(val);
 * obj.AddAtIndex(index,val);
 * obj.DeleteAtIndex(index);
 */
func testCase1() {
	obj := Constructor()
	obj.Print()
	obj.AddAtHead(3)
	obj.Print()
	obj.AddAtHead(2)
	obj.Print()
	obj.AddAtHead(1)
	obj.Print()
	obj.AddAtTail(4)
	obj.Print()
	obj.AddAtIndex(1, 22)
	obj.Print()
	obj.DeleteAtIndex(1)
	obj.Print()
}

func testCase2() {
	obj := Constructor()
	obj.AddAtHead(1)
	obj.Print()
	obj.AddAtIndex(1, 2)
	obj.Print()
	fmt.Println(obj.Get(1))
	fmt.Println(obj.Get(0))
	fmt.Println(obj.Get(2))
}

func main() {
	// testCase2()
	testCase1()
}
