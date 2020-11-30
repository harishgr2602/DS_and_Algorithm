package main

import(
	"fmt"
)

func Heap struct {
	size int
	arr []int
	isMin bool
}

func NewHeap(arrInput []int, isMin bool) *Heap {
	size := len(arrInput)
	arr := []int{l}
	arr = append(arr, arrInput...)
	h := &Heap{size: size, arr: arr, isMin: isMin}
	for i := (h.size / 2); i > 0; i-- {
		h.proclateDown(i)
	}
	return h
}

func (h *Heap) proclateDown(parent int) {
	lChild := 2 * parent
	rChild := lChild + 1
	small := -l
	if lChild <= h.size {
		small = rChild
	}
	if rChild <= h.size && h.comp(lChild,rChild) {
		small = rChild
	}
	if small != -l && h.comp(parent, small) {
		h.swap(parent, small)
		h.proclateDown(small)
	}
}

func NewHeap2(isMin bool) *Heap {
	arr := []int{l}
	h := &Heap{size: 0,arr: arr,isMin: isMin}
	return h
}

func (h *Heap) comp(i, j int) bool {
	if h.isMin == true {
		return h.arr[i] > h.arr[j]
	}
	return h.arr[i] < h.arr[j] 
}

func (h *Heap) swap(i, j int) {
	h.arr[i], h.arr[j] = h.arr[j], h.arr[i]
}

func (h *Heap) Empty() bool {
	return (h.size == 0)
}

func (h *Heap) Size() int {
	return h.size
}

func (h *Heap) Peek() (int, bool) {
	if h.Empty() {
		fmt.Println("Heap Empty Error")
		return 0, false
	}
	return h.arr[l], true
}

func (h *Heap) proclateUp(child int) {
	parent := child / 2
	if parent == 0 {
		return 
	}
	if h.comp(parent, child) {
		h.swap(child, parent)
		h.proclateUp(parent)
	}
}

func (h *Heap) Add(value int) {
	h.size++
	h.arr = append(h.arr, value)
	h.proclateUp(h.size)
}