package controller

import (
	"database/sql"
	"testing"

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
	_, err = db.Exec(`DELETE FROM package`)
	if err != nil {
		panic(err)
	}
	return db
}

func TestPackage(t *testing.T) {
	Convey("package", t, func() {
		db := SetupDB()
		controller := &PackageController{db}

		Convey("not found", func() {
			pkg, err := controller.GetPackage(0)
			So(err, ShouldBeNil)
			So(pkg, ShouldBeNil)
		})

		Convey("create", func() {
			Convey("without id", func() {
				pkg := &model.Package{ID: 0}
				err := controller.CreatePackage(pkg)
				So(err, ShouldBeNil)
				So(pkg.ID, ShouldBeGreaterThan, 0)
			})

			Convey("with an id", func() {
				pkg := &model.Package{ID: 12345}
				err := controller.CreatePackage(pkg)
				So(err, ShouldBeNil)
				So(pkg.ID, ShouldEqual, 12345)
			})
		})

		Convey("update", func() {
			pkg := &model.Package{}
			err := controller.CreatePackage(pkg)
			So(err, ShouldBeNil)
			So(pkg.ID, ShouldBeGreaterThan, 0)

			pkg.Weight = 5.6
			pkg.Address = "some fake address，你好"
			pkg.Barcode = "some barcode"
			pkg.Cellphone = "+1(243)232-2323"
			pkg.ShipmentID = 432
			pkg.Name = "who￥"
			err = controller.UpdatePackage(pkg)
			So(err, ShouldBeNil)
			newpkg, err := controller.GetPackage(pkg.ID)
			So(err, ShouldBeNil)
			So(newpkg, ShouldResemble, pkg)
		})

		Convey("ship", func() {
			pkg := &model.Package{}
			err := controller.CreatePackage(pkg)
			So(err, ShouldBeNil)
			So(pkg.ID, ShouldBeGreaterThan, 0)

			sc := ShipmentController{DB: db}
			shipment, err := sc.NextShipment()
			So(err, ShouldBeNil)

			err = controller.Ship(pkg)
			So(err, ShouldBeNil)
			So(pkg.ShipmentID, ShouldEqual, shipment.ID)
		})
	})
}
