package controller

// import (
// 	"testing"

// 	uuid "github.com/satori/go.uuid"
// 	. "github.com/smartystreets/goconvey/convey"

// 	"github.com/yunfeizuo/liam/model"
// )

// func TestProduct(t *testing.T) {
// 	Convey("product", t, func() {
// 		db := SetupDB()
// 		controller := &ProductController{db}

// 		Convey("not found", func() {
// 			products, err := controller.GetProducts()
// 			So(err, ShouldBeNil)
// 			So(products, ShouldBeEmpty)
// 		})

// 		Convey("create", func() {
// 			product := model.Product{ID: uuid.NewV4(), Title: "some title", ImageURLs: []string{"foo", "bar"}}
// 			err := controller.CreateProduct(&product)
// 			So(err, ShouldBeNil)
// 		})

// 		Convey("create then find", func() {
// 			product := model.Product{ID: uuid.NewV4(), Title: "some title", ImageURLs: []string{"foo", "bar"}}
// 			err := controller.CreateProduct(&product)
// 			So(err, ShouldBeNil)

// 			products, err := controller.GetProducts()
// 			So(err, ShouldBeNil)
// 			expected := []model.Product{product}
// 			So(products, ShouldResemble, expected)
// 		})

// 		// Convey("update", func() {
// 		// 	pkg := &model.Package{}
// 		// 	err := controller.CreatePackage(pkg)
// 		// 	So(err, ShouldBeNil)
// 		// 	So(pkg.ID, ShouldBeGreaterThan, 0)

// 		// 	pkg.Weight = 5.6
// 		// 	pkg.Address = "some fake address，你好"
// 		// 	pkg.Barcode = "some barcode"
// 		// 	pkg.Cellphone = "+1(243)232-2323"
// 		// 	pkg.ShipmentID = 432
// 		// 	pkg.Name = "who￥"
// 		// 	err = controller.UpdatePackage(pkg)
// 		// 	So(err, ShouldBeNil)
// 		// 	newpkg, err := controller.GetPackage(pkg.ID)
// 		// 	So(err, ShouldBeNil)
// 		// 	So(newpkg, ShouldResemble, pkg)
// 		// })

// 		// Convey("ship", func() {
// 		// 	pkg := &model.Package{}
// 		// 	err := controller.CreatePackage(pkg)
// 		// 	So(err, ShouldBeNil)
// 		// 	So(pkg.ID, ShouldBeGreaterThan, 0)

// 		// 	sc := ShipmentController{DB: db}
// 		// 	shipment, err := sc.NextShipment()
// 		// 	So(err, ShouldBeNil)

// 		// 	err = controller.Ship(pkg)
// 		// 	So(err, ShouldBeNil)
// 		// 	So(pkg.ShipmentID, ShouldEqual, shipment.ID)
// 		// })
// 	})
// }
