package heaps

import (
	"testing"
	"errors"
)

type Integer int

func (i Integer) Compare(j Comparable) int8 {
	jInt, ok := j.(Integer)
	if !ok {
		return 0
	}

	if jInt > i {
		return -1
	} else if jInt < i {
		return 1
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


func TestHeap_Right(t *testing.T) {
	testCases := []struct {
		Input, Expected int
	}{

		{2, 6},
		{0, 2},
	}

	for i, c := range testCases {
		heap := Heap{}

		if heap.Right(c.Input) != c.Expected {
			t.Errorf("TC %d: Expect parent of %d to be %d; got %d instead.",
				i, c.Input, c.Expected, heap.Parent(c.Input),
			)
		}
	}
}


func TestHeap_Get(t *testing.T) {
	testCases := []struct {
		Heap
		Index int
		Expected Integer
		ExpectedError error
	}{
		{
			Heap: Heap(
				[]Comparable{
					Integer(10),
					Integer(2),
					Integer(6),
					Integer(18),
				},
			),
			Index: 0,
			Expected: 10,
		},
		{
			Heap: Heap(
				[]Comparable{
					Integer(10),
					Integer(2),
					Integer(6),
					Integer(18),
				},
			),
			Index: 2,
			Expected: 6,
		},
		{
			Heap: Heap(
				[]Comparable{
					Integer(10),
					Integer(2),
					Integer(6),
					Integer(18),
				},
			),
			Index: -1,
			ExpectedError: errors.New("Index out of bounds"),
		},
		{
			Heap: Heap(
				[]Comparable{
					Integer(10),
					Integer(2),
					Integer(6),
					Integer(18),
				},
			),
			Index: 4,
			ExpectedError: errors.New("Index out of bounds"),
		},
		{
			Heap: Heap(
				[]Comparable{
				},
			),
			Index: 0,
			ExpectedError: errors.New("Index out of bounds"),
		},
	}

	for i, c := range testCases {
		comparable, err := c.Heap.Get(c.Index)
		if err != nil {
			if c.ExpectedError == nil {
				t.Errorf("TC %d: expected no error from calling Get(%d). Got %v instead.",
					i, c.Index, err,
				)
			}
			return
		}

		integer, ok := comparable.(Integer)
		if !ok {
			t.Errorf("TC %d: Expects to be an integer.", i)
			return
		}

		if integer != c.Expected {
			t.Errorf("TC %d: Expected %d from Get. Got %d instead",
				i, c.Expected, integer,
			)
		}
	}
}

func TestHeap_IsMaxHeap(t *testing.T) {
	testCases := []struct {
		Heap
		Expected bool
	}{
		{
			Heap([]Comparable{}),
			true,
		},
		{
			Heap([]Comparable{
				Integer(1),
			}),
			true,
		},
		{
			Heap([]Comparable{
				Integer(1),
				Integer(2),
			}),
			false,
		},
		{
			Heap([]Comparable{
				Integer(2),
				Integer(1),
				Integer(1),
			}),
			true,
		},
		{
			Heap([]Comparable{
				Integer(16),
				Integer(14),
				Integer(10),
				Integer(8),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(4),
				Integer(1),
			}),
			true,
		},
		{
			Heap([]Comparable{
				Integer(16),
				Integer(14),
				Integer(10),
				Integer(8),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(8),
				Integer(1),
			}),
			true,
		},
	}

	for i, c := range testCases {
		isMaxHeap := c.Heap.IsMaxHeap()
		if isMaxHeap != c.Expected {
			t.Errorf("TC %d: Expected to be max-heap: %t; got %t instead.",
				i, c.Expected, isMaxHeap,
			)
		}
	}
}

func areLikelyTheSame(a, b Heap) bool {
	if len(a) != len(b) {
		return false
	}

	for i, c := range a {
		if b[i].Compare(c) != 0 {
			return false
		}
	}

	return true
}

func TestHeap_MaxHeapify(t *testing.T) {
	testCases := []struct {
		HeapifyAt int
		Heap, Expected Heap
	}{
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{}),
			Expected: Heap([]Comparable{}),
		},
		{
			HeapifyAt: -1,
			Heap: Heap([]Comparable{}),
			Expected: Heap([]Comparable{}),
		},
		{
			HeapifyAt: 1,
			Heap: Heap([]Comparable{}),
			Expected: Heap([]Comparable{}),
		},
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(1),
			}),
		},
		{
			HeapifyAt: 1,
			Heap: Heap([]Comparable{
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(1),
			}),
		},
		{
			HeapifyAt: 2,
			Heap: Heap([]Comparable{
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(1),
			}),
		},
		{
			HeapifyAt: -1,
			Heap: Heap([]Comparable{
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(1),
			}),
		},
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{
				Integer(1),
				Integer(2),
			}),
			Expected: Heap([]Comparable{
				Integer(2),
				Integer(1),
			}),
		},
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{
				Integer(1),
				Integer(2),
				Integer(3),
			}),
			Expected: Heap([]Comparable{
				Integer(3),
				Integer(2),
				Integer(1),
			}),
		},
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{
				Integer(1),
				Integer(3),
				Integer(2),
			}),
			Expected: Heap([]Comparable{
				Integer(3),
				Integer(1),
				Integer(2),
			}),
		},
		{
			HeapifyAt: 0,
			Heap: Heap([]Comparable{
				Integer(16),
				Integer(4),
				Integer(10),
				Integer(14),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(8),
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(16),
				Integer(4),
				Integer(10),
				Integer(14),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(8),
				Integer(1),
			}),
		},
		{
			HeapifyAt: 1,
			Heap: Heap([]Comparable{
				Integer(16),
				Integer(4),
				Integer(10),
				Integer(14),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(8),
				Integer(1),
			}),
			Expected: Heap([]Comparable{
				Integer(16),
				Integer(14),
				Integer(10),
				Integer(8),
				Integer(7),
				Integer(9),
				Integer(3),
				Integer(2),
				Integer(4),
				Integer(1),
			}),
		},
	}

	for i, c := range testCases {
		c.Heap.MaxHeapify(c.HeapifyAt)

		if !areLikelyTheSame(c.Heap, c.Expected) {
			t.Errorf("TC %d: Expected %v. Got %v instead.",
				i, c.Expected, c.Heap,
			)
		}
	}
}