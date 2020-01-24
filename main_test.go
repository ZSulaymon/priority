package main

import (
	"testing"
)

func Test_queue_length(t *testing.T) {
	q := queue{}
	q.equeue(0)
	if q.len() != 1 {
		t.Error("after adding one item to queue size must be 1, got: ", q.len())
	}
}

func Test_queue_first_and_last_item(t *testing.T) {
	q := queue{}
	q.equeue(1)
	if q.first() != q.last() {
		t.Error("after adding one item to queue first and last item must be same, got: ", q.first())
	}
}

func Test_queue_for_deleting_item(t *testing.T) {
	q := queue{}
	q.equeue(2)
	q.dequeue()
	if q.first() != nil {
		t.Error("deleting one item from queue it must be empty, got: ", q)
	}
}