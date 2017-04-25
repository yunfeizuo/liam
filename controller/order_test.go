package controller

import (
	"database/sql"
	"testing"

	uuid "github.com/satori/go.uuid"
	. "github.com/smartystreets/goconvey/convey"

	"github.com/yunfeizuo/liam/model"
	"github.com/yunfeizuo/liam/utils"
)

func SetupDB() *sql.DB {
	opts, err := utils.ParseConfigFile("../config/integration.json")
	if err != nil {
		panic(err)
	}
	db := utils.ConnectDB(opts)
	_, err = db.Exec(`DELETE FROM orders`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`DELETE FROM package`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`DELETE FROM shipment`)
	if err != nil {
		panic(err)
	}
	_, err = db.Exec(`DELETE FROM products`)
	if err != nil {
		panic(err)
	}
	return db
}

func TestOrder(t *testing.T) {
	Convey("order", t, func() {
		db := SetupDB()
		controller := &OrderController{db}

		Convey("not found", func() {
			orders, err := controller.GetMyOrders()
			So(err, ShouldBeNil)
			So(orders, ShouldBeEmpty)
		})

		Convey("create", func() {
			Convey("with an id", func() {
				buyer := model.UserInfo{ID: uuid.NewV4(), Name: "bill", Address: "some address", Cellphone: "1234567"}
				seller := model.UserInfo{ID: uuid.NewV4()}
				item := model.OrderItem{ID: uuid.NewV4(), Title: "item title", Note: "item note", ImageURL: "http://some.com/sxlas"}
				items := []model.OrderItem{item}
				order := &model.Order{ID: uuid.NewV4(), Title: "some title", Note: "bala", Buyer: buyer, Seller: seller, Items: items}
				err := controller.CreateOrder(order)
				So(err, ShouldBeNil)
				// So(pkg.ID, ShouldEqual, 12345)

				orders, err := controller.GetMyOrders()
				So(err, ShouldBeNil)
				o := orders[0]
				So(o.Buyer, ShouldResemble, buyer)
				So(o.Seller, ShouldResemble, seller)
				So(o.Title, ShouldResemble, order.Title)
				So(o.Note, ShouldResemble, order.Note)
				So(o.Items, ShouldResemble, order.Items)
				So(o.ID, ShouldResemble, order.ID)
			})
		})

		// Convey("update", func() {
		// 	pkg := &model.Package{}
		// 	err := controller.CreatePackage(pkg)
		// 	So(err, ShouldBeNil)
		// 	So(pkg.ID, ShouldBeGreaterThan, 0)

		// 	pkg.Weight = 5.6
		// 	pkg.Address = "some fake address，你好"
		// 	pkg.Barcode = "some barcode"
		// 	pkg.Cellphone = "+1(243)232-2323"
		// 	pkg.ShipmentID = 432
		// 	pkg.Name = "who￥"
		// 	err = controller.UpdatePackage(pkg)
		// 	So(err, ShouldBeNil)
		// 	newpkg, err := controller.GetPackage(pkg.ID)
		// 	So(err, ShouldBeNil)
		// 	So(newpkg, ShouldResemble, pkg)
		// })

		// Convey("ship", func() {
		// 	pkg := &model.Package{}
		// 	err := controller.CreatePackage(pkg)
		// 	So(err, ShouldBeNil)
		// 	So(pkg.ID, ShouldBeGreaterThan, 0)

		// 	sc := ShipmentController{DB: db}
		// 	shipment, err := sc.NextShipment()
		// 	So(err, ShouldBeNil)

		// 	err = controller.Ship(pkg)
		// 	So(err, ShouldBeNil)
		// 	So(pkg.ShipmentID, ShouldEqual, shipment.ID)
		// })
	})
}
