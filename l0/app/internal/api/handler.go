package api

import (
	"encoding/json"
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
		http.Error(w, `{"error":"missing order_uid"}`, http.StatusBadRequest)
		return
	}

	// сначала кеш
	if order, ok := h.Cache.Get(orderID); ok {
		json.NewEncoder(w).Encode(order)
		return
	}

	// если нет в кеше — из БД
	order, err := h.Repo.GetOrder(orderID)
	if err != nil {
		http.Error(w, `{"error":"order not found"}`, http.StatusNotFound)
		return
	}

	// положим в кеш и отдадим
	h.Cache.Set(order)
	json.NewEncoder(w).Encode(order)
}
