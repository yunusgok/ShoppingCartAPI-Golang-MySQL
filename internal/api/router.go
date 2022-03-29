package api

import (
	categoryApi "picnshop/internal/api/category"
	"picnshop/internal/domain/category"
	"picnshop/pkg/database_handler"

	"github.com/gin-gonic/gin"
)

func RegisterHandlers(r *gin.Engine) {
	db := database_handler.NewMySQLDB("go_test:password@tcp(127.0.0.1:3306)/go_database?parseTime=True&loc=Local")
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(*categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	categoryGroup := r.Group("/category")
	categoryGroup.GET("/:name/:code", func(c *gin.Context) {
		name := c.Param("name")
		code := c.Param("code")
		c.JSON(200, gin.H{"name": name, "code": code})
	})
	categoryGroup.POST("", categoryController.CreateCategory)
}
