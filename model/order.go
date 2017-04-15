package model

import "time"

type OrderItem struct {
	ID          int64
	Name        string
	Description string
	Brand       string
}

type Package struct {
	ID      int64
	Items   []OrderItem
	Barcode string
	Weight  float32
	Contact
	ShipmentID int64
}

type Contact struct {
	Name      string
	Address   string
	Cellphone string
}

type Shipment struct {
	ID       int64
	Status   string
	Packages []Package
	ShipDate *time.Time
}
