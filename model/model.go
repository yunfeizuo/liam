package model

import (
	"time"

	jwt "github.com/dgrijalva/jwt-go"

	"gopkg.in/mgo.v2/bson"
)

type Claims struct {
	UserInfo
	jwt.StandardClaims
}

type User struct {
	ID     bson.ObjectId `json:"id,omitempty"  bson:"_id,omitempty"`
	OpenID string
}

type UserInfo struct {
	UserID    bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	OpenID    string        `json:"openid"`
	NickName  string        `json:"nickName"`
	AvatarURL string        `json:"avatarUrl"`
	Name      string        `json:"name"`
	Address   string        `json:"address"`
	Cellphone string        `json:"cellphone"`
}

type Order struct {
	ID       bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
	Title    string        `json:"title"`
	Note     string        `json:"note"`
	Buyer    UserInfo      `json:"buyer"`
	SellerID bson.ObjectId `json:"sellerid,omitempty" bson:"sellerid,omitempty"`
	Items    []OrderItem   `json:"items"`
}

type OrderItem struct {
	ImageURL string  `json:"imgUrl"`
	Title    string  `json:"title"`
	Note     string  `json:"note"`
	Brand    string  `json:"brand"`
	Category string  `json:"category"`
	Price    float64 `json:"price"`
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
