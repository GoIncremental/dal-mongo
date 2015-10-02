package mongo

import (
	"github.com/goincremental/dal"
	"gopkg.in/mgo.v2"
)

type mongo struct {
	session *mgo.Session
	cloned  bool
}

func (m *mongo) isConnected() bool {
	return m.session != nil
}

// func (m *mongo) isClone() bool {
// 	return m.cloned
// }

//Connect takes a connection string and connects to the underlying mongo database
func (m *mongo) Connect(s string) (c dal.Connection, err error) {
	if s == "" {
		err = dal.ErrInvalidConnectionString
		return
	}
	if m.isConnected() {
		err = dal.ErrAlreadyConnected
		return
	}
	session, err := mgo.Dial(s)
	if err == nil {
		m.session = session
	}
	return m, err
}

func (m *mongo) Clone() (c dal.Connection, err error) {
	if !m.isConnected() {
		err = dal.ErrNotConnected
		return
	}
	session := m.session.Clone()
	c = &mongo{session: session}
	return c, nil
}

func (m *mongo) Close() error {
	if m.isConnected() {
		m.session.Close()
	}
	m.session = nil
	return nil
}

func (m *mongo) Tag(t string) (dal.Connection, error) {
	return m, nil
}

func (m *mongo) Create(dal.Item) (interface{}, error) {
	return m, nil
}

func (m *mongo) Read(dal.ID) (interface{}, error) {
	return m, nil
}

//NewConnection returns a dal Connection to a mongo datastore
func NewConnection() dal.Connection {
	m := &mongo{}
	return m
}
