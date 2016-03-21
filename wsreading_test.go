package coinbase

import "testing"

func TestWSRead(t *testing.T) {
	wsConn, err := NewWSConnection()
	if err != nil {
		wsConn.Subscribe()
		wsConn.StartReadLoop()
	}

}
