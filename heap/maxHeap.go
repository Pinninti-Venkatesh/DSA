package heap

type MaxHeap struct {
	data []int
}

func (h *MaxHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func NewMaxHeap() *MaxHeap {
	return &MaxHeap{data: make([]int, 0)}
}

func (h *MaxHeap) siftUp(i int) {
	for i > 0 && h.data[i] > h.data[parent(i)] {
		// Violation: current element is smaller than its parent (Min-Heap violation).
		// Swap the current element with its parent.
		h.swap(i, parent(i))
		// Move up to the parent's index for the next check.
		i = parent(i)
	}
}

func (h *MaxHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MaxHeap) SiftDown(i int) {
	n := len(h.data)
	for {
		l, r := left(i), right(i)
		smallest := i
		if l < n && h.data[l] > h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] > h.data[smallest] {
			smallest = r
		}
		if i == smallest {
			break
		}
		h.swap(i, smallest)
		i = smallest
	}
}

func (h *MaxHeap) ExtractMax() (int, bool) {
	n := len(h.data)
	if n == 0 {
		return 0, false
	}
	minVal := h.data[0]
	h.data[0] = h.data[n-1]
	h.data = h.data[:n-1]
	if len(h.data) > 0 {
		h.SiftDown(0)
	}
	return minVal, true
}
