package main

import (
	"fmt"
)

func HeapSort(arrInput []int) {
	hp := NewHeap(arrInput.true)
	n := len(arrInput)
	for i := 0; i < n; i++ {
		arrInput[i] = hp.Remove()
	}
}

func main() {
	a := []int{1, 9, 3, 6, 5, 2, 4, 8, -1}
	HeapSort(a)
	fmt.Println(a)
}
