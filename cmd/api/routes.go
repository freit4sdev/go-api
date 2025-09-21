package main

import (
	"github.com/freit4sdev/go-api/cmd/api/controller"
	"github.com/freit4sdev/go-api/internal/repository"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")
	categoryRepo := repository.NewCategoryRepository()

	categoryRoutes.GET("/", func(c *gin.Context) {
		controller.ListCategory(c, categoryRepo)
	})
	categoryRoutes.POST("/", func(c *gin.Context) {
		controller.CreateCategory(c, categoryRepo)
	})
}
