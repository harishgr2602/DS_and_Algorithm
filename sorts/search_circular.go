package main

import "fmt"

type CircularLinkedList struct {
	tail  *Node
	count int
}

type Node struct {
	value int
	next  *Node
}

func (list *CircularLinkedList) Size() int {
	return list.count
}

func (list *CircularLinkedList) IsEmpty() bool {
	return list.count == 0
}

func (list *CircularLinkedList) Peek() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("Empty list Error")
		return 0, false
	}
	return list.tail.next.value, true
}

func (list *CircularLinkedList) IsPresent(data int) bool {
	temp := list.tail
	for i := 0; i < list.count; i++ {
		if temp.value == data {
			return true
		}
		temp = temp.next
	}
	return false
}
