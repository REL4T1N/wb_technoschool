package cache

import (
	"sync"
	"wb_technoschool/internal/models"
)

type Cache struct {
	mu     sync.RWMutex
	orders map[string]models.Order
}

func NewCache() *Cache {
	return &Cache{orders: make(map[string]models.Order)}
}

func (c *Cache) Set(order models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.orders[order.OrderUID] = order
}

func (c *Cache) Get(id string) (models.Order, bool) {
	c.mu.RLock()
	defer c.mu.RUnlock()
	order, ok := c.orders[id]
	return order, ok
}

func (c *Cache) LoadFromDB(orders []models.Order) {
	c.mu.Lock()
	defer c.mu.Unlock()
	for _, o := range orders {
		c.orders[o.OrderUID] = o
	}
}
