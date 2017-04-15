package route

import (
	"database/sql"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunfeizuo/liam/controller"
)

func NewRouter(db *sql.DB) *mux.Router {
	shipmentController := controller.ShipmentController{db}
	router := mux.NewRouter()
	router.HandleFunc("/shipment/next", func(res http.ResponseWriter, req *http.Request) {
		shipmentController.NextShipment()
	}).Methods("GET")
	// Router.HandleFunc("/movie/{imdbKey}", handleMovie).Methods("GET", "DELETE", "POST")
	return router

}
