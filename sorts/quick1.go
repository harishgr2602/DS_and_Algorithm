package main

import (
	"fmt"
)

func QuickSelect(arr []int, key int) int {
	size := len(arr)
	QuickSelectUtil(arr, 0, size-1, key)
	return arr[key-1]
}

func QuickSelectUtil(arr []int, lower int, upper int, key int) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for arr[lower] <= pivot && lower < upper {
			lower++
		}
		for arr[upper] > pivot && lower <= upper {
			upper--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)
	QuickSelectUtil(arr, start, upper-1, key)
	QuickSelectUtil(arr, upper+1, stop, key)
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func main() {
	data := []int{9, 1, 5, 2, 6, 3, 7, 4}
	fmt.Println(QuickSelect(data, 5))
}
