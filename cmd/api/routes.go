package main

import (
	"github.com/freit4sdev/go-api/cmd/api/controller"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")
	categoryRoutes.POST("/", controller.CreateCategory)
}
