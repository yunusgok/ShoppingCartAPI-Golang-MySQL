package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/product"
	"picnshop/pkg/pagination"
	"picnshop/pkg/response"
)

type Controller struct {
	productService product.Service
}

func NewProductController(service product.Service) *Controller {
	return &Controller{
		productService: service,
	}
}

func (c *Controller) GetProducts(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.productService.GetAll(page)
	g.JSON(http.StatusOK, page)

}

func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		response.HandleError(g, err)
		return
	}

	err := c.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		response.HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, CreateProductResponse{
		Message: "product created",
	})
}
