package coinbase

import (
	"fmt"
	"testing"
	"time"
)

func TestWSReadLoop(t *testing.T) {
	wsConn, err := NewWSConnection()
	if err != nil {
		t.Error("NewWSConnection failed.")
	}

	wsConn.Subscribe("BTC-USD")

	queue := NewMessageQueue()
	var done chan bool

	go wsConn.ReadLoop(queue, done)

	for i := 0; i < 30; i++ {
		time.Sleep(time.Second * 2)
		fmt.Printf("Current queue size: %d\n", queue.Size())
		fmt.Printf("Last message recieved: %v\n\n", queue.Tail().Next().Value())
	}

	<-done
}
