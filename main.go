package main

import (
	"context"
	"encoding/json"
	"log"
	"os"

	"github.com/segmentio/kafka-go"

	"envelope_db_writer/dao"
	"envelope_db_writer/entity"
)

func main() {
	dao.InitDB()

	// configure kafka consumer
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{os.Getenv("KAFKA_HOST")},
		GroupID:  os.Getenv("KAFKA_GROUPID"),
		Topic:    os.Getenv("KAFKA_TOPIC"),
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
	// kafka Reader connection will be automatically closed
	// when Docker container stops

	// subscribe for message
	for {
		msg, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Fatal("error occured when accepting message")
			break
		}
		envelope := entity.Envelope{}
		json.Unmarshal(msg.Value, &envelope)
		// decode operation type
		if envelope.Opened {
			dao.UpdateOpenState(&envelope)
		} else {
			dao.InsertEnvelope(&envelope)
		}
	}

	if err := r.Close(); err != nil {
		log.Fatal("failed to close reader:", err)
	}
}
