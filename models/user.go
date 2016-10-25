package models

import (
	"gopkg.in/mgo.v2/bson"
	"errors"
	"encoding/json"
	"github.com/jakecoffman/go-rest-mongo/datastore"
)

type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Cats     []Cat
	DogIds   []bson.ObjectId `json:"-"`
}

func (u *User) MarshalJSON() ([]byte, error) {
	dogs := []Dog{}
	if err := datastore.Dog().Find(bson.M{"_id": bson.M{"$in": u.DogIds}}).All(&dogs); err != nil {
		return nil, err
	}

	type Alias User
	return json.Marshal(&struct {
		*Alias
		Dogs []Dog
	}{
		Alias: (*Alias)(u),
		Dogs: dogs,
	})
}

type Repository interface {
	List(query map[string]interface{}, limit int, sort ...string) (interface{}, error)
	Get(id string) (interface{}, error)
	//Insert(interface{}) error
	//Update(id string, interface{}) error
	//Delete(id string)
}

type UserRepository struct {}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) List(query map[string]interface{}, limit int, sort ...string) (interface{}, error) {
	users := []User{}
	err := datastore.User().Find(query).Sort(sort...).Limit(limit).All(&users)
	return users, err
}

func (u *UserRepository) Get(id string) (interface{}, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("that's no id")
	}
	user := User{}
	err := datastore.User().FindId(bson.ObjectIdHex(id)).One(&user)
	return user, err
}
