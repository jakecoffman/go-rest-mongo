package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"errors"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Cats     []Cat
}

type Repository interface {
	List(query map[string]interface{}, limit int, sort ...string) (interface{}, error)
	Get(id string) (interface{}, error)
	//Insert(interface{}) error
	//Update(id string, interface{}) error
	//Delete(id string)
}

type UserRepository struct {
	collection *mgo.Collection
}

func NewUserRepository(collection *mgo.Collection) *UserRepository {
	return &UserRepository{
		collection: collection,
	}
}

func (u *UserRepository) List(query map[string]interface{}, limit int, sort ...string) (interface{}, error) {
	users := []User{}
	err := u.collection.Find(query).Sort(sort...).Limit(limit).All(&users)
	return users, err
}

func (u *UserRepository) Get(id string) (interface{}, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("that's no id")
	}
	user := User{}
	err := u.collection.FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}
