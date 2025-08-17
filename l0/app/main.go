package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"wb_technoschool/internal/api"
	"wb_technoschool/internal/cache"
	"wb_technoschool/internal/config"
	"wb_technoschool/internal/db"
	"wb_technoschool/internal/kafka"
)

// loggingMiddleware логирует каждый HTTP-запрос
func loggingMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		next.ServeHTTP(w, r)
		duration := time.Since(start)
		log.Printf("%s %s %s", r.Method, r.URL.Path, duration)
	})
}

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

	// Оборачиваем Handler в логирование
	http.Handle("/api/order", loggingMiddleware(http.HandlerFunc(h.GetOrder)))

	// Статика тоже через логирование
	http.Handle("/", loggingMiddleware(http.FileServer(http.Dir("./static"))))

	// Kafka
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
