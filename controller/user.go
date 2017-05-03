package controller

// // UserController user related stuff
// type UserController struct {
// 	Store *model.Store
// }

// // NewUserController create controller
// func NewUserController(store *model.Store) *UserController {
// 	return &UserController{Secret: secret, Store: store}
// }

// // User exchange code for session key
// func (c *UserController) UpdateUserInfo(userInfo *model.UserInfo) (bool, string, error) {
// 	// save user
// 	userNode := model.Node{
// 		ID:   openID,
// 		Type: "user",
// 		Properties: map[string]interface{}{
// 			"openid":     openID,
// 			"sessionkey": sessionKey,
// 		},
// 	}
// 	fmt.Printf("%+v", userNode)
// 	if err = c.Store.SaveNode(&userNode); err != nil {
// 		return false, "", err
// 	}

// 	// Create the Claims
// 	claims := model.Claims{
// 		ID:     openID,
// 		OpenID: openID,
// 		StandardClaims: jwt.StandardClaims{
// 			ExpiresAt: time.Now().Add(time.Hour * 24 * 30).Unix(),
// 			Issuer:    "liam",
// 		},
// 	}

// 	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
// 	tokenString, err := token.SignedString(c.Secret)

// 	return true, tokenString, err
// }

// func getWxOpenID(code string) (openid, sessionkey string, err error) {
// 	appid := os.Getenv("APPID")
// 	appsecrete := os.Getenv("APPSECRET")
// 	url := fmt.Sprintf(`https://api.weixin.qq.com/sns/jscode2session?appid=%s&secret=%s&js_code=%s&grant_type=authorization_code`,
// 		appid, appsecrete, code)

// 	fmt.Println(url)
// 	var resp *http.Response
// 	resp, err = http.DefaultClient.Get(url)
// 	if err != nil {
// 		return
// 	}
// 	if resp.StatusCode >= 400 {
// 		return "", "", errors.New("failed to get wx session key" + err.Error())
// 	}
// 	defer resp.Body.Close()

// 	var bytes []byte
// 	bytes, err = ioutil.ReadAll(resp.Body)
// 	if err != nil {
// 		return
// 	}
// 	fmt.Println(string(bytes))

// 	var result wxResp
// 	if err = json.Unmarshal(bytes, &result); err != nil {
// 		return
// 	}
// 	if result.Errcode != 0 {
// 		err = fmt.Errorf("Wechat returns %s", string(bytes))
// 		return
// 	}

// 	openid = result.OpenID
// 	sessionkey = result.SessionKey

// 	return
// }

// type wxResp struct {
// 	OpenID     string
// 	SessionKey string `json:"session_key"`
// 	Errcode    int
// 	Errmsg     string
// }
