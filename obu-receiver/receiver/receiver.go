package receiver

import (
	"context"
	"log"
	"net/http"

	"github.com/FadyGamilM/obureceiver/producer"
	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type receiver struct {
	conn          *websocket.Conn
	kafkaProducer *producer.Producer
}

func NewReceiver(p *producer.Producer) *receiver {
	return &receiver{
		kafkaProducer: p,
	}
}

func (rec *receiver) HandleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println("error upgrading to WebSocket : ", err)
		return
	}
	defer conn.Close()
	rec.conn = conn

	for {
		_, message, err := rec.conn.ReadMessage()
		if err != nil {
			log.Println("error reading WebSocket message:", err)
			return
		}

		log.Printf("received OBU data: %s\n", message)

		err = rec.kafkaProducer.SendMessage(context.Background(), message)
		if err != nil {
			log.Printf("error publishing reads to kafka : %v\n", err.Error())
		}
	}
}
