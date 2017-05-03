package model

import (
	"errors"
	"fmt"

	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

type Store struct {
	MDB *mgo.Database
}

// SaveNode save a node
func (m *Store) SaveNode(n *Node) error {
	if n.ID == "" {
		n.ID = bson.NewObjectId().Hex()
	}
	if n.Type == "" {
		return errors.New("invalid data, missing node type")
	}
	c := m.MDB.C("node")
	_, err := c.UpsertId(n.ID, n)
	for refName, refVal := range n.Refs {
		if err := m.SaveRelationship(n.ID, refVal, refName); err != nil {
			return err
		}
	}
	return err
}

// SaveRelationship save node relationship
func (m *Store) SaveRelationship(from, to, name string) error {
	c := m.MDB.C("connections")
	id := fmt.Sprintf("%s:%s:%s", from, to, name)
	u := map[string]string{
		"from": from,
		"to":   to,
		"name": name,
	}
	_, err := c.UpsertId(id, u)
	return err
}

// ReadNode reads one node
func (m *Store) ReadNode(id string) (*Node, error) {
	var result Node
	err := m.MDB.C("node").FindId(id).One(&result)
	return &result, err
}

// Query run mongo query
func (m *Store) Query(q interface{}) ([]*Node, error) {
	var result []*Node
	query := m.MDB.C("node").Find(q)
	err := query.All(&result)
	return result, err
}

// Discover obtains a sub-graph for app start ups. It returns all nodes the user may use with certain limits?
// The discover rules is hardcoded here.
// THIS SHOULD BE DONE ON CLIENT!!!
// func (m *Store) Discover(userID string) ([]Node, error) {
// 	// App data
// 	// users, posts, offers, packages, payments ....
// 	//
// 	// `{ ownerid: '1234', type: 'post' }`
// 	// find all visible posts of an user
// 	// 1. get user's friend list 	{_id: 'user-id', friends }	= expand connections
// 	// 2. fore each connections friend do
// 	//		2.1 get friend's posts {_id: 'friend-id', posts} = expand connections
// 	// 3. get user's posts = expend connections
// 	// 4. get user's offers = expand connections
// 	// 5. get user's packages = expand
// 	// 6. get user's payments = expand
// }
