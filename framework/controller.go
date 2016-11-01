package framework

import (
	"log"

	"github.com/gin-gonic/gin"
)

type GenericController struct {
	repository Repository
	resource   Resource
}

func NewGenericController(repository Repository) *GenericController {
	return &GenericController{
		repository: repository,
	}
}

func (u *GenericController) List(ctx *gin.Context) {
	if resources, err := u.repository.List(nil, 0); err != nil {
		ctx.JSON(500, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, resources)
	}
}

func (u *GenericController) Get(ctx *gin.Context) {
	id := ctx.Param("id")

	if resource, err := u.repository.Get(id); err != nil {
		ctx.JSON(404, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, resource)
	}
}

func (u *GenericController) Create(ctx *gin.Context) {
	resource := u.repository.New()
	if err := ctx.BindJSON(resource); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !resource.IsValid() {
		log.Println("resource is not valid")
		ctx.JSON(422, gin.H{"error": "resource is not valid"})
		return
	}
	if resource, err := u.repository.Insert(resource); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, resource)
	}
}

func (u *GenericController) Update(ctx *gin.Context) {
	id := ctx.Param("id")
	resource := u.repository.New()
	if err := ctx.BindJSON(resource); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	if !resource.IsValid() {
		log.Println("resource is not valid")
		ctx.JSON(422, gin.H{"error": "resource is not valid"})
		return
	}
	if resource, err := u.repository.Update(id, resource); err != nil {
		log.Println(err)
		ctx.JSON(400, gin.H{"error": err.Error()})
	} else {
		ctx.JSON(200, resource)
	}
}

func (u *GenericController) Delete(ctx *gin.Context) {
	id := ctx.Param("id")
	if err := u.repository.Delete(id); err != nil {
		log.Println(err)
		ctx.JSON(404, gin.H{"error": "resource not found"})
	} else {
		ctx.JSON(200, gin.H{"message": "deleted"})
	}
}
