package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/cart"
	"picnshop/pkg/pagination"
	"picnshop/pkg/response"
)

type Controller struct {
	cartService *cart.Service
}

func NewCartController(service *cart.Service) *Controller {
	return &Controller{
		cartService: service,
	}
}

func (c *Controller) AddItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		response.HandleError(g, err)
		return
	}

	err := c.cartService.AddItem(req.UserId, req.SKU, req.Count)
	if err != nil {
		response.HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, CreateCategoryResponse{
		Message: "created",
	})
}

func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		response.HandleError(g, err)
		return
	}

	err := c.cartService.UpdateItem(req.UserId, req.SKU, req.Count)
	if err != nil {
		response.HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, CreateCategoryResponse{
		Message: "updated",
	})
}

func (c *Controller) GetCart(g *gin.Context) {
	userId := pagination.ParseInt(g.Query("userId"), -1)

	result, err := c.cartService.GetCartItems(uint(userId))
	if err != nil {
		response.HandleError(g, err)
		return
	}
	g.JSON(200, result)
}
