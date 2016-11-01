package datastore

import (
	"log"

	"gopkg.in/mgo.v2"
)

type DataStore struct {
	session *mgo.Session
}

const (
	database = "test"

	user = "user"
	dog  = "dog"
)

var dataStore *DataStore

func init() {
	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	dataStore = &DataStore{session: session}
}

func DB() *mgo.Database {
	return dataStore.session.DB(database)
}

func User() *mgo.Collection {
	return DB().C(user)
}

func Dog() *mgo.Collection {
	return DB().C(dog)
}
