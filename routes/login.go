package routes

import (
	"net/http"

	"github.com/yunfeizuo/liam/controller"
)

func NewLoginHandler(loginController *controller.LoginController) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		code := r.FormValue("code")
		if code == "" {
			http.Error(w, "missing code", http.StatusBadRequest)
		} else if ok, token, err := loginController.Login(code); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		} else if ok {
			w.Write([]byte(token))
		} else {
			http.Error(w, "登录失败："+err.Error(), http.StatusForbidden)
		}
	}
}
