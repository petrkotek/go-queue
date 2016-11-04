package queue

// Queue is a standard queue data structure, which allows operations Enqueue and Dequeue.
type Queue interface {
	// Enqueue
	Enqueue(interface{})
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
