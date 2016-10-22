package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/jakecoffman/go-rest-mongo/models"
)

type GenericController struct {
	repository models.Repository
}

func NewGenericController(repository models.Repository) *GenericController {
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
