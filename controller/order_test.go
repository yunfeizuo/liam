package controller

// import (
// 	"testing"

// 	mgo "gopkg.in/mgo.v2"
// 	"gopkg.in/mgo.v2/bson"

// 	. "github.com/smartystreets/goconvey/convey"

// 	"github.com/yunfeizuo/liam/model"
// )

// // func SetupDB() *mgo.Collection {
// // 	session, err := mgo.Dial("127.0.0.1")
// // 	if err != nil {
// // 		panic(err)
// // 	}
// // 	// defer session.Close()

// // 	session.SetMode(mgo.Monotonic, true)

// // 	// Collection order
// // 	c := session.DB("liam").C("order_test")
// // 	c.RemoveAll(nil)
// // 	return c
// // }

// // func TestOrder(t *testing.T) {
// // 	Convey("order", t, func() {
// // 		db := SetupDB()
// // 		controller := &OrderController{db}

// // 		Convey("not found", func() {
// // 			orders, err := controller.GetMyOrders(bson.NewObjectId())
// // 			So(err, ShouldBeNil)
// // 			So(orders, ShouldBeEmpty)
// // 		})

// // 		buyer := model.UserInfo{UserID: bson.NewObjectId(), Name: "bill", Address: "some address", Cellphone: "1234567"}
// // 		sellerid := bson.NewObjectId()
// // 		order := &model.Order{Title: "some title", Note: "bala", ImageURLs: []string{"a", "b", "c"}, Buyer: buyer, SellerID: sellerid}

// // 		Convey("create", func() {
// // 			err := controller.SaveOrder(order)
// // 			So(err, ShouldBeNil)
// // 			So(order.ID.Valid(), ShouldBeTrue)

// // 			orders, err := controller.GetMyOrders(sellerid)
// // 			So(err, ShouldBeNil)
// // 			o := orders[0]
// // 			So(o, ShouldResemble, *order)
// // 		})

// // 		Convey("update", func() {
// // 			order.Buyer.Address = "new adress"
// // 			order.ImageURLs = append(order.ImageURLs, "http://abc.com/sxlas")
// // 			order.Brand = "coc"
// // 			err := controller.SaveOrder(order)
// // 			So(err, ShouldBeNil)
// // 			orders, err := controller.GetMyOrders(sellerid)
// // 			So(err, ShouldBeNil)
// // 			So(len(orders), ShouldEqual, 1)
// // 			o := orders[0]
// // 			So(o, ShouldResemble, *order)
// // 		})
// // 	})
// // }
