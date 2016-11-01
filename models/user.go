package models

import (
	"encoding/json"
	"errors"

	"github.com/jakecoffman/go-rest-mongo/datastore"
	"github.com/jakecoffman/go-rest-mongo/framework"
	"gopkg.in/mgo.v2/bson"
)

// User implements Resource
type User struct {
	Id       bson.ObjectId `bson:"_id"`
	Name     string
	Username string
	Cats     []Cat
	DogIds   []bson.ObjectId `json:"-"`
}

func (u *User) IsValid() bool {
	return u.Name != "" && u.Username != ""
}

func (u *User) MarshalJSON() ([]byte, error) {
	dogs := []Dog{}
	if err := datastore.Dog().Find(bson.M{"_id": bson.M{"$in": u.DogIds}}).All(&dogs); err != nil {
		return nil, err
	}

	type UserAlias User // prevents circular reference
	return json.Marshal(&struct {
		*UserAlias
		Dogs []Dog
	}{
		UserAlias: (*UserAlias)(u),
		Dogs:      dogs,
	})
}

// UserRepository implements Repository
type UserRepository struct{}

func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

func (u *UserRepository) New() framework.Resource {
	return &User{}
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

func (u *UserRepository) Insert(data interface{}) (interface{}, error) {
	user := data.(*User)
	user.Id = bson.NewObjectId()

	err := datastore.User().Insert(user)
	return user, err
}

func (u *UserRepository) Update(id string, data interface{}) (interface{}, error) {
	if !bson.IsObjectIdHex(id) {
		return nil, errors.New("that's no id")
	}
	user := data.(*User)
	err := datastore.User().UpdateId(bson.ObjectIdHex(id), user)
	return user, err
}

func (u *UserRepository) Delete(id string) error {
	if !bson.IsObjectIdHex(id) {
		return errors.New("that's no id")
	}
	return datastore.User().RemoveId(bson.ObjectIdHex(id))
}
