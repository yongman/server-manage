package db

import (
	"labix.org/v2/mgo"
)

type Mongo struct {
	Url      string
	Database string
	Session  *mgo.Session
}

func NewMongo(url string, database string) *Mongo {
	return &Mongo{url, database, nil}
}
func NewMongoDefault() *Mongo {
	return &Mongo{"st01-arch-agent00.st01:27017", "test", nil}
}

func (m *Mongo) NewSession() (err error) {
	if m.Session == nil {
		m.Session, err = mgo.Dial(m.Url)
	}
	return err
}

func (m *Mongo) GetDB() (database *mgo.Database, err error) {
	m.NewSession()
	database = m.Session.DB(m.Database)
	if err != nil {
		return nil, err
	}
	m.Session.SetMode(mgo.Monotonic, true)
	return database, nil
}

func (m *Mongo) Close() (err error) {
	m.Session.Close()
	return nil
}
