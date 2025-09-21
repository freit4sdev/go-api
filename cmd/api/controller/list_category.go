package controller

import (
	"net/http"

	"github.com/freit4sdev/go-api/internal/repository"
	"github.com/freit4sdev/go-api/internal/use_case"
	"github.com/gin-gonic/gin"
)


func ListCategory(c *gin.Context, categoryRepo repository.ICategoryRepository) {
	useCase := use_case.NewListCategoryUseCase(categoryRepo)
	categories, err := useCase.Execute()
	if err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError,
			gin.H{
				"success": false,
				"error":   err.Error(),
			})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"success": true,
		"categories": categories,
	})
}
