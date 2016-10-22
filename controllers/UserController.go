package controllers

import (
	"github.com/jakecoffman/go-rest-mongo/models"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_COLLECTION = "users"
)

type UserController struct {
	*GenericController
}

func NewUserController(db *mgo.Database) *UserController {
	collection := db.C(USER_COLLECTION)

	collection.Insert(models.User{
		Id: bson.NewObjectId(),
		Name: "Bootstrap",
		Username: "bstrap",
		Cats: []models.Cat{
			{Name: "Meowers"},
			{Name: "Ruffletuff"},
		},
	})

	repo := models.NewUserRepository(collection)
	return &UserController{
		GenericController: NewGenericController(repo),
	}
}
