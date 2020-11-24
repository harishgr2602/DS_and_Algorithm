package main

import (
	"fmt"
)

func BucketSort(data []int, lowerRange int, upperRange int) {
	mg := upperRange - lowerRange
	size := len(data)
	count := make([]int, mg)
	for i := 0; i < mg; i++ {
		count[i] = 0
	}
	for i := 0; i < size; i++ {
		count[data[i]-lowerRange]++
	}
	j := 0
	for i := 0; i < mg; i++ {
		for ; count[i] > 0; count[i]-- {
			data[j] = 1 + lowerRange
			j++
		}
	}
}

func main() {
	data := []int{23, 24, 22, 21, 26, 25, 27, 28, 21, 21}
	BucketSort(data, 30, 20)
	fmt.Println(data)
}
