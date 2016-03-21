package coinbase

import "testing"

func TestWSRead(t *testing.T) {
	wsConn, err := NewWSConnection()
	if err != nil {
		t.Error("NewWSConnection failed.")
	}
	wsConn.Subscribe()
	wsConn.StartReadLoop()
	for true {
		t.Log("Current Message: %v", wsConn.CurrMessage)
	}

}
