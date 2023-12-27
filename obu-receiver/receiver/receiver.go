package receiver

import (
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

type receiver struct {
	conn *websocket.Conn
}

func NewReceiver() *receiver {
	return &receiver{}
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
			log.Println("Error reading WebSocket message:", err)
			return
		}

		log.Printf("received OBU data: %s\n", message)
	}
}
