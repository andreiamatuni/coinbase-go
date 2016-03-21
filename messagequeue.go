package coinbase

import (
	"errors"
	"fmt"
)

/*
MessageQueue is a simple FIFO queue of Messages
*/
type MessageQueue struct {
	head *Node
	tail *Node
	size int
}

type Node struct {
	value *Message
	next  *Node
	prev  *Node
}

func NewMessageQueue() *MessageQueue {
	tail := Node{
		value: nil,
		prev:  nil,
	}

	head := Node{
		value: nil,
		next:  nil,
		prev:  &tail,
	}
	tail.next = &head

	return &MessageQueue{
		head: &head,
		tail: &tail,
		size: 0,
	}
}

func (q *MessageQueue) Push(message *Message) {
	node := Node{
		value: message,
		next:  q.tail.next,
		prev:  q.tail,
	}

	q.tail.next.prev = &node
	q.tail.next = &node
	q.size++
}

func (q *MessageQueue) Pop() (*Node, error) {
	node := q.head.prev
	fmt.Printf("%+v\n\n", node)
	if q.head.prev.prev == nil {
		return nil, errors.New("queue is empty")
	}
	q.head.prev = q.head.prev.prev
	q.size--
	return node, nil
}

func (q *MessageQueue) Size() int {
	return q.size
}

func (q *MessageQueue) Head() *Node {
	return q.head
}

func (q *MessageQueue) Tail() *Node {
	return q.tail
}

func (n *Node) Value() *Message {
	return n.value
}

func (n *Node) Next() *Node {
	return n.next
}

func (n *Node) Prev() *Node {
	return n.prev
}
