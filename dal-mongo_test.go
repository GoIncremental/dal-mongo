package mongo

import (
	"testing"

	"github.com/goincremental/dal"
	. "github.com/smartystreets/goconvey/convey"
)

func getOpenConnectionTo(s string) (dal.Connection, error) {
	return getNewConnection().Connect(s)
}

func getNewConnection() dal.Connection {
	return NewConnection()
}

func TestConnect(t *testing.T) {
	Convey("Given a new connection", t, func() {
		c := getNewConnection()
		Convey("When I call connect with an empty string", func() {
			_, err := c.Connect("")
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("And the error should be an invalid Connection String", func() {
				So(err, ShouldEqual, dal.ErrInvalidConnectionString)
			})
		})
		Convey("When I call connect with localhost", func() {
			conn, err := c.Connect("localhost")
			Convey("Then I should not get an error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I should have a connection", func() {
				So(conn, ShouldNotBeNil)
			})
			Convey("When I call connect to localhost again", func() {
				_, err := conn.Connect("localhost")
				Convey("Then I should get an error", func() {
					So(err, ShouldNotBeNil)
				})
				Convey("And the error should be AlreadyConnected", func() {
					So(err, ShouldEqual, dal.ErrAlreadyConnected)
				})
			})
		})
	})
}

func TestClone(t *testing.T) {
	Convey("Given a new connection", t, func() {
		conn := NewConnection()
		Convey("When I create a clone", func() {
			_, err := conn.Clone()
			Convey("Then I should get an error", func() {
				So(err, ShouldNotBeNil)
			})
			Convey("And the error should be Not Connected", func() {
				So(err, ShouldEqual, dal.ErrNotConnected)
			})
		})
	})
	Convey("Given a connection to localhost", t, func() {
		conn, _ := getOpenConnectionTo("localhost")
		Convey("When I create a clone", func() {
			clone, err := conn.Clone()
			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I should have a connection", func() {
				So(clone, ShouldNotBeNil)
			})
		})
	})
}

func TestClose(t *testing.T) {
	Convey("Given a new connection", t, func() {
		conn := getNewConnection()
		Convey("When I close the connection", func() {
			err := conn.Close()
			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})
		})
	})
	Convey("Given an open connection to localhost", t, func() {
		conn, _ := getOpenConnectionTo("localhost")
		Convey("When I close the connection", func() {
			err := conn.Close()
			Convey("Then there should be no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I connect to localhost again", func() {
				_, err := conn.Connect("localhost")
				Convey("Then I should have no error", func() {
					So(err, ShouldBeNil)
				})
			})
		})
	})
}

func TestTag(t *testing.T) {
	Convey("Given a new connection to localhost", t, func() {
		conn, _ := getOpenConnectionTo("localhost")
		Convey("When I receive an empty tag", func() {
			conn, err := conn.Tag("")
			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I still have a connection", func() {
				So(conn, ShouldNotBeNil)
			})
		})
	})
}

func TestCreate(t *testing.T) {
	Convey("Given a new connection to localhost", t, func() {
		conn, _ := getOpenConnectionTo("localhost")
		Convey("When I receive an item to create", func() {
			item := &dummyItem{}
			conn, err := conn.Create(item)
			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I still have a connection", func() {
				So(conn, ShouldNotBeNil)
			})
		})
	})
}

func TestRead(t *testing.T) {
	Convey("Given a new connection to localhost", t, func() {
		conn, _ := getOpenConnectionTo("localhost")
		Convey("When I receive an readRequest", func() {
			conn, err := conn.Read(nil)
			Convey("Then there is no error", func() {
				So(err, ShouldBeNil)
			})
			Convey("And I still have a connection", func() {
				So(conn, ShouldNotBeNil)
			})
		})
	})
}
