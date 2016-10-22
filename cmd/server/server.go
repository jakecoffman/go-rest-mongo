package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
	"gopkg.in/mgo.v2"
	"github.com/jakecoffman/go-rest-mongo/controllers"
	"github.com/jakecoffman/go-rest-mongo/models"
)

const (
	USER_COLLECTION = "users"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	session, err := mgo.Dial("127.0.0.1:27017")
	if err != nil {
		log.Fatal(err)
	}
	defer session.Close()
	db := session.DB("test")

	// For bootstrapping
	db.DropDatabase()
	for i := 0; i < 1000; i++ {
		db.C(USER_COLLECTION).Insert(models.User{
			Id: bson.NewObjectId(),
			Name: "Bootstrap",
			Username: "bstrap",
			Cats: []models.Cat{
				{Name: "Meowers"},
				{Name: "Ruffletuff"},
				{Name: "Zebra"},
				{Name: "Paws"},
				{Name: "Tiger"},
			},
		})
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, bson.M{"hello": "world!"})
	})
	userGroup := r.Group("/users")
	{
		collection := db.C(USER_COLLECTION)
		repo := models.NewUserRepository(collection)
		userResource := controllers.NewGenericController(repo)
		userGroup.GET("/", userResource.List)
		userGroup.GET("/:id", userResource.Get)
	}
	log.Fatal(r.Run("0.0.0.0:9898"))
}
