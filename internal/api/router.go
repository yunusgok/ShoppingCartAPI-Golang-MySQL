package api

import (
	cartApi "picnshop/internal/api/cart"
	categoryApi "picnshop/internal/api/category"
	orderApi "picnshop/internal/api/order"
	productApi "picnshop/internal/api/product"
	userApi "picnshop/internal/api/user"

	"picnshop/internal/domain/cart"
	"picnshop/internal/domain/order"
	"picnshop/internal/domain/product"

	"picnshop/internal/domain/category"
	"picnshop/internal/domain/user"
	"picnshop/pkg/database_handler"

	"github.com/gin-gonic/gin"
)

type Databases struct {
	categoryRepository    *category.Repository
	userRepository        *user.Repository
	productRepository     *product.Repository
	cartRepository        *cart.Repository
	cartItemRepository    *cart.ItemRepository
	orderRepository       *order.Repository
	orderedItemRepository *order.OrderedItemRepository
}

func CreateDBs() *Databases {
	//TODO get settings from config
	db := database_handler.NewMySQLDB("go_test:password@tcp(127.0.0.1:3306)/go_database?parseTime=True&loc=Local")
	return &Databases{
		categoryRepository:    category.NewCategoryRepository(db),
		cartRepository:        cart.NewCartRepository(db),
		userRepository:        user.NewUserRepository(db),
		productRepository:     product.NewProductRepository(db),
		cartItemRepository:    cart.NewCartItemRepository(db),
		orderRepository:       order.NewOrderRepository(db),
		orderedItemRepository: order.NewOrderedItemRepository(db),
	}
}
func RegisterHandlers(r *gin.Engine) {
	dbs := *CreateDBs()
	RegisterUserHandlers(r, dbs)
	RegisterCategoryHandlers(r, dbs)
	RegisterCartHandlers(r, dbs)
	RegisterProductHandlers(r, dbs)
	RegisterOrderHandlers(r, dbs)
}

func RegisterCategoryHandlers(r *gin.Engine, dbs Databases) {
	categoryService := category.NewCategoryService(*dbs.categoryRepository)
	categoryController := categoryApi.NewCategoryController(categoryService)
	categoryGroup := r.Group("/category")
	categoryGroup.POST("", categoryController.CreateCategory)
	categoryGroup.GET("", categoryController.GetCategories)
	categoryGroup.POST("/upload", categoryController.BulkCreateCategory)
}

func RegisterUserHandlers(r *gin.Engine, dbs Databases) {
	userService := user.NewUserService(*dbs.userRepository)
	userController := userApi.NewUserController(userService)
	userGroup := r.Group("/user")
	userGroup.POST("", userController.CreateUser)
	userGroup.POST("/login", userController.Login)

}

func RegisterCartHandlers(r *gin.Engine, dbs Databases) {
	cartService := cart.NewService(*dbs.cartRepository, *dbs.cartItemRepository, *dbs.productRepository)
	cartController := cartApi.NewCartController(cartService)
	cartGroup := r.Group("/cart")
	cartGroup.POST("/item", cartController.AddItem)
	cartGroup.PATCH("/item", cartController.UpdateItem)
	cartGroup.GET("/", cartController.GetCart)
}
func RegisterProductHandlers(r *gin.Engine, dbs Databases) {
	productService := product.NewService(*dbs.productRepository)
	productController := productApi.NewProductController(*productService)
	productGroup := r.Group("/product")
	productGroup.GET("", productController.GetProducts)
	productGroup.POST("", productController.CreateProduct)
	productGroup.DELETE("", productController.DeleteProduct)
	productGroup.PATCH("", productController.UpdateProduct)

}

func RegisterOrderHandlers(r *gin.Engine, dbs Databases) {
	orderService := order.NewService(*dbs.orderRepository, *dbs.orderedItemRepository, *dbs.productRepository, *dbs.cartRepository, *dbs.cartItemRepository)
	productController := orderApi.NewOrderController(orderService)
	productGroup := r.Group("/order")
	productGroup.POST("", productController.CompleteOrder)

}
