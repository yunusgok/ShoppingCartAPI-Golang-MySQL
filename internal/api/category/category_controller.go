package category

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"picnshop/internal/domain/category"
	"picnshop/pkg/pagination"
)

type Controller struct {
	categoryService *category.CategoryService
}

func NewCategoryController(service *category.CategoryService) *Controller {
	return &Controller{
		categoryService: service,
	}
}

//TODO: add swagger
func (c *Controller) CreateCategory(g *gin.Context) {
	var req CreateCategoryRequest
	if err := g.ShouldBind(&req); err != nil {
		HandleError(g, err)
		return
	}
	newCategory := category.NewCategory(req.Name, req.Desc)
	err := c.categoryService.Create(newCategory)
	if err != nil {
		HandleError(g, err)
		return
	}

	g.JSON(http.StatusCreated, CreateCategoryResponse{
		Name: newCategory.Name,
		Desc: newCategory.Desc,
	})
}
func (c *Controller) BulkCreateCategory(g *gin.Context) {
	fileHeader, _ := g.FormFile("file")
	count, err := c.categoryService.BulkCreate(fileHeader)
	if err != nil {
		HandleError(g, err)
		return
	}
	g.String(http.StatusOK, fmt.Sprintf("'%s' uploaded! '%d' new categories created", fileHeader.Filename, count))
}

func (c *Controller) GetCategories(g *gin.Context) {
	page := pagination.NewFromGinRequest(g, -1)
	page = c.categoryService.GetAll(page)
	g.JSON(http.StatusOK, page)

}

func HandleError(g *gin.Context, err error) {

	g.JSON(http.StatusBadRequest, gin.H{
		"error_message": err.Error(),
	})
	g.Abort()
	return

}
