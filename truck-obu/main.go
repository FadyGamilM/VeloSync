package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"time"

	obu "github.com/FadyGamilM/truckobu/OBU"
	"github.com/gorilla/websocket"
)

const obuReceiverServerURL = "ws://localhost:8081/ws"

func main() {
	log.Fatal(SendReads())
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func PrintCoordinateEverySecond() {
	numOfTrucks := 50
	obuReads := obu.GenerateObuData(numOfTrucks)
	counter := 0
	for counter < numOfTrucks {
		log.Printf("Truck.{%v} ➜ %v\n", counter, obuReads[counter])
		counter++
		time.Sleep(1 * time.Second)
	}
}

func SendReads() error {
	// define a connection
	conn, _, err := websocket.DefaultDialer.Dial(obuReceiverServerURL, nil)
	if err != nil {
		return fmt.Errorf("error establishing a websocket connection to receiver microservice : %v", err.Error())
	}
	defer conn.Close()

	// get reads ready (here we simulate the generation)
	// for {
	obuReads := obu.GenerateObuData(10)
	for _, data := range obuReads {
		jsonRead, err := json.Marshal(data)
		if err != nil {
			log.Printf("error marshling data to json : %v", err)
			continue // beacuse we need to avoid the data that caused the problem
		}

		// send the data to the receiver service
		err = conn.WriteMessage(websocket.TextMessage, jsonRead)
		if err != nil {
			log.Printf("error sending data to receiver microservice : %v", err)
			continue // beacuse we need to avoid the data that caused the problem
		}

		log.Printf("Truck.{%v} Sent ➜ %v\n", data.ID, data)
		time.Sleep(1 * time.Second)
	}
	// }
}
