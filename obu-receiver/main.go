package main

import (
	"log"
	"net/http"

	"github.com/FadyGamilM/obureceiver/receiver"
)

func main() {
	receiver := receiver.NewReceiver()
	http.HandleFunc("/ws", receiver.HandleWS)
	log.Println("receiver microservice is up and running")
	if err := http.ListenAndServe(":8081", nil); err != nil {
		log.Fatalf("error starting the receiver server : %v", err)
	}
}
