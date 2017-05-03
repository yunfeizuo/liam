package model

import jwt "github.com/dgrijalva/jwt-go"

type Ref struct {
	Name string
	ID   string
}

// Node graph node
type Node struct {
	ID         string                 `json:"id" bson:"_id"`
	OwnerID    string                 `json:"ownerId"`
	Type       string                 `json:"type"`
	Title      string                 `json:"title"`
	ImageUrls  []string               `json:"imageUrls"`
	Tags       []string               `json:"tags"`
	Properties map[string]interface{} `json:"properties,omitempty"`
	Refs       map[string]string      `json:"refs,omitempty"`
}

// UserInfo wx user info
type UserInfo struct {
	ID        string `json:"id,omitempty" bson:"_id,omitempty"`
	OpenID    string `json:"openid"`
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
	Name      string `json:"name"`
	Address   string `json:"address"`
	Cellphone string `json:"cellphone"`
}

// Claims token claims
type Claims struct {
	ID     string `json:"id"`
	OpenID string `json:"openid"`
	jwt.StandardClaims
}

// Content:
// product, order, payment,
// order - product,payment
// payment > order∂

type Product struct {
}

// type Order struct {
// 	ID        bson.ObjectId `json:"id,omitempty" bson:"_id,omitempty"`
// 	Title     string        `json:"title"`
// 	Note      string        `json:"note"`
// 	ImageURLs []string      `json:"imageUrls"`
// 	Brand     string        `json:"brand"`
// 	Category  string        `json:"category"`
// 	Price     float64       `json:"price"`
// 	// Buyer     UserInfo      `json:"buyer"`
// 	// SellerID  bson.ObjectId `json:"sellerid,omitempty" bson:"sellerid,omitempty"`
// // }

// type Package struct {
// 	ID      int64
// 	Orders  []Order
// 	Barcode string
// 	Weight  float32
// 	Contact
// 	ShipmentID int64
// }

// type Contact struct {
// 	Name      string
// 	Address   string
// 	Cellphone string
// }

// type Shipment struct {
// 	ID       int64
// 	Status   string
// 	Packages []Package
// 	ShipDate *time.Time
// }
