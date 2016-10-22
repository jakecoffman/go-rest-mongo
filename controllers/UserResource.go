package controllers

import (
	"github.com/jakecoffman/go-rest-mongo/models"
	"gopkg.in/mgo.v2"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
)

const (
	USER_COLLECTION = "users"
)

type UserResource struct {
	collection *mgo.Collection
}

func NewUserResource(db *mgo.Database) *UserResource {
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

	return &UserResource{collection: collection}
}

func (u *UserResource) List(ctx *gin.Context) {
	users := []models.User{}
	if err := u.collection.Find(nil).All(&users); err != nil {
		ctx.JSON(500, gin.H{"error": "failed listing users"})
		return
	}
	ctx.JSON(200, users)
}

func (u *UserResource) Get(ctx *gin.Context) {
	id := ctx.Param("id")
	if !bson.IsObjectIdHex(id) {
		ctx.JSON(422, gin.H{"error": "that's no id"})
		return
	}

	user := models.User{}
	if err := u.collection.FindId(bson.ObjectIdHex(id)).One(&user); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
		return
	}

	ctx.JSON(200, user)
}
