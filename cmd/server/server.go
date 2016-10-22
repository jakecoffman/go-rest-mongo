package main

import (
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
)

func main() {
	r := gin.Default()
	r.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, bson.M{"hello": "world!"})
	})
	log.Fatal(r.Run("0.0.0.0:9898"))
}
