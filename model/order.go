package model

import (
	"time"

	"github.com/jinzhu/gorm"
	_ "github.com/mattes/migrate/driver/postgres"
	uuid "github.com/satori/go.uuid"
)

type UserInfo struct {
	gorm.Model
	ID        uuid.UUID `json:"id"`
	Name      string    `json:"name"`
	Address   string    `json:"address"`
	Cellphone string    `json:"cellphone"`
}

type Order struct {
	gorm.Model
	ID     uuid.UUID   `json:"id"`
	Title  string      `json:"title"`
	Note   string      `json:"note"`
	Buyer  UserInfo    `json:"buyer",sql:"buyerid"`
	Seller UserInfo    `json:"seller"`
	Items  []OrderItem `json:"items"`
}

type OrderItem struct {
	gorm.Model
	ID       uuid.UUID `json:"id"`
	ImageURL string    `json:"imgUrl"`
	Title    string    `json:"title"`
	Note     string    `json:"note"`
	Brand    string    `json:"brand"`
	Category string    `json:"category"`
	Price    float64   `json:"price"`
}

type Package struct {
	gorm.Model
	ID      int64
	Items   []OrderItem
	Barcode string
	Weight  float32
	Contact
	ShipmentID int64
}

type Contact struct {
	gorm.Model
	Name      string
	Address   string
	Cellphone string
}

type Shipment struct {
	gorm.Model
	ID       int64
	Status   string
	Packages []Package
	ShipDate *time.Time
}
