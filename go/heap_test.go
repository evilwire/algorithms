package heaps

import (
	"testing"
)

type Integer int

func (i Integer) Compare(j Comparable) int8 {
	jInt, ok := j.(Integer)
	if !ok {
		return 0
	}

	if jInt > i {
		return 1
	} else if jInt < i {
		return -1
	}

	return 0
}

func TestHeap_Length_Capacity(t *testing.T) {
	testCases := []struct{
		Heap
		ExpectedLength int
		ExpectedCapacity int
	}{
		{
			Heap(make([]Comparable, 2, 4)),
			2,
			4,
		},
		{
			Heap(make([]Comparable, 0, 0)),
			0,
			0,
		},
		{
			Heap(make([]Comparable, 0, 4)),
			0,
			4,
		},
		{
			Heap([]Comparable{}),
			0,
			0,
		},
		{
			Heap([]Comparable{
				Integer(1),
			}),
			1,
			1,
		},
	}

	for i, c := range testCases {
		if len(c.Heap) != c.ExpectedLength {
			t.Errorf("TC %d: expected length %d. Actually %d",
				i, c.ExpectedLength, len(c.Heap),
			)
		}

		if cap(c.Heap) != c.ExpectedCapacity {
			t.Errorf("TC %d: expected capacity %d. Actually %d",
				i, c.ExpectedCapacity, cap(c.Heap),
			)
		}
	}

}

func TestHeap_Parent(t *testing.T) {
	testCases := []struct {
		Input, Expected int
	}{

		{5, 2},
		{6, 2},
		{0, 0},
	}

	for i, c := range testCases {
		heap := Heap{}

		if heap.Parent(c.Input) != c.Expected {
			t.Errorf("TC %d: Expect parent of %d to be %d; got %d instead.",
				i, c.Input, c.Expected, heap.Parent(c.Input),
			)
		}
	}
}


func TestHeap_Left(t *testing.T) {
	testCases := []struct {
		Input, Expected int
	}{

		{2, 5},
		{0, 1},
	}

	for i, c := range testCases {
		heap := Heap{}

		if heap.Left(c.Input) != c.Expected {
			t.Errorf("TC %d: Expect parent of %d to be %d; got %d instead.",
				i, c.Input, c.Expected, heap.Parent(c.Input),
			)
		}
	}
}
