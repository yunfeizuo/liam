package routes

import (
	"net/http"

	mgo "gopkg.in/mgo.v2"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
	jwtmiddleware "github.com/yunfeizuo/go-jwt-middleware"
	"github.com/yunfeizuo/liam/controller"
)

var AuthSecret = "somepriveate;akfl;samykey"

func NewRouter(db *mgo.Database) http.Handler {
	router := mux.NewRouter()
	orderController := &controller.OrderController{Collection: db.C("order")}
	router.Handle("/orders", applyMiddlewaresWithAuth(NewOrderHandler(orderController))).Methods("GET", "POST")
	router.Handle("/upload", applyMiddlewaresWithAuth(upload)).Methods("POST")
	router.Handle("/login", applyMiddlewares(NewLoginHandler(&controller.LoginController{Secret: []byte(AuthSecret)})))
	router.PathPrefix("/download/").Handler(http.StripPrefix("/download/", http.FileServer(http.Dir("./upload"))))

	return router
}

func applyMiddlewares(handler http.HandlerFunc) http.Handler {
	n := negroni.Classic() // Includes some default middlewares
	n.UseHandler(handler)
	return n
}

func applyMiddlewaresWithAuth(handler http.HandlerFunc) http.Handler {
	ex := jwtmiddleware.FromFirst(jwtmiddleware.FromParameter("Authorization"), jwtmiddleware.FromParameter("authorization"), jwtmiddleware.FromAuthHeader)
	jwtMiddleware := jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(token *jwt.Token) (interface{}, error) {
			return []byte(AuthSecret), nil
		},
		SigningMethod: jwt.SigningMethodHS256,
		Extractor:     ex,
	})
	am := func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
		if err := jwtMiddleware.CheckJWT(w, r); err == nil {
			next(w, r)
		}
	}

	n := negroni.Classic() // Includes some default middlewares
	n.UseFunc(am)
	n.UseHandler(handler)
	return n
}
