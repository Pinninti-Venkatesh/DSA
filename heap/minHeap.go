package heap

type MinHeap struct {
	data []int
}

func parent(i int) int { return (i - 1) / 2 }
func left(i int) int   { return 2*i + 1 }
func right(i int) int  { return 2*i + 2 }

func (h *MinHeap) swap(i, j int) {
	h.data[i], h.data[j] = h.data[j], h.data[i]
}

func NewMinHeap() *MinHeap {
	return &MinHeap{data: make([]int, 0)}
}

func (h *MinHeap) siftUp(i int) {
	for i > 0 && h.data[i] < h.data[parent(i)] {
		// Violation: current element is smaller than its parent (Min-Heap violation).
		// Swap the current element with its parent.
		h.swap(i, parent(i))
		// Move up to the parent's index for the next check.
		i = parent(i)
	}
}

func (h *MinHeap) Insert(val int) {
	h.data = append(h.data, val)
	h.siftUp(len(h.data) - 1)
}

func (h *MinHeap) SiftDown(i int) {
	n := len(h.data)
	for {
		l, r := left(i), right(i)
		smallest := i
		if l < n && h.data[l] < h.data[smallest] {
			smallest = l
		}
		if r < n && h.data[r] < h.data[smallest] {
			smallest = r
		}
		if i == smallest {
			break
		}
		h.swap(i, smallest)
		i = smallest
	}
}

func (h *MinHeap) ExtractMin() (int, bool) {
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
