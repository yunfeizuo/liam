package main

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	"github.com/yunfeizuo/liam/routes"
)

func main() {

	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}
	defer session.Close()

	session.SetMode(mgo.Monotonic, true)
	db := session.DB("liam")

	// // Index
	// index := mgo.Index{
	// 	Key:        []string{"name", "phone"},
	// 	Unique:     true,
	// 	DropDups:   true,
	// 	Background: true,
	// 	Sparse:     true,
	// }

	// err = c.EnsureIndex(index)
	// if err != nil {
	// 	panic(err)
	// }

	router := routes.NewRouter(db)

	http.ListenAndServe(":8080", router)
}
