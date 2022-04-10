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

// AddItem godoc
// @Summary AddItem add product with given amount to cart of user
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [post]
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
			Message: "Item added to cart",
		})
}

// UpdateItem godoc
// @Summary UpdateItem in cart of the user
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Param ItemCartRequest body ItemCartRequest true "product information"
// @Success 200 {object} api_helper.Response
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [patch]
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

// GetCart godoc
// @Summary GetCart list ot items in user's cart
// @Tags Cart
// @Accept json
// @Produce json
// @Param        Authorization  header    string  true  "Authentication header"
// @Success 200 {array} cart.Item
// @Failure 400  {object} api_helper.ErrorResponse
// @Router /cart [get]
func (c *Controller) GetCart(g *gin.Context) {
	userId := api_helper.GetUserId(g)

	result, err := c.cartService.GetCartItems(userId)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.JSON(200, result)
}
