package heaps

import "errors"

type Comparable interface {
	Compare(Comparable) int8
}

type Heap []Comparable

func (heap *Heap) Parent(index int) int {
	if index == 0 {
		return 0
	}

	return (index - 1) >> 1
}

func (heap *Heap) Left(index int) int {
	return (index << 1) + 1
}

func (heap *Heap) Right(index int) int {
	return (index << 1) + 2
}

func (heap *Heap) Get(index int) (Comparable, error) {
	if index >= len(*heap) || index < 0 {
		return nil, errors.New("Index out of bounds")
	}

	return (*heap)[index], nil
}

func (heap *Heap) IsMaxHeap() bool {
	for i, c := range *heap {
		if i == 0 {
			continue
		}

		parent, _ := heap.Get(heap.Parent(i))
		if c.Compare(parent) > 0 {
			return false
		}
	}

	return true
}

func (heap *Heap) MaxHeapify(index int) {
	heapLen := len(*heap)
	currentElt, err := heap.Get(index)
	if err != nil {
		return
	}

	largest := index
	var largestElt Comparable = currentElt

	left := heap.Left(index)
	if left < heapLen {
		leftElt, _ := heap.Get(left)
		if leftElt.Compare(currentElt) > 0 {
			largest = left
			largestElt = leftElt
		}
	}

	right := heap.Right(index)
	if right < heapLen {
		rightElt, _ := heap.Get(right)
		if rightElt.Compare(largestElt) > 0 {
			largest = right
			largestElt = rightElt
		}
	}

	if largest != index {
		(*heap)[index], (*heap)[largest] = (*heap)[largest], (*heap)[index]
		heap.MaxHeapify(largest)
	}
}