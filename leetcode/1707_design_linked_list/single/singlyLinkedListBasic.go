// Package provides singly Linked List implementation
// 实现单链表注意点
// Get Delete等index判断应该提前判断并返回
// 只用一个头指针表示node的话，速度在尾插入的时候会比较慢
// go指针不能显示的改变
// 例如:
// type L struct {
//	val int
//  next *L
// }
// func (l *L) addHead(val int) {
//   p := l
//   newL := L{val:val, next: p}
//   *l = &newL      // wrong
//   l = newL        // could not work, still wrong
//   // 这里发现l的地址是改变不了, 不知道为什么
//}
package main

import "fmt"

// 改成头尾节点更加好
//type NodeTwoSide struct {
	//val int
	//head *NodeTwoSide
	//tail *NodeTwoSide
//}

type Node struct {
	val int
	next *Node
}

type MyLinkedList struct {
	pn *Node
	length int
}


/** Initialize your data structure here. */
func Constructor() MyLinkedList {
	return MyLinkedList{}
}


/** Get the value of the index-th node in the linked list. If the index is invalid, return -1. */
func (this *MyLinkedList) Get(index int) int {
	if index < 0 || index >= this.length {
		return -1
	}
    pHead := this.pn
    i := 0
    for pHead != nil {
        if index == i {
			break
        }
        i++
        pHead = pHead.next
    }
	return pHead.val
}


/** Add a node of value val before the first element of the linked list. After the insertion, the new node will be the first node of the linked list. */
func (this *MyLinkedList) AddAtHead(val int)  {
	pHead := this.pn
    newNode := &Node{val: val, next: pHead}
	this.pn = newNode
	this.length++
}


/** Append a node of value val to the last element of the linked list. */
func (this *MyLinkedList) AddAtTail(val int)  {
    pTail := this.pn
    for pTail.next != nil {
		pTail = pTail.next
	}
	newNode := &Node{val: val, next: nil}
    pTail.next = newNode
	this.length++
}


/** Add a node of value val before the index-th node in the linked list. If index equals to the length of linked list, the node will be appended to the end of linked list. If index is greater than the length, the node will not be inserted. */
func (this *MyLinkedList) AddAtIndex(index int, val int)  {
	if this.length < index || index < 0 {
		return
	} else if index == 0 {
		this.AddAtHead(val)
		return
	} else {
		pHead := this.pn
		i := 1
		for pHead != nil {
			if i == index {
				newNode := &Node{val: val, next: pHead.next}
				pHead.next = newNode
				this.length++
				break
			}
			i++
			pHead = pHead.next
		}
	}
}


/** Delete the index-th node in the linked list, if the index is valid. */
func (this *MyLinkedList) DeleteAtIndex(index int)  {
	if index >= this.length || index < 0 {
		return
	}
    pHead := this.pn
	if index == 0 {
		this.pn = this.pn.next
		this.length--
        // how about origin first node?
		return
	}
	i := 1
	for pHead != nil {
		if i == index {
			// left one not deal with
			pHead.next = pHead.next.next
			this.length--
			break
		}
		i++
		pHead = pHead.next
	}
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

func (this *MyLinkedList) Print() {
	pHead := this.pn
	for pHead != nil {
		fmt.Printf("%d ", pHead.val)
		pHead = pHead.next
	}
	fmt.Printf(" with length: %d\n", this.length)
}

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
	testCase2()
}
