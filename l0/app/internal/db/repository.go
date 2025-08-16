package db

import (
	"wb_technoschool/internal/models"
)

// Получить все заказы для восстановления кеша при старте
func (r *Repository) GetAllOrders() ([]models.Order, error) {
	rows, err := r.db.Query(`SELECT order_uid FROM orders`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var orders []models.Order
	for rows.Next() {
		var id string
		if err := rows.Scan(&id); err != nil {
			return nil, err
		}
		order, err := r.GetOrder(id)
		if err == nil {
			orders = append(orders, order)
		}
	}
	return orders, nil
}

// Получить заказ по ID
func (r *Repository) GetOrder(orderUID string) (models.Order, error) {
	var order models.Order

	// orders
	err := r.db.QueryRow(`
		SELECT order_uid, track_number, entry, customer_id, delivery_service, date_created
		FROM orders WHERE order_uid = $1
	`, orderUID).Scan(
		&order.OrderUID, &order.TrackNumber, &order.Entry,
		&order.CustomerID, &order.DeliveryService, &order.DateCreated,
	)
	if err != nil {
		return order, err
	}

	// delivery
	_ = r.db.QueryRow(`
		SELECT name, phone, zip, city, address, region, email
		FROM delivery WHERE order_uid = $1
	`, orderUID).Scan(
		&order.Delivery.Name, &order.Delivery.Phone, &order.Delivery.Zip,
		&order.Delivery.City, &order.Delivery.Address,
		&order.Delivery.Region, &order.Delivery.Email,
	)

	// ОБЯЗАТЕЛЬНО: чтобы фронт видел "order.delivery.service"
	order.Delivery.Service = order.DeliveryService

	// payment
	_ = r.db.QueryRow(`
		SELECT transaction, currency, provider, amount, bank, delivery_cost, goods_total
		FROM payment WHERE order_uid = $1
	`, orderUID).Scan(
		&order.Payment.Transaction, &order.Payment.Currency,
		&order.Payment.Provider, &order.Payment.Amount,
		&order.Payment.Bank, &order.Payment.DeliveryCost,
		&order.Payment.GoodsTotal,
	)

	// items
	rows, err := r.db.Query(`
		SELECT chrt_id, track_number, price, name, sale, size, total_price, nm_id, brand, status
		FROM items WHERE order_uid = $1
	`, orderUID)
	if err == nil {
		defer rows.Close()
		for rows.Next() {
			var item models.Item
			_ = rows.Scan(
				&item.ChrtID, &item.TrackNumber, &item.Price, &item.Name,
				&item.Sale, &item.Size, &item.TotalPrice,
				&item.NmID, &item.Brand, &item.Status,
			)
			item.Quantity = 1 // в схеме quantity нет — вернём 1
			order.Items = append(order.Items, item)
		}
	}

	return order, nil
}

// Вставить заказ (например, из Kafka)
func (r *Repository) InsertOrder(order models.Order) error {
	tx, err := r.db.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	// orders
	if _, err = tx.Exec(`
		INSERT INTO orders (order_uid, track_number, entry, customer_id, delivery_service, date_created)
		VALUES ($1,$2,$3,$4,$5,$6)
		ON CONFLICT (order_uid) DO NOTHING
	`, order.OrderUID, order.TrackNumber, order.Entry, order.CustomerID, order.DeliveryService, order.DateCreated); err != nil {
		return err
	}

	// delivery
	if _, err = tx.Exec(`
		INSERT INTO delivery (order_uid, name, phone, zip, city, address, region, email)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (order_uid) DO NOTHING
	`, order.OrderUID, order.Delivery.Name, order.Delivery.Phone,
		order.Delivery.Zip, order.Delivery.City, order.Delivery.Address,
		order.Delivery.Region, order.Delivery.Email); err != nil {
		return err
	}

	// payment
	if _, err = tx.Exec(`
		INSERT INTO payment (order_uid, transaction, currency, provider, amount, bank, delivery_cost, goods_total)
		VALUES ($1,$2,$3,$4,$5,$6,$7,$8)
		ON CONFLICT (order_uid) DO NOTHING
	`, order.OrderUID, order.Payment.Transaction, order.Payment.Currency,
		order.Payment.Provider, order.Payment.Amount, order.Payment.Bank,
		order.Payment.DeliveryCost, order.Payment.GoodsTotal); err != nil {
		return err
	}

	// items
	for _, item := range order.Items {
		if _, err = tx.Exec(`
			INSERT INTO items (chrt_id, order_uid, track_number, price, name, sale, size, total_price, nm_id, brand, status)
			VALUES ($1,$2,$3,$4,$5,$6,$7,$8,$9,$10,$11)
			ON CONFLICT (chrt_id) DO NOTHING
		`, item.ChrtID, order.OrderUID, item.TrackNumber,
			item.Price, item.Name, item.Sale, item.Size,
			item.TotalPrice, item.NmID, item.Brand, item.Status); err != nil {
			return err
		}
	}

	return tx.Commit()
}
