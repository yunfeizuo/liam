package controller

import (
	"database/sql"
	"log"

	"github.com/yunfeizuo/liam/model"
)

type ShipmentController struct {
	DB *sql.DB
}

func (c *ShipmentController) NextShipment() (*model.Shipment, error) {

	shipment := model.Shipment{}
	tx, err := c.DB.Begin()
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()

	err = tx.QueryRow(`SELECT id FROM shipment WHERE status='new' FOR UPDATE`).Scan(&shipment.ID)
	if err == sql.ErrNoRows {
		log.Println("no new shipment, creating one ...")
	} else if err != nil {
		return nil, err
	}

	if shipment.ID == 0 {
		_, err := tx.Query(`INSERT INTO shipment DEFAULT VALUES`)
		if err != nil {
			return nil, err
		}
	}

	err = tx.QueryRow(`SELECT id, ship_date, status FROM shipment WHERE status = 'new'`).Scan(&shipment.ID, &shipment.ShipDate, &shipment.Status)
	if err != nil {
		return nil, err
	}

	err = tx.Commit()
	return &shipment, err
}
