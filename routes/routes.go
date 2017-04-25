package route

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/yunfeizuo/liam/controller"
	"github.com/yunfeizuo/liam/model"
)

func NewRouter(db *sql.DB) *mux.Router {
	orderController := controller.OrderController{DB: db}
	shipmentController := controller.ShipmentController{DB: db}
	productController := controller.ProductController{DB: db}

	router := mux.NewRouter()
	router.HandleFunc("/orders", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			var order model.Order
			if err := decoder.Decode(&order); err != nil {
				log.Println("bad request", err)
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			defer req.Body.Close()

			defer req.Body.Close()
			if err := orderController.CreateOrder(&order); err != nil {
				log.Println("create order error", err)
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		orders, _ := orderController.GetMyOrders()
		js, err := json.Marshal(orders)
		if err != nil {
			log.Println("marshal response error", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}).Methods("GET", "POST")
	router.HandleFunc("/shipment/next", func(res http.ResponseWriter, req *http.Request) {
		shipment, _ := shipmentController.NextShipment()
		res.Write([]byte(fmt.Sprintf("%+v", shipment)))
	}).Methods("GET")

	router.HandleFunc("/products", func(res http.ResponseWriter, req *http.Request) {
		if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			var product model.Product
			if err := decoder.Decode(&product); err != nil {
				log.Println("bad request", err)
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			defer req.Body.Close()

			defer req.Body.Close()
			if err := productController.CreateProduct(&product); err != nil {
				log.Println("create order error", err)
				http.Error(res, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}

		products, _ := productController.GetProducts()
		js, err := json.Marshal(products)
		if err != nil {
			log.Println("marshal response error", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}

		res.Header().Set("Content-Type", "application/json")
		res.Write(js)
	}).Methods("GET", "POST")

	router.HandleFunc("/upload", upload).Methods("POST")
	router.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir("./upload"))))

	return router

}
