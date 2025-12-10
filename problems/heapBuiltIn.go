package Problems

import (
	"container/heap"
	"fmt"
)

type IntHeap []int

func (h IntHeap) Len() int {
	return len(h)
}

func (h IntHeap) Less(i, j int) bool {
	return h[i] < h[j]
}

func (h IntHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *IntHeap) Push(x interface{}) {
	*h = append(*h, x.(int))
}

func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func main() {
	// 1. Initialize the Heap
	h := &IntHeap{2, 1, 5}
	heap.Init(h) // This calls heapify to arrange the initial elements

	fmt.Printf("Min-Heap root (smallest): %d\n", (*h)[0]) // Output: 1

	// 2. Push (Insert)
	heap.Push(h, 3)
	fmt.Printf("After Push(3), root: %d\n", (*h)[0]) // Output: 1 (still the smallest)

	// 3. Pop (Extract Minimum)
	fmt.Printf("Extracted minimum: %d\n", heap.Pop(h)) // Output: 1
	fmt.Printf("New root: %d\n", (*h)[0])              // Output: 2
}
