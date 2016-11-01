package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/go-rest-mongo/datastore"
	"github.com/jakecoffman/go-rest-mongo/framework"
	"github.com/jakecoffman/go-rest-mongo/models"
	"gopkg.in/mgo.v2/bson"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// For bootstrapping
	datastore.DB().DropDatabase()
	for i := 0; i < 10; i++ {
		dogId := bson.NewObjectId()
		datastore.Dog().Insert(models.Dog{
			Id:   dogId,
			Name: "Rex",
		})
		datastore.User().Insert(models.User{
			Id:       bson.NewObjectId(),
			Name:     "Bootstrap",
			Username: "bstrap",
			Cats: []models.Cat{
				{Name: "Meowers"},
				{Name: "Ruffletuff"},
				{Name: "Zebra"},
				{Name: "Paws"},
				{Name: "Tiger"},
			},
			DogIds: []bson.ObjectId{dogId},
		})
	}

	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, bson.M{"hello": "world!"})
	})
	userGroup := r.Group("/users")
	{
		repo := models.NewUserRepository()
		userResource := framework.NewGenericController(repo)
		userGroup.GET("/", userResource.List)
		userGroup.GET("/:id", userResource.Get)
		userGroup.POST("/", userResource.Create)
		userGroup.PUT("/:id", userResource.Update)
		userGroup.DELETE("/:id", userResource.Delete)
	}
	log.Fatal(r.Run("0.0.0.0:9898"))
}
