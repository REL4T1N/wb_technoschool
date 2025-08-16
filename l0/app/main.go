package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"wb_technoschool/internal/api"
	"wb_technoschool/internal/cache"
	"wb_technoschool/internal/config"
	"wb_technoschool/internal/db"
	"wb_technoschool/internal/kafka"
)

func main() {
	cfg := config.LoadConfig()
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		cfg.PostgresHost, cfg.PostgresPort, cfg.PostgresUser, cfg.PostgresPassword, cfg.PostgresDB,
	)

	repo, err := db.NewRepository(connStr)
	if err != nil {
		log.Fatal("DB error: ", err)
	}

	c := cache.NewCache()

	// восстановим кеш
	if orders, err := repo.GetAllOrders(); err == nil {
		c.LoadFromDB(orders)
	}

	h := &api.Handler{Cache: c, Repo: repo}

	http.Handle("/api/order", http.HandlerFunc(h.GetOrder))
	// статические файлы
	http.Handle("/", http.FileServer(http.Dir("./static")))

	// Если нужна Kafka — раскомментируй:
	broker := getenv("KAFKA_BROKER", "localhost:9092")
	topic := getenv("KAFKA_TOPIC", "orders")
	go kafka.StartConsumer(broker, topic, c, repo)

	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getenv(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}
