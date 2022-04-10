package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/order"
	"picnshop/pkg/api_helper"
	"picnshop/pkg/pagination"
)

type Controller struct {
	orderService *order.Service
}

func NewOrderController(orderService *order.Service) *Controller {
	return &Controller{
		orderService: orderService,
	}
}

func (c *Controller) CompleteOrder(g *gin.Context) {
	var req CompleteOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)

	err := c.orderService.CompleteOrder(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Created",
		})
}

func (c *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	userId := api_helper.GetUserId(g)
	err := c.orderService.CancelOrder(userId, req.OrderId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, api_helper.Response{
			Message: "Order Created",
		})
}

func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)

}
