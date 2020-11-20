package main

import (
	"fmt"
)

func binarySearch(data int, value []int) bool {

	low := 0
	high := len(value) - 1

	for low <= high {
		median := (low + high) / 2

		if value[median] < data {
			low = median + 1
		} else {
			high = median - 1
		}
	}

	if low == len(value) || value[low] != data {
		return false
	}

	return true
}

func main() {
	items := []int{1, 2, 9, 20, 31, 45, 63, 70, 100}
	fmt.Println(binarySearch(63, items))
}
