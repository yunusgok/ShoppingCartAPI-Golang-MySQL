package api

import (
	categoryApi "picnshop/internal/api/category"
	userApi "picnshop/internal/api/user"

	"picnshop/internal/domain/category"
	"picnshop/internal/domain/user"
	"picnshop/pkg/database_handler"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterHandlers(r *gin.Engine) {
	//TODO get settings from config
	db := database_handler.NewMySQLDB("go_test:password@tcp(127.0.0.1:3306)/go_database?parseTime=True&loc=Local")
	RegisterUserHandlers(db, r)
	RegisterCategoryHandlers(db, r)
	//TODO: delete ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

}

func RegisterCategoryHandlers(db *gorm.DB, r *gin.Engine) {
	categoryRepository := category.NewCategoryRepository(db)
	categoryService := category.NewCategoryService(*categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")
	categoryGroup.GET("/:name/:code", func(c *gin.Context) {
		name := c.Param("name")
		code := c.Param("code")
		c.JSON(200, gin.H{"name": name, "code": code})
	})
	categoryGroup.POST("", categoryController.CreateCategory)
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST("/upload", categoryController.BulkCreateCategory)
}

func RegisterUserHandlers(db *gorm.DB, r *gin.Engine) {
	userRepository := user.NewUserRepository(db)
	userService := user.NewUserService(*userRepository)
	userController := userApi.NewUserController(userService)
	userGroup := r.Group("/user")
	userGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/login", userController.Login)

}
