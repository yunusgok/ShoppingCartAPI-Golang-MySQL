package product

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/product"
	"picnshop/pkg/api_helper"
	"picnshop/pkg/pagination"
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
	queryText := g.Query("qt")
	if queryText != "" {
		page = c.productService.SearchProduct(queryText, page)
	} else {
		page = c.productService.GetAll(page)

	}
	g.JSON(http.StatusOK, page)

}

func (c *Controller) CreateProduct(g *gin.Context) {
	var req CreateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.CreateProduct(req.Name, req.Desc, req.Count, req.Price, req.CategoryID)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateProductResponse{
			Message: "product created",
		})
}

func (c *Controller) DeleteProduct(g *gin.Context) {
	var req DeleteProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.DeleteProduct(req.SKU)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateProductResponse{
			Message: "Product Deleted",
		})
}

func (c *Controller) UpdateProduct(g *gin.Context) {
	var req UpdateProductRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}

	err := c.productService.UpdateProduct(req.ToProduct())
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateProductResponse{
			Message: "Product Updated",
		})
}
