package controller

import (
	"database/sql"
	"log"

	uuid "github.com/satori/go.uuid"
	"github.com/yunfeizuo/liam/model"
)

type OrderController struct {
	DB *sql.DB
}

func (c *OrderController) CreateOrder(order *model.Order) error {
	tx, err := c.DB.Begin()
	if err != nil {
		return err
	}
	defer tx.Rollback()

	_, err = tx.Exec(`
		INSERT INTO orders 
			(id, seller_id, buyer_id, buyer_name, buyer_address, buyer_cellphone, title, note) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
		order.ID, order.Seller.ID, order.Buyer.ID, order.Buyer.Name, order.Buyer.Address, order.Buyer.Cellphone, order.Title, order.Note)
	if err != nil {
		log.Println("insert order error", err)
		return err
	}

	for _, item := range order.Items {
		_, err := tx.Exec(`
		INSERT INTO orderitems 
			(id, order_id, image_url, title, note, brand, category, price) 
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)`,
			item.ID, order.ID, item.ImageURL, item.Title, item.Note, item.Brand, item.Category, item.Price)
		if err != nil {
			log.Println("inert order item error", err)
			return err
		}
	}

	err = tx.Commit()
	return nil
}

func (c *OrderController) UpdateBuyer(orderId uuid.UUID, buyer *model.Buyer) {
	_, err = c.DB.Exec(`UPDATE TABLE orders SET ()`)
}

func (c *OrderController) GetMyOrders() ([]model.Order, error) {
	orders := make([]model.Order, 0, 128)

	rows, err := c.DB.Query(`SELECT id, seller_id, buyer_id, buyer_name, buyer_address, buyer_cellphone, title, note FROM orders`)
	if err != nil {
		log.Println("query my order error", err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		order := model.Order{}
		if err = rows.Scan(&order.ID, &order.Seller.ID, &order.Buyer.ID, &order.Buyer.Name, &order.Buyer.Address,
			&order.Buyer.Cellphone, &order.Title, &order.Note); err != nil {
			log.Println("scan my order error", err)
			return nil, err
		}

		// find all items
		order.Items = make([]model.OrderItem, 0, 2)
		rows, err := c.DB.Query(`SELECT id, image_url, title, note, brand, category, price FROM orderitems WHERE order_id = $1`, order.ID)
		if err != nil {
			log.Println("query my order item error", err)
			return nil, err
		}
		defer rows.Close()
		for rows.Next() {
			item := model.OrderItem{}
			if err = rows.Scan(&item.ID, &item.ImageURL, &item.Title, &item.Note, &item.Brand, &item.Category, &item.Price); err != nil {
				log.Println("scan my order error", err)
				return nil, err
			}
			order.Items = append(order.Items, item)
		}

		orders = append(orders, order)
	}

	return orders, nil
}
