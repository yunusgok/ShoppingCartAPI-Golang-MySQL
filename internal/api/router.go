package api

import (
	cartApi "picnshop/internal/api/cart"
	categoryApi "picnshop/internal/api/category"
	productApi "picnshop/internal/api/product"
	userApi "picnshop/internal/api/user"

	"picnshop/internal/domain/cart"
	"picnshop/internal/domain/product"

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
	RegisterCartHandlers(db, r)
	RegisterProductHandlers(db, r)
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

func RegisterCartHandlers(db *gorm.DB, r *gin.Engine) {
	cartRepo := cart.NewCartRepository(db)
	cartItemRepo := cart.NewCartItemRepository(db)
	productRepo := product.NewProductRepository(db)
	cartService := cart.NewService(*cartRepo, *cartItemRepo, *productRepo)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart")
	cartGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.GET("/", cartController.GetCart)
}
func RegisterProductHandlers(db *gorm.DB, r *gin.Engine) {
	productRepo := product.NewProductRepository(db)
	productService := product.NewService(*productRepo)
	productController := productApi.NewProductController(*productService)
	productGroup := r.Group("/product")
	productGroup.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})
	productGroup.GET("/", productController.GetProducts)
	productGroup.POST("/product", productController.CreateProduct)

}
