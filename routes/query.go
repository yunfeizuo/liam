package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	"github.com/yunfeizuo/liam/model"
	"gopkg.in/mgo.v2"
)

// NewQueryHandler create Query API handler
func NewQueryHandler(mdb *mgo.Database) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		// TODO: security check
		//
		// tkn := req.Context().Value("user").(*jwt.Token)
		// clms := tkn.Claims.(jwt.MapClaims)
		// myUserID := clms["id"].(string)
		store := model.Store{MDB: mdb}

		var respBody interface{}
		if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			var query map[string]interface{}
			if err := decoder.Decode(&query); err != nil {
				log.Println("bad request", err)
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			defer req.Body.Close()

			nodes, err := store.Query(query)
			if err != nil {
				log.Println("query failed", err)
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			respBody = nodes
		} else {
			fmt.Println("nooooo")
		}

		js, err := json.Marshal(respBody)
		if err != nil {
			log.Println("marshal response error", err)
			http.Error(res, err.Error(), http.StatusInternalServerError)
			return
		}
		res.Header().Set("Content-Type", "application/json")
		res.Write([]byte(js))
	}
}
