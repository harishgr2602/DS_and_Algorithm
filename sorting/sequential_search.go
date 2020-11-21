package main

import (
	"fmt"
)

func SequentialSearch(data []int, value int) bool {
	size := len(data)
	for i := 0; i < size; i++ {
		if value == data[i] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println("Data is:", data)
}
