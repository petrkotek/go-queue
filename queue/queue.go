package queue

// Queue is a standard queue data structure, which allows operations Enqueue and Dequeue.
type Queue interface {
	// Enqueue adds an item to a queue. If it was added successfully, returns true, otherwise false (e.g. in case
	// when the queue capacity is limited & was reached).
	Enqueue(interface{}) bool
	// Dequeue gets an item from the (FIFO) queue and removes it from the queue.
	Dequeue() interface{}
}

// ChanQueue is a Queue powered by a channel.
type ChanQueue interface {
	Queue
	// DequeueChan returns underlying channel of the queue, where you can read the items in the queue from.
	DequeueChan() <-chan interface{}
	// Close closes the underlying channel. Necessary for proper cleanup.
	Close()
}
