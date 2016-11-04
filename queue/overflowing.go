package queue

// OverflowingQueue is a queue with given capacity. If capacity is reached, it starts discarding newly enqueued items.
type OverflowingQueue struct {
	items chan interface{}
}

func NewOverflowingQueue(capacity int) *OverflowingQueue {
	return &OverflowingQueue{
		items: make(chan interface{}, capacity),
	}
}

func (q *OverflowingQueue) Enqueue(item interface{}) bool {
	select {
	case q.items <- item:
		return true
	default:
	// channel full, discarding item
	}
	return false
}

func (q *OverflowingQueue) Dequeue() interface{} {
	select {
	case item := <-q.items:
		return item
	default:
	// no item available, just return nil
	}
	return nil
}

func (q *OverflowingQueue) DequeueChan() <-chan interface{} {
	return q.items
}
