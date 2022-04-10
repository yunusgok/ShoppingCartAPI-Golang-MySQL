package cart

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/cart"
	"picnshop/pkg/api_helper"
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
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.cartService.AddItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "created",
		})
}

func (c *Controller) UpdateItem(g *gin.Context) {
	var req ItemCartRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)

	err := c.cartService.UpdateItem(userId, req.SKU, req.Count)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Message: "updated",
		})
}

func (c *Controller) GetCart(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	result, err := c.cartService.GetCartItems(uint(userId))
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, result)
}
