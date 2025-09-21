package controller

import (
	"net/http"

	"github.com/freit4sdev/go-api/internal/use_case"
	"github.com/gin-gonic/gin"
)

type createCategoryInput struct {
	Name string `json:"name" binding:"required"`
}

func CreateCategory(c *gin.Context) {
	var body createCategoryInput
	if err := c.ShouldBindJSON(&body); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"success": false, "error": err.Error()})
		return
	}
	useCase := use_case.NewCreateCategoryUseCase()
	useCase.Execute(name string)
	categoryRoutes := r.Group("/categories")
	categoryRoutes.POST("/", use_case.NewCreateCategoryUseCase().Execute("test"))
}
