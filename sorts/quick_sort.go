package main

import "fmt"

func QuickSort(arr []int, comp func(int, int) bool) {
	size := len(arr)
	quickSortUtil(arr, 0, size-1, comp)
}

func quickSortUtil(arr []int, lower int, upper int, comp func(int, int) bool) {
	if upper <= lower {
		return
	}
	pivot := arr[lower]
	start := lower
	stop := upper

	for lower < upper {
		for comp(arr[lower], pivot) == false && lower < upper {
			lower++
		}
		for comp(arr[upper], pivot) && lower <= upper {
			lower--
		}
		if lower < upper {
			swap(arr, upper, lower)
		}
	}
	swap(arr, upper, start)
	quickSortUtil(arr, start, upper-1, comp)
	quickSortUtil(arr, upper+1, stop, comp)
}

func swap(arr []int, first int, second int) {
	arr[first], arr[second] = arr[second], arr[first]
}

func main() {
	data := []int{9, 1, 8, 4, 7, 2, 0, 3, 5}
	QuickSort(data, more)
	fmt.Println(data)
}
