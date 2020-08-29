package priorityqueue

import (
	"testing"
)

func assert(condition bool, t *testing.T, msg string, args ...interface{}) {
	if !condition {
		t.Errorf(msg, args...)
	}
}

func TestInititalization(t *testing.T) {
	queue := NewPriorityQueue()

	assert(queue.Len() == 0, t, "Empty priority queue length should be 0, not %d", queue.Len())
}

func TestPush(t *testing.T) {
	queue := NewPriorityQueue()
	queue.Push(&Element{Value: 1, Priority: 5})

	assert(queue.Len() == 1, t, "Priority queue length should be 1, not %d", queue.Len())
	assert(queue.At(0).Value == 1, t, "Element at index 0 should have value 1")

	queue.Push(&Element{Value: 0, Priority: 6})
	assert(queue.At(0).Value == 0, t, "Element at index 0 should have value 0")
	assert(queue.At(1).Value == 1, t, "Element at index 1 should have value 1")
}

func TestPop(t *testing.T) {
	queue := NewPriorityQueue()
	queue.Push(&Element{Value: 1, Priority: 10})
	queue.Push(&Element{Value: 2, Priority: 3})
	queue.Push(&Element{Value: 3, Priority: 15})
	queue.Push(&Element{Value: 4, Priority: -12})

	assert(queue.Pop().Value == 3, t, "First element to be removed should have value 3")
	assert(queue.Pop().Value == 1, t, "Second element to be removed should have value 2")
	assert(queue.Pop().Value == 2, t, "Third element to be removed should have value 1")
	assert(queue.Pop().Value == 4, t, "Fourth element to be removed should have value 4")
	assert(queue.Len() == 0, t, "Priority queue length should be 0, not %d", queue.Len())
}

func TestRemove(t *testing.T) {
	queue := NewPriorityQueue()
	queue.Push(&Element{Value: 1, Priority: 15})
	queue.Push(&Element{Value: 2, Priority: -12})
	queue.Push(&Element{Value: 3, Priority: 10})
	queue.Push(&Element{Value: 4, Priority: 14})
	queue.Push(&Element{Value: 5, Priority: 4})

	// The heap will keep the elements in the following order (by priority) [15, 14, 10, -12, 4]

	assert(queue.Remove(1).Value == 4, t, "Removed element at index 1 should have value 4")
	assert(queue.Remove(1).Value == 5, t, "Removed element at index 1 should have value 5")
	assert(queue.Remove(2).Value == 3, t, "Removed element at index 2 should have value 3")
	assert(queue.Remove(1).Value == 2, t, "Removed element at index 1 should have value 2")
	assert(queue.Remove(0).Value == 1, t, "Removed element at index 0 should have value 1")
	assert(queue.Len() == 0, t, "Priority queue length should be 0, not %d", queue.Len())
}
