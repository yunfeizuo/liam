package controller

// import (
// 	"reflect"
// 	"testing"

// 	"github.com/yunfeizuo/liam/utils"
// )

// func TestGetOrCreateNextShipment(t *testing.T) {
// 	opts, err := utils.ParseConfigFile("../config/integration.json")
// 	if err != nil {
// 		panic(err)
// 	}
// 	db := utils.ConnectDB(opts)
// 	sc := ShipmentController{DB: db}
// 	_, err = db.Exec(`DELETE FROM shipment`)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// insert new row
// 	shipment1, err := sc.NextShipment()
// 	if err != nil {
// 		t.Fatal("unexpected error", err)
// 		return
// 	}
// 	if shipment1.Status != "new" {
// 		t.Fatal("should be new, actual", shipment1.Status)
// 	}

// 	// do it again
// 	shipment2, err := sc.NextShipment()
// 	if err != nil {
// 		t.Fatal("unexpected error", err)
// 		return
// 	}
// 	if !reflect.DeepEqual(shipment1, shipment2) {
// 		t.Fatal("should be the same", shipment1, shipment2)
// 	}

// 	var cnt int
// 	err = db.QueryRow(`SELECT count(*) FROM shipment`).Scan(&cnt)
// 	if err != nil {
// 		t.Fatal(err)
// 	}
// 	if cnt != 1 {
// 		t.Fatal("expect just one new row")
// 	}
// }
