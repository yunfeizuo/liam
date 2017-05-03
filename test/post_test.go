package test

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"io/ioutil"

	"encoding/json"

	. "github.com/smartystreets/goconvey/convey"
	"github.com/yunfeizuo/liam/model"
	"github.com/yunfeizuo/liam/routes"
	mgo "gopkg.in/mgo.v2"
)

func SetupDB() *mgo.Database {
	session, err := mgo.Dial("127.0.0.1")
	if err != nil {
		panic(err)
	}

	session.SetMode(mgo.Monotonic, true)

	// Collection order
	db := session.DB("liam-test")
	err = db.DropDatabase()
	if err != nil {
		panic(err)
	}

	db = session.DB("liam-test")
	return db
}

func GetTestToken(svrURL string) string {
	// url := fmt.Sprintf("%s/login?code=123", svrURL)
	// resp, err := http.Get(url)
	// if err != nil {
	// 	panic(err)
	// }

	// token, err := ioutil.ReadAll(resp.Body)
	// if err != nil {
	// 	panic(err)
	// }
	// return string(token)
	return `eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJpZCI6Im96NWtMMFlMYjREWmpfT0NsNm9wRGF0RXNEWkUiLCJvcGVuaWQiOiJvejVrTDBZTGI0RFpqX09DbDZvcERhdEVzRFpFIiwiZXhwIjoxNDk2Mjk0NzY1LCJpc3MiOiJsaWFtIn0.BzpRRhwhoKroohF16EIgDRMMaaQtQ-bWItwS-f7oGmQ`
}

func TestPost(t *testing.T) {
	Convey("Post", t, func() {
		db := SetupDB()
		defer db.Session.Close()

		router := routes.NewRouter(db)
		ts := httptest.NewServer(router)
		defer ts.Close()
		svrURL := ts.URL

		token := GetTestToken(svrURL)
		So(token, ShouldNotBeEmpty)

		Convey("create", func() {
			b, _ := os.Open("./samples/post.json")
			res, err := http.Post(fmt.Sprintf("%s/nodes?authorization=%s", svrURL, token), "application/json", b)
			So(err, ShouldBeNil)
			So(res.StatusCode, ShouldEqual, 200)
			res.Body.Close()

			Convey("query", func() {
				b, _ := os.Open("./samples/queryByType.json")
				res, err := http.Post(fmt.Sprintf("%s/query?authorization=%s", svrURL, token), "application/json", b)
				So(err, ShouldBeNil)
				So(res.StatusCode, ShouldEqual, 200)

				body, err := ioutil.ReadAll(res.Body)
				So(err, ShouldBeNil)
				var nodes []*model.Node
				err = json.Unmarshal(body, &nodes)
				So(err, ShouldBeNil)
				So(nodes[0].ID, ShouldEqual, "123")
				So(nodes[1].OwnerID, ShouldEqual, "oz5kL0YLb4DZj_OCl6opDatEsDZE")

				res.Body.Close()
			})
		})

	})
}
