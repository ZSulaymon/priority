package main

type queue struct {
	firstEl *queueNode
	lastEl  *queueNode
	size    int
}

type queueNode struct {
	nextEl   *queueNode
	prevEl   *queueNode
	priority int
}

func (receiver *queue) len() int {
	return receiver.size
}

func (receiver *queue) first() *queueNode {
	return receiver.firstEl
}

func (receiver *queue) last() *queueNode {
	return receiver.lastEl
}

func (receiver *queue) equeue(priority int) {
	if receiver.len() == 0 {
		receiver.firstEl = &queueNode{
			nextEl:   nil,
			prevEl:   nil,
			priority: priority,
		}
		receiver.lastEl = receiver.firstEl
		receiver.size++
		return
	}
	receiver.size++
	current := receiver.firstEl
	if current != nil {
		for {
			if current.nextEl == nil {
				current.nextEl = &queueNode{
					nextEl:   nil,
					prevEl:   current,
					priority: priority,
				}
				receiver.lastEl = current.nextEl
				break
			}
			current = current.nextEl
		}
	}
}

func (receiver *queue) dequeue() queue {
	if receiver.len() == 0 {
		return queue{}
	}

	BeThefirstToReturn := queue{
		firstEl: receiver.firstEl,
		lastEl:  nil,
		size:    0,
	}
	receiver.firstEl = receiver.firstEl.nextEl
	receiver.size--
	if receiver.size == 0 {
		receiver.lastEl = receiver.firstEl
	}
	return BeThefirstToReturn
}

func main() {}