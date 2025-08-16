package kafka

import (
	"context"
	"encoding/json"
	"log"

	"wb_technoschool/internal/cache"
	"wb_technoschool/internal/db"
	"wb_technoschool/internal/models"

	"github.com/segmentio/kafka-go"
)

func StartConsumer(broker, topic string, c *cache.Cache, repo *db.Repository) {
	r := kafka.NewReader(kafka.ReaderConfig{
		Brokers: []string{broker},
		Topic:   topic,
		GroupID: "order-consumer",
	})

	for {
		m, err := r.ReadMessage(context.Background())
		if err != nil {
			log.Println("Kafka error:", err)
			continue
		}

		var order models.Order
		if err := json.Unmarshal(m.Value, &order); err != nil {
			log.Println("JSON error:", err)
			continue
		}

		if err := repo.InsertOrder(order); err != nil {
			log.Println("DB error:", err)
			continue
		}

		c.Set(order)
	}
}
