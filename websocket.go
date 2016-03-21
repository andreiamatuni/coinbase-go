package coinbase

import (
	"fmt"

	"github.com/gorilla/websocket"
)

const cbWSAddress = "wss://ws-feed.exchange.coinbase.com"

var subscribeMessage = map[string]string{
	"type":       "subscribe",
	"product_id": "BTC-USD",
}

type WSConnection struct {
	Connection  *websocket.Conn
	MessageChan chan Message
	CurrMessage Message
}

func NewWSConnection() (WSConnection, error) {
	var wsDialer websocket.Dialer
	connection, _, err := wsDialer.Dial(cbWSAddress, nil)

	if err != nil {
		return WSConnection{}, err
	}

	wsConn := WSConnection{
		Connection:  connection,
		MessageChan: make(chan Message, 1024),
	}
	return wsConn, nil
}

func (wsConn *WSConnection) Subscribe() error {
	if err := wsConn.Connection.WriteJSON(subscribeMessage); err != nil {
		return err
	}
	return nil
}

func (wsConn *WSConnection) StartReadLoop() error {
	for true {
		if err := wsConn.Connection.ReadJSON(wsConn.CurrMessage); err != nil {
			return err
		}
		wsConn.MessageChan <- wsConn.CurrMessage
		fmt.Println(wsConn.CurrMessage)
	}
	return nil
}

/*func thing() {*/
//var wsDialer ws.Dialer
//wsConn, _, err := wsDialer.Dial(wsAddress, nil)
//if err != nil {
//fmt.Println(err.Error())
//}

//subscribe := map[string]string{
//"type":       "subscribe",
//"product_id": "BTC-USD",
//}

//if err := wsConn.WriteJSON(subscribe); err != nil {
//fmt.Println(err.Error())
//}

//message := Message{}

//for true {
//if err := wsConn.ReadJSON(&message); err != nil {
//fmt.Println(err.Error())
//break
//}
//fmt.Println(message)

//if message.Type == "match" {
//fmt.Println("Got a match")
//}
//}
/*}*/
