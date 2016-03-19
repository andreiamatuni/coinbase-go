package main

import (
	"fmt"

	ws "github.com/gorilla/websocket"
)

const wsAddress = "wss://ws-feed.exchange.coinbase.com"

type Message struct {
	Type          string  `json:"type"`
	TradeId       int     `json:"trade_id,number"`
	OrderId       string  `json:"order_id"`
	Sequence      int     `json:"sequence,number"`
	MakerOrderId  string  `json:"maker_order_id"`
	TakerOrderId  string  `json:"taker_order_id"`
	Time          string  `json:"time"`
	RemainingSize float64 `json:"remaining_size,string"`
	NewSize       float64 `json:"new_size,string"`
	OldSize       float64 `json:"old_size,string"`
	Size          float64 `json:"size,string"`
	Price         float64 `json:"price,string"`
	Side          string  `json:"side"`
	Reason        string  `json:"reason"`
	OrderType     string  `json:"order_type"`
	Funds         float64 `json:"funds,string"`
	NewFunds      float64 `json:"new_funds,string"`
	OldFunds      float64 `json:"old_funds,string"`
}

func main() {
	var wsDialer ws.Dialer
	wsConn, _, err := wsDialer.Dial(wsAddress, nil)
	if err != nil {
		fmt.Println(err.Error())
	}

	subscribe := map[string]string{
		"type":       "subscribe",
		"product_id": "BTC-USD",
	}

	if err := wsConn.WriteJSON(subscribe); err != nil {
		fmt.Println(err.Error())
	}

	message := Message{}

	for true {
		if err := wsConn.ReadJSON(&message); err != nil {
			fmt.Println(err.Error())
			break
		}
		fmt.Println(message)

		if message.Type == "match" {
			fmt.Println("Got a match")
		}
	}

}
