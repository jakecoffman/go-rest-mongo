package models

import (
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
	"errors"
)

const (
	USER_COLLECTION = "users"
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
	db *mgo.Database
}

func NewUserRepository(db *mgo.Database) *UserRepository {
	return &UserRepository{
		db: db,
	}
}

func (u *UserRepository) List(query map[string]interface{}, limit int, sort ...string) (interface{}, error) {
	users := []User{}
	err := u.db.C(USER_COLLECTION).Find(query).Sort(sort...).Limit(limit).All(&users)
	return users, err
}

func (u *UserRepository) Get(id string) (interface{}, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("that's no id")
	}
	user := User{}
	err := u.db.C(USER_COLLECTION).FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}
