package api

import (
	"encoding/json"
	"log"
	"net/http"

	"wb_technoschool/internal/cache"
	"wb_technoschool/internal/db"
)

type Handler struct {
	Cache *cache.Cache
	Repo  *db.Repository
}

func (h *Handler) GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	orderID := r.URL.Query().Get("order_uid")
	if orderID == "" {
		log.Println("GetOrder error: missing order_uid")
		http.Error(w, `{"error":"missing order_uid"}`, http.StatusBadRequest)
		return
	}

	// сначала кеш
	if order, ok := h.Cache.Get(orderID); ok {
		json.NewEncoder(w).Encode(order)
		log.Printf("GetOrder served from cache: order_uid=%s", orderID)
		return
	}

	// если нет в кеше — из БД
	order, err := h.Repo.GetOrder(orderID)
	if err != nil {
		log.Printf("GetOrder error: order not found, order_uid=%s", orderID)
		http.Error(w, `{"error":"order not found"}`, http.StatusNotFound)
		return
	}

	// положим в кеш и отдадим
	h.Cache.Set(order)
	json.NewEncoder(w).Encode(order)
	log.Printf("GetOrder served from DB and cached: order_uid=%s", orderID)
}
