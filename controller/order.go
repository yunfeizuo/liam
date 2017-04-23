package controller

import (
	"database/sql"
	"log"

	"github.com/yunfeizuo/liam/model"
)

type OrderController struct {
	DB *sql.DB
}

func (c *OrderController) CreateOrder(order *model.Order) error {
	rows, err := c.DB.Query(`INSERT INTO orders (id, customer_id, title, note) VALUES ($1, $2, $3, $4)`,
		order.ID, order.CustomerID, order.Title, order.Note)
	if err != nil {
		log.Println("query my order error", err)
		return err
	}
	defer rows.Close()
	return nil
}

func (c *OrderController) GetMyOrders() ([]model.Order, error) {
	orders := make([]model.Order, 0, 128)

	rows, err := c.DB.Query(`SELECT id, customer_id, title, note FROM orders`)
	if err != nil {
		log.Println("query my order error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := model.Order{}
		if err = rows.Scan(&order.ID, &order.CustomerID, &order.Title, &order.Note); err != nil {
			log.Println("scan my order error", err)
			return nil, err
		}
		orders = append(orders, order)
	}

	return orders, nil
}
