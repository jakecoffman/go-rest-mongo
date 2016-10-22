package models

import "gopkg.in/mgo.v2/bson"

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Cats     []Cat
}
