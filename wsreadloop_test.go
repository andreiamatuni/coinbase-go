package coinbase

import "testing"

func TestWSReadLoop(t *testing.T) {
	wsConn, err := NewWSConnection()
	if err != nil {
		t.Error("NewWSConnection failed.")
	}
	wsConn.Subscribe("BTC-USD")
	var message Message
	var done chan bool

	go wsConn.ReadLoop(&message, done)

	<-done
}
