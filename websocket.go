package coinbase

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const cbWSAddress = "wss://ws-feed.exchange.coinbase.com"

// subscription products (different currency markets)
var (
	subscribeBTCUSD = map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-USD",
	}

	subscribeBTCEUR = map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-EUR",
	}

	subscribeBTCGBP = map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-GBP",
	}

	subscribeBTCCAD = map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-CAD",
	}
)

type WSConnection struct {
	Connection *websocket.Conn
	Message    *Message
}

func NewWSConnection() (WSConnection, error) {
	var wsDialer websocket.Dialer
	connection, _, err := wsDialer.Dial(cbWSAddress, nil)

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
		if err := wsConn.Connection.WriteJSON(subscribeBTCUSD); err != nil {
			return err
		}
	case "BTC-EUR":
		if err := wsConn.Connection.WriteJSON(subscribeBTCEUR); err != nil {
			return err
		}

	case "BTC-GBP":
		if err := wsConn.Connection.WriteJSON(subscribeBTCGBP); err != nil {
			return err
		}
	case "BTC-CAD":
		if err := wsConn.Connection.WriteJSON(subscribeBTCCAD); err != nil {
			return err
		}

	}
	return nil
}

func (wsConn *WSConnection) ReadLoop(message *Message, done chan bool) error {
	wsConn.Message = message
	for {
		if err := wsConn.Connection.ReadJSON(message); err != nil {
			fmt.Println("something went wrong")
			return err
		}

		fmt.Println(message)

	}
	done <- true
	return nil
}
