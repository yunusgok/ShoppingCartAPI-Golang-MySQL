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

// CompleteOrder godoc
// @Summary CompleteOrder with items in user's cart
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [post]
func (c *Controller) CompleteOrder(g *gin.Context) {
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

// CancelOrder godoc
// @Summary CancelOrder of user with given order id
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param CancelOrderRequest body CancelOrderRequest true "order information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /order [delete]
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
			Message: "Order Canceled",
		})
}

// GetOrders godoc
// @Summary GetOrders list of user's orders
// @Tags Order
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param page query int false "Page number"
// @Param pageSize query int false "Page size"
// @Success 200 {object} pagination.Pages
// @Router /order [get]
func (c *Controller) GetOrders(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	userId := api_helper.GetUserId(g)
	page = c.orderService.GetAll(page, userId)
	g.JSON(http.StatusOK, page)

}
