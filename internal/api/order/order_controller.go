package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/order"
	"picnshop/pkg/pagination"
	"picnshop/pkg/response"
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
		response.HandleError(g, err)
		return
	}

	err := c.orderService.CompleteOrder(req.UserId)
	if err != nil {
		response.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, response.Response{
			Message: "Order Created",
		})
}

func (c *Controller) CancelOrder(g *gin.Context) {
	var req CancelOrderRequest
	if err := g.ShouldBind(&req); err != nil {
		response.HandleError(g, err)
		return
	}
	err := c.orderService.CancelOrder(req.UserId, req.OrderId)
	if err != nil {
		response.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, response.Response{
			Message: "Order Created",
		})
}

func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	//TODO: get from jwt
	var userId uint = 1
	page = c.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)

}
