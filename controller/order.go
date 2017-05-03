package controller

// import (
// 	"fmt"

// 	"github.com/yunfeizuo/liam/model"
// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"
// )

// type OrderController struct {
// 	Collection *mgo.Collection
// }

// func (c *OrderController) SaveOrder(order *model.Order) error {
// 	if !order.ID.Valid() {
// 		order.ID = bson.NewObjectId()
// 		fmt.Println("new order....")
// 	} else {
// 		fmt.Println("updating order...")
// 	}
// 	_, err := c.Collection.UpsertId(order.ID, order)
// 	return err
// }

// func (c *OrderController) GetMyOrders(sellerID bson.ObjectId) ([]model.Order, error) {
// 	var orders []model.Order
// 	err := c.Collection.Find(bson.M{"sellerid": sellerID}).All(&orders)
// 	return orders, err
// }
