package producer

import (
	"context"
	"encoding/json"
	"fmt"
	"log"

	"github.com/FadyGamilM/obureceiver/producer/events"
	"github.com/segmentio/kafka-go"
)

type Producer struct {
	topic      string
	brokerAddr string
}

func NewProducer(topic, brokerAddr string) *Producer {
	return &Producer{
		topic:      topic,
		brokerAddr: brokerAddr,
	}
}

func (p *Producer) SendMessage(ctx context.Context, msg []byte) error {

	// define kafka writer
	w := kafka.NewWriter(kafka.WriterConfig{
		Brokers:      []string{p.brokerAddr},
		Topic:        p.topic,
		RequiredAcks: int(kafka.RequireOne), // wait for the leader replica to ack and don't care about followers
	})

	// write the message to the specified topic
	// jsonRead, err := json.Marshal(msg)
	// if err != nil {
	// 	return fmt.Errorf("error trying to marshal msg to json format :%v", err)
	// }
	err := w.WriteMessages(ctx, kafka.Message{
		Value: msg,
	})
	if err != nil {
		return fmt.Errorf("error trying to send msg into kafka topic : %v", err)
	}
	publishedEvent := new(events.ObuData)
	err = json.Unmarshal(msg, publishedEvent)
	if err != nil {
		log.Printf("error to unmarshal published event : %v\n", err)
	}
	log.Printf("obu read with id = {%v} published to kafka successfully \n", publishedEvent.ID)

	return nil
}
