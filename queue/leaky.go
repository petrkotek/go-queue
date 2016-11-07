package queue

// LeakyQueue is a queue with given capacity. If capacity is reached, it starts discarding newly enqueued items.
// Underlying implementation is a buffered channel.
type LeakyQueue struct {
	items chan interface{}
}

// NewLeakyQueue returns new LeakyQueue of specified capacity.
func NewLeakyQueue(capacity int) *LeakyQueue {
	return &LeakyQueue{
		items: make(chan interface{}, capacity),
	}
}

// Enqueue adds an item to the queue.
func (q *LeakyQueue) Enqueue(item interface{}) bool {
	select {
	case q.items <- item:
		return true
	default:
	// channel full, discarding item
	}
	return false
}

// Dequeue gets an item from the queue (and removes it from the queue). If no items are in the queue, returns nil.
func (q *LeakyQueue) Dequeue() interface{} {
	select {
	case item := <-q.items:
		return item
	default:
	// no item available, just return nil
	}
	return nil
}

// DequeueChan returns a channel, which provides items from the queue.
func (q *LeakyQueue) DequeueChan() <-chan interface{} {
	return q.items
}

func (q *LeakyQueue) Close() {
	close(q.items)
}
