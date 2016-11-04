package queue

// OverflowingQueue is a queue with given capacity. If capacity is reached, it starts discarding newly enqueued items.
// Underlying implementation is a buffered channel.
type OverflowingQueue struct {
	items chan interface{}
}

// NewOverflowingQueue returns new OverflowingQueue of specified capacity.
func NewOverflowingQueue(capacity int) *OverflowingQueue {
	return &OverflowingQueue{
		items: make(chan interface{}, capacity),
	}
}

// Enqueue adds an item to the queue.
func (q *OverflowingQueue) Enqueue(item interface{}) bool {
	select {
	case q.items <- item:
		return true
	default:
	// channel full, discarding item
	}
	return false
}

// Dequeue gets an item from the queue (and removes it from the queue). If no items are in the queue, returns nil.
func (q *OverflowingQueue) Dequeue() interface{} {
	select {
	case item := <-q.items:
		return item
	default:
	// no item available, just return nil
	}
	return nil
}

// DequeueChan returns a channel, which provides items from the queue.
func (q *OverflowingQueue) DequeueChan() <-chan interface{} {
	return q.items
}

func (q *OverflowingQueue) Close() {
	close(q.items)
}
