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

func (list *CircularLinkedList) RemoveNode(key int) bool {
	if list.IsEmpty() {
		return false
	}
	prev := list.tail
	curr := list.tail.next
	head := list.tail.next

	if curr.value == key {
		if curr == curr.next {
			list.tail = nil
		} else {
			list.tail.next = list.tail.next.next
		}
		return true
	}
	prev = curr
	curr = curr.next

	for curr != head {
		if curr.value == key {
			if curr == list.tail {
				list.tail = prev
			}
			prev.next = curr.next
			return true
		}
		prev = curr
		curr = curr.next
	}
	return false
}
