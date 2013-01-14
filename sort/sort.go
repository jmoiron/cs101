package main

import (
	"fmt"
	"math/rand"
)

func randomList(length ...int) []int {
	size := 10
	if len(length) == 0 {
		size = length[0]
	}
	list := make([]int, size)
	for i := 0; i < size; i++ {
		list[i] = rand.Int() / 100000
	}
	return list
}

func isSorted(list []int) bool {
	if len(list) <= 1 {
		return true
	}
	i := 0
	j := 1
	for ; j < len(list); j++ {
		if list[i] > list[j] {
			return false
		}
		i++
	}
	return true
}

func partition(list []int, left, right, pivot int) int {
	saved := left
	pval := list[pivot]
	list[pivot], list[right] = list[right], list[pivot]
	for i := left; i < right; i++ {
		if list[i] < pval {
			list[i], list[saved] = list[saved], list[i]
			saved++
		}
	}
	list[saved], list[right] = list[right], list[saved]
	return saved
}

func _quicksort(list []int, left, right int) {
	if left < right {
		pivot := left + (right-left)/2
		newpivot := partition(list, left, right, pivot)
		_quicksort(list, left, newpivot-1)
		_quicksort(list, newpivot+1, right)
	}
}

// perform quicksort on a list of integers in place
func qsort(list []int) {
	_quicksort(list, 0, len(list)-1)
}

func heapify(list []int) {
	size := len(list)
	for start := (size - 2) / 2; start >= 0; start-- {
		siftDown(list, start, size-1)
	}
}

func siftDown(list []int, start, end int) {
	for root := start; root*2+1 <= end; {
		child := root*2 + 1
		swap := root
		/* find the greatest number between root & children */
		if list[swap] < list[child] {
			swap = child
		}
		if child+1 <= end && list[swap] < list[child+1] {
			swap = child + 1
		}
		if swap != root {
			list[root], list[swap] = list[swap], list[root]
			root = swap
		} else {
			return
		}
	}
}

// perform heapsort on a list of integers
func heapsort(list []int) {
	heapify(list)

	for end := len(list) - 1; end > 0; {
		list[end], list[0] = list[0], list[end]
		end--
		siftDown(list, 0, end)
	}
}

func main() {
	fmt.Println("qsort")
	for i := 0; i < 10; i++ {
		l := randomList(10)
		fmt.Println(l, isSorted(l))
		qsort(l)
		fmt.Println(l, isSorted(l))
		fmt.Println()
	}

	fmt.Println("heapsort")
	for i := 0; i < 10; i++ {
		l := randomList(10)
		fmt.Println(l, isSorted(l))
		heapsort(l)
		fmt.Println(l, isSorted(l))
		fmt.Println()
	}
}
