package heaps


type Heap []int


func NewHeap(l []int) *Heap {
	h := Heap(l)
	return &h
}


func (h *Heap) get(index int) (*int, bool) {
	if index >= h.len() {
		return nil, false
	}

	return &([]int(*h)[index]), true
}


func (h *Heap) parent(index int) int {
	return index >> 1
}


func (h *Heap) left(index int) int {
	return index << 1
}


func (h *Heap) right(index int) int {
	return (index << 1) | 1
}


func (h *Heap) len() int {
	return len([]int(*h))
}


func (h *Heap) Swap(i, j int) {
	if i >= h.len() || j >= h.len() {
		return
	}

	l := []int(*h)
	l[i], l[j] = l[j], l[i]
}


func (h *Heap) maxHeapify(i, r int) {
	index := i
	for index < r {
		current, _ := h.get(index)
		maxIndex := index
		max := *current
		leftIndex := h.left(index)
		left, exists := h.get(leftIndex)
		if exists && max < *left {
			maxIndex = leftIndex
			max = *left
		}

		rightIndex := h.right(index)
		right, exists := h.get(rightIndex)
		if exists && max < *right {
			maxIndex = rightIndex
		}

		if maxIndex != index {
			h.Swap(maxIndex, index)
			index = maxIndex
		} else {
			return
		}

	}
}


func (h *Heap) MaxHeapify() {
	hLen := h.len()
	for i := hLen >> 1; i >= 0; i-- {
		h.maxHeapify(i, hLen)
	}
}


func (h *Heap) Sort() {
	h.MaxHeapify()
	for i := h.len(); i > 0; i-- {
		h.Swap(0, i)
		h.maxHeapify(0, i - 1)
	}
}
