package order

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/order"
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

	g.JSON(http.StatusCreated, response.Response{
		Message: "Order Created",
	})
}
