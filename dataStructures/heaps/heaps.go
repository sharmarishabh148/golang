package main

import (
	"container/heap"
	"fmt"
)

//According to the heap order (maximum heap) property, the value
//stored at each node is greater than or equal to its children.
//If the order is descending, it is referred to as a maximum heap

// implementing minHeap
type IntegerMinHeap []int

func (iHeap IntegerMinHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntegerMinHeap) Less(i, j int) bool {
	//MinHeap property [0]th is minimum
	return iHeap[i] < iHeap[j]
}

func (iHeap IntegerMinHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerMinHeap) Push(x interface{}) {
	*iHeap = append(*iHeap, x.(int))
}

func (iHeap *IntegerMinHeap) Pop() interface{} {
	var n int
	var x1 int
	var previous IntegerMinHeap = *iHeap
	n = len(previous)
	x1 = previous[n-1]
	*iHeap = previous[0 : n-1]
	return x1

}

type IntegerMaxHeap []int

func (iHeap IntegerMaxHeap) Len() int {
	return len(iHeap)
}

func (iHeap IntegerMaxHeap) Less(i, j int) bool {
	//MinHeap property [0]th is maximum
	return iHeap[i] > iHeap[j]
}

func (iHeap IntegerMaxHeap) Swap(i, j int) {
	iHeap[i], iHeap[j] = iHeap[j], iHeap[i]
}

func (iHeap *IntegerMaxHeap) Push(x interface{}) {
	*iHeap = append(*iHeap, x.(int))
}

func (iHeap *IntegerMaxHeap) Pop() interface{} {
	var n int
	var x1 int
	var old IntegerMaxHeap = *iHeap
	n = len(old)
	x1 = old[n-1]
	*iHeap = old[0 : n-1]
	return x1

}
func main() {
	var intHeap *IntegerMinHeap = &IntegerMinHeap{1, 4, 5}
	heap.Init(intHeap)
	heap.Push(intHeap, 2)
	fmt.Printf("minimum: %d\n", (*intHeap)[0])
	for intHeap.Len() > 0 {
		fmt.Printf("%d \n", heap.Pop(intHeap))
	}

	h := &IntegerMaxHeap{3, 1, 4, 1, 5, 9, 2, 6, 5, 3, 5}
	heap.Init(h)
	heap.Push(h, 10) // Push an element onto the max-heap
	//pop := heap.Pop(h).(int) // Pop the maximum element from the max-heap
	fmt.Printf("maximum: %d\n", (*h)[0])
	for h.Len() > 0 {
		max := heap.Pop(h).(int)
		fmt.Println(max)
	}

}
