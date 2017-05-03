package routes

import (
	"encoding/json"
	"log"
	"net/http"

	"fmt"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yunfeizuo/liam/model"
	"gopkg.in/mgo.v2"
)

// NewGraphHandler create graph API handler
func NewGraphHandler(mdb *mgo.Database) http.HandlerFunc {
	return func(res http.ResponseWriter, req *http.Request) {
		tkn := req.Context().Value("user").(*jwt.Token)
		clms := tkn.Claims.(jwt.MapClaims)
		myUserID := clms["id"].(string)
		store := model.Store{MDB: mdb}

		respBody := make([]*model.Node, 0)
		if req.Method == "POST" {
			decoder := json.NewDecoder(req.Body)
			var nodes []*model.Node
			if err := decoder.Decode(&nodes); err != nil {
				log.Println("bad request", err)
				http.Error(res, err.Error(), http.StatusBadRequest)
				return
			}
			defer req.Body.Close()

			for _, node := range nodes {
				node.OwnerID = myUserID
				if node.Refs == nil {
					node.Refs = make(map[string]string)
				}
				node.Refs["owner"] = myUserID
				if err := store.SaveNode(node); err != nil {
					log.Println("save node error", err)
					http.Error(res, err.Error(), http.StatusInternalServerError)
					return
				}
			}
			if nodes != nil {
				respBody = nodes
			}
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
