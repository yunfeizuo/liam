package controller

import (
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/yunfeizuo/liam/model"
	"gopkg.in/mgo.v2/bson"
)

// LoginController user related stuff
type LoginController struct {
	Secret []byte
}

// Login exchange code for session key
func (c *LoginController) Login(code string) (bool, string, error) {
	// get user info
	// getUserInfo()

	// Create the Claims
	claims := model.Claims{
		model.UserInfo{UserID: bson.NewObjectId()},
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
			Issuer:    "test",
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString(c.Secret)
	fmt.Println(c.Secret, tokenString, err)

	return true, tokenString, err
}
