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

func (list *CircularLinkedList) AddHead(value int) {
	temp := &Node{value, nil}
	if list.IsEmpty() {
		list.tail = temp
		temp.next = temp
	} else {
		temp.next = list.tail.next
		list.tail.next = temp
	}
	list.count++
}

func(list *CircularLinkedList) RemoveHead() (int, bool) {
	if list.IsEmpty() {
		fmt.Println("Empty List Error")
		return 0, fasle
	}
	value := list.tail.next.value

	if list.tail == list.tail.next {
		list.tail = nil
	} else {
		list.tail.next = liat.tail.next.next
	}
	list count--
	return value, true
}
