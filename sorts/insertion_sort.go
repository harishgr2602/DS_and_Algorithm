package main

import (
	"fmt"
)

func InsertionSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	var temp, i, j int
	for i = 1; i < size; i++ {
		temp = arr[i]
		for j = i; j > 0 && comp(arr[j-1], temp); j-- {
			arr[j] = arr[j-1]
		}
		arr[j] = temp
	}
}

func main() {
	data := []int{9, 3, 1, 7, 4, 2, 0, 5, 8}
	InsertionSort(data)
	fmt.Println(data)
}
