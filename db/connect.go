package db

import (
	"time"

	mgo "gopkg.in/mgo.v2"
)

type DBConnection struct {
	Session *mgo.Session
}

func (conn *DBConnection) NewConnection() {
	mongoDBDialInfo := &mgo.DialInfo{
		Addrs:    []string{"127.0.0.1:27017"},
		Timeout:  60 * time.Second,
		Database: "DBrsbundapamursih",
	}

	session, err := mgo.DialWithInfo(mongoDBDialInfo)
	if err != nil {
		panic(err)
	}
	session.SetMode(mgo.Monotonic, true)
	conn.Session = session
}

func (conn *DBConnection) Close() {
	conn.Session.Close()
	return
}

var DBconnect = DBConnection{}
