package main

import (
	"github.com/freit4sdev/go-api/internal/use_case"
	"github.com/gin-gonic/gin"
)

func CategoryRoutes(r *gin.Engine) {
	categoryRoutes := r.Group("/categories")
	categoryRoutes.POST("/", use_case.NewCreateCategoryUseCase().Execute("test"))
}
