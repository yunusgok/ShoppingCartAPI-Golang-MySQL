package category

import (
	"net/http"
	"picnshop/internal/domain/category"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	categoryService *category.CategoryService
}

func NewCategoryController(service *category.CategoryService) *CategoryController {
	return &CategoryController{
		categoryService: service,
	}
}

// TODO: change comment
// CreateCategory godoc
// @Summary Creates a new city
// @Tags City
// @Accept  json
// @Produce  json
// @Param createRequest body CreateCityRequest true "City informations"
// @Success 200 {object} CityResponse
// @Failure 400 {object} map[string]string
// @Failure 401 {object} map[string]string
// @Failure 404 {object} map[string]string
// @Failure 500 {object} map[string]string
// @Security ApiKeyAuth
// @param Authorization header string true "Authorization"
// @Router /city [post]
func (c *CategoryController) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": "Check your request body.",
		})
		g.Abort()
		return
	}
	category := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(category)
	if err != nil {
		g.JSON(http.StatusBadRequest, gin.H{
			"error_message": err.Error(),
		})
		g.Abort()
		return
	}

	g.JSON(http.StatusCreated, CategoryResponse{
		Name: category.Name,
		Desc: category.Desc,
	})
}
