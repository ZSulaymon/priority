package main

import (
	"testing"
)

func Test_empty_queue(t *testing.T) {
	q := queue{}
	if q.len() != 0 {
		t.Error("empty queue length must be zero, got:", q.len())
	}
	if q.first() != nil {
		t.Error("first element of queue must have value of 0, got:", q.first())
	}
	if q.last() != nil {
		t.Error("last element of queue must have value of 0, got:", q.last())
	}
	if q.last() != q.first() {
		t.Error("first and last elements of one item queue must have same priorities, got:", q.last(), ";", q.first())
	}
}

func Test_queue_with_one_item(t *testing.T) {
	q := queue{}
	q.equeue(1)
	if q.len() != 1 {
		t.Error("after adding one item queue length must be 1, got:", q.len())
	}
	if q.first() != 1 {
		t.Error("first element of queue must have value of 1, got:", q.first())
	}
	if q.last() != 1 {
		t.Error("last element of queue must have value of 1, got:", q.last())
	}
	if q.last() != q.first() {
		t.Error("first and last elements of one item queue must have same values, got:", q.last(), ";", q.first())
	}
}

func Test_queue_with_multiple_item(t *testing.T) {
	q := queue{}
	q.equeue(1)
	q.equeue(2)
	q.equeue(3)
	if q.len() != 3 {
		t.Error("after adding three item queue length must be 2, got:", q.len())
	}
	if q.first() != 1 {
		t.Error("first element of queue must have value of 1, got:", q.first())
	}
	if q.last() != 3 {
		t.Error("last element of queue must have value of 2, got:", q.last())
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