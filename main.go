package main

type queue struct {
	firstEl *queueNode
	lastEl  *queueNode
	size    int
}

type queueNode struct {
	nextEl *queueNode
	prevEl *queueNode
	value  int
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) first() interface{} {
	if receiver.firstEl == nil {
		return nil
	}
	return receiver.firstEl.value
}

func (receiver *queue) last() interface{} {
	if receiver.lastEl == nil {
		return nil
	}
	return receiver.lastEl.value
}

func (receiver *queue) equeue(value int) {
	if receiver.len() == 0 {
		receiver.firstEl = &queueNode{
			nextEl: nil,
			prevEl: nil,
			value:  value,
		}
		receiver.lastEl = receiver.firstEl
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.firstEl
	if current == nil {
		return
	}

	for {
		if current.nextEl == nil {
			current.nextEl = &queueNode{
				nextEl: nil,
				prevEl: current,
				value:  value,
			}
			receiver.lastEl = current.nextEl
			break
		}
		current = current.nextEl
	}
}

func (receiver *queue) dequeue() queue {
	if receiver.len() == 0 {
		return queue{}
	}

	firstToReturn := queue{
		firstEl: receiver.firstEl,
		lastEl:  nil,
		size:    0,
	}
	receiver.firstEl = receiver.firstEl.nextEl
	receiver.size--
	if receiver.size == 0 {
		receiver.lastEl = receiver.firstEl
	}
	return firstToReturn
}

func main() {}