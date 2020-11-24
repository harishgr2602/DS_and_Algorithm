package main

import (
	"fmt"
)

func MergeSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	temp := make([]int, size)
	mergeSrt(arr, temp, 0, size-1, comp)
}

func Merge(arr []int, temp []int, LIndex int, MIndex int, UIndex int, comp func(int, int) bool) {
	LStart := LIndex
	LStop := MIndex
	UStart := MIndex + 1
	UStop := UIndex
	count := LIndex
	for LStart <= LStop && UStart <= UStop {
		if comp(arr[LStart], arr[UStart]) == false {
			temp[count] = arr[LStart]
			LStart++
		} else {
			temp[count] = arr[UStart]
			UStart++
		}
		count++
	}
	for LStart <= LStop {
		temp[count] = arr[LStart]
		count++
		LStart++
	}
	for UStart <= UStop {
		temp[count] = arr[UStart]
		count++
		UStart++
	}
	for i := LIndex; i <= UIndex; i++ {
		arr[i] = temp[i]
	}
}

func mergeSrt(arr []int, temp []int, LIndex int, UIndex int, comp func(int, int) bool) {
	if LIndex >= UIndex {
		return
	}
	MIndex := (LIndex + UIndex) / 2
	mergeSrt(arr, temp, LIndex, MIndex, comp)
	mergeSrt(arr, temp, MIndex+1, UIndex, comp)
	Merge(arr, temp, LIndex, MIndex, UIndex, comp)
}

func main() {
	data := more []int{9, 1, 8, 2, 7, 3, 6, 4, 5}
	MergeSort(data, more)
	fmt.Println(data)
}
