package queue

import "testing"

func TestOverflowingQueue_Dequeue_WithoutItems(t *testing.T) {
	queue := NewOverflowingQueue(3)
	item := queue.Dequeue()
	if item != nil {
		t.Errorf("Expected to return nil, got %v", item)
	}
}

func TestOverflowingQueue_Enqueue_DoesntBlock(t *testing.T) {
	queue := NewOverflowingQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)
	queue.Enqueue(5)
}

func TestOverflowingQueue_Enqueue_DiscardsItems(t *testing.T) {
	queue := NewOverflowingQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Enqueue(3)
	queue.Enqueue(4)

	if val := queue.Dequeue(); val != 1 {
		t.Fatalf("Expected 1, got %v", val)
	}
	if val := queue.Dequeue(); val != 2 {
		t.Fatalf("Expected 2, got %v", val)
	}
	if val := queue.Dequeue(); val != 3 {
		t.Fatalf("Expected 3, got %v", val)
	}
	if val := queue.Dequeue(); val != nil {
		t.Fatalf("Expected nil, got %v", val)
	}
}

func TestOverflowingQueue_Enqueue_WithZeroSize(t *testing.T) {
	queue := NewOverflowingQueue(0)
	queue.Enqueue(1)
	queue.Enqueue(2)
	item := queue.Dequeue()
	if item != nil {
		t.Fatalf("Expected nil, got %v", item)
	}
}

func TestOverflowingQueue_EnqueueDequeueCombination(t *testing.T) {
	queue := NewOverflowingQueue(3)
	queue.Enqueue(1)
	if val := queue.Dequeue(); val != 1 {
		t.Fatalf("Expected 1, got %v", val)
	}
	queue.Enqueue(2)
	if val := queue.Dequeue(); val != 2 {
		t.Fatalf("Expected 2, got %v", val)
	}
}

func TestOverflowingQueue_DequeueChan(t *testing.T) {
	queue := NewOverflowingQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Close()

	dequeueChan := queue.DequeueChan()
	item1 := <-dequeueChan
	item2 := <-dequeueChan

	if item1 != 1 {
		t.Fatalf("Expected 1, got %v", item1)
	}

	if item2 != 2 {
		t.Fatalf("Expected 2, got %v", item2)
	}
}

func TestOverflowingQueue_DequeueAndDequeueChan(t *testing.T) {
	queue := NewOverflowingQueue(3)
	queue.Enqueue(1)
	queue.Enqueue(2)
	queue.Close()

	dequeueChan := queue.DequeueChan()
	item1 := <-dequeueChan
	item2 := queue.Dequeue()

	if item1 != 1 {
		t.Fatalf("Expected 1, got %v", item1)
	}

	if item2 != 2 {
		t.Fatalf("Expected 2, got %v", item2)
	}
}
