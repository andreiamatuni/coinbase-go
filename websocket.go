package coinbase

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const CoinbaseWSAddress = "wss://ws-feed.exchange.coinbase.com"

// subscription products (different currency markets)
var subscribeMessage = map[string]string{
	"type":       "subscribe",
	"product_id": "BTC-USD",
}

type WSConnection struct {
	Connection   *websocket.Conn
	Message      *Message
	MessageQueue *MessageQueue
}

func NewWSConnection() (WSConnection, error) {
	var wsDialer websocket.Dialer
	connection, _, err := wsDialer.Dial(CoinbaseWSAddress, nil)

	if err != nil {
		return WSConnection{}, err
	}

	wsConn := WSConnection{
		Connection: connection,
	}
	return wsConn, nil
}

func (wsConn *WSConnection) Subscribe(product string) error {
	switch product {
	case "BTC-USD":
		if err := wsConn.Connection.WriteJSON(subscribeMessage); err != nil {
			return err
		}
	case "BTC-EUR":
		subscribeMessage["product_id"] = product
		if err := wsConn.Connection.WriteJSON(subscribeMessage); err != nil {
			return err
		}
	case "BTC-GBP":
		subscribeMessage["product_id"] = product
		if err := wsConn.Connection.WriteJSON(subscribeMessage); err != nil {
			return err
		}
	case "BTC-CAD":
		subscribeMessage["product_id"] = product
		if err := wsConn.Connection.WriteJSON(subscribeMessage); err != nil {
			return err
		}
	}
	return nil
}

func (wsConn *WSConnection) ReadLoop(queue *MessageQueue, done chan bool) error {
	message := Message{}
	//wsConn.Message = message
	wsConn.MessageQueue = queue
	for {
		if err := wsConn.Connection.ReadJSON(&message); err != nil {
			fmt.Println("something went wrong")
			return err
		}
		wsConn.MessageQueue.Push(&message)
	}
	done <- true
	return nil
}
