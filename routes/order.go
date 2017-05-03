package routes

// NewOrderHandler create order handler
// func NewOrderHandler(orderController *controller.OrderController) http.HandlerFunc {
// 	return func(res http.ResponseWriter, req *http.Request) {
// 		tkn := req.Context().Value("user").(*jwt.Token)
// 		clms := tkn.Claims.(jwt.MapClaims)
// 		s := clms["id"].(string)
// 		myUserID := bson.ObjectIdHex(s)

// 		var respBody interface{}
// 		if req.Method == "POST" {
// 			decoder := json.NewDecoder(req.Body)
// 			var order model.Order
// 			order.SellerID = myUserID
// 			if err := decoder.Decode(&order); err != nil {
// 				log.Println("bad request", err)
// 				http.Error(res, err.Error(), http.StatusBadRequest)
// 				return
// 			}
// 			defer req.Body.Close()
// 			fmt.Printf("hey: %#v\n", order)

// 			defer req.Body.Close()
// 			if err := orderController.SaveOrder(&order); err != nil {
// 				log.Println("create order error", err)
// 				http.Error(res, err.Error(), http.StatusInternalServerError)
// 				return
// 			}
// 			respBody = order
// 		} else {
// 			orders, _ := orderController.GetMyOrders(myUserID)
// 			respBody = orders
// 		}

// 		js, err := json.Marshal(respBody)
// 		if err != nil {
// 			log.Println("marshal response error", err)
// 			http.Error(res, err.Error(), http.StatusInternalServerError)
// 			return
// 		}
// 		res.Header().Set("Content-Type", "application/json")
// 		res.Write([]byte(js))
// 	}
// }
