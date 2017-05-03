package model

import (
	"encoding/json"
	"testing"

	mgo "gopkg.in/mgo.v2"

	. "github.com/smartystreets/goconvey/convey"
)

func SetupStore() *mgo.Database {
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

func TestNode(t *testing.T) {
	js := `[{
    "type": "post",
    "id": "123",
    "title": "foobar.....",
    "imageUrls": ["http://abc.com/xyz"],
    "tags": ["ta", "tb"],
    "properties": {
        "price": "123.3",
        "brand": "Coach",
        "category": "包"
    },
    "refs": {}
}, {
    "type": "offer",
    "id": "231",
    "title": "",
    "imageUrls": [],
    "tags": [],
    "properties": {
        "name": "张三",
        "address": "北京市房山区。。。",
        "cellphone": "12421230123"
    },
    "refs": {
        "user": "af231s33sase",
        "order": "231x"
    }
}, {
    "type": "order",
    "id": "231x",
    "title": "",
    "imageUrls": [],
    "tags": [],
    "properties": {
        "name": "张三",
        "address": "北京市房山区。。。",
        "cellphone": "12421230123"
    },
    "refs": {
        "user": "af231s33sase",
        "package": "xc300001998hk"
    }
}, {
    "type": "package",
    "id": "xc300001998hk",
	"title": "",
    "imageUrls": [],
    "tags": [],
	"properties": {},
    "trackingID": "xc300001998hk",
	"refs": {}
}]`
	Convey("order", t, func() {
		var nodes []Node
		err := json.Unmarshal([]byte(js), &nodes)
		So(err, ShouldBeNil)
		So(len(nodes), ShouldEqual, 4)

		db := SetupStore()

		Convey("save", func() {
			d := Store{db}
			for _, node := range nodes {
				err := d.SaveNode(&node)
				So(err, ShouldBeNil)
				n, err := d.ReadNode(node.ID)
				So(err, ShouldBeNil)
				So(*n, ShouldResemble, node)
			}

			Convey("query", func() {
				Convey("all", func() {
					nodes, err := d.Query(nil)
					So(err, ShouldBeNil)
					So(nodes, ShouldNotBeNil)
				})

				Convey("by type", func() {
					q := map[string]interface{}{"type": "offer"}
					nodes, err := d.Query(q)
					So(err, ShouldBeNil)
					So(nodes, ShouldNotBeNil)
					So(nodes[0].ID, ShouldEqual, "231")
				})
			})
		})
	})
}
