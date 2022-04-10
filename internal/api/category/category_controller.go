package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/category"
	"picnshop/pkg/api_helper"
	"picnshop/pkg/pagination"
)

type Controller struct {
	categoryService *category.Service
}

func NewCategoryController(service *category.Service) *Controller {
	return &Controller{
		categoryService: service,
	}
}

//TODO: add swagger
func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		api_helper.HandleError(g, err)
		return
	}
	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}

	g.JSON(
		http.StatusCreated, CreateCategoryResponse{
			Name: newCategory.Name,
			Desc: newCategory.Desc,
		})
}
func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, _ := g.FormFile("file")
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		api_helper.HandleError(g, err)
		return
	}
	g.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! '%d' new categories created", fileHeader.Filename, count))
}

func (c *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)

}
