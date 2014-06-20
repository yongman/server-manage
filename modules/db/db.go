package db

import (
	"fmt"
	"labix.org/v2/mgo"
)

type Mongo struct {
	Url        string
	Database   string
	Collection string
	Session    *mgo.Session
}

func NewMongo(url string, database string, collection string) *Mongo {
	return &Mongo{url, database, collection, nil}
}

func (m *Mongo) NewSession() (err error) {
	if m.Session == nil {
		m.Session, err = mgo.Dial(m.Url)
		fmt.Println("New Session")
	}
	return err
}

func (m *Mongo) GetDB() (database *mgo.Database, err error) {
	m.NewSession()
	database = m.Session.DB(m.Database)
	if err != nil {
		return nil, err
	}
	return database, nil
}

func (m *Mongo) Close() (err error) {
	m.Session.Close()
	return nil
}
