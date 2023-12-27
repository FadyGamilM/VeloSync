package main

import (
	"log"
	"math/rand"
	"time"

	obu "github.com/FadyGamilM/truckobu/OBU"
)

func main() {
	PrintCoordinateEverySecond()
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
