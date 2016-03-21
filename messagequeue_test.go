package coinbase

import (
	"fmt"
	"log"
	"testing"
)

func TestMessageQueue(t *testing.T) {
	mq := NewMessageQueue()

	for i := 0; i < 100; i++ {
		msg := &Message{Sequence: i}
		mq.Push(msg)
		fmt.Println("size of queue: ", mq.size)
	}

	for i := 0; ; i++ {
		node, err := mq.Pop()
		if err != nil {
			log.Fatal(err)
			break
		}

		fmt.Printf("popping #%d: %+v\n", i, node)
		fmt.Printf("head: %+p\n", mq.Head())
		fmt.Printf("tail: %+p\n", mq.Tail())

	}

	fmt.Println("Final size of queue: %d", mq.Size())
}
