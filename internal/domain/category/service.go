package category

import (
	"mime/multipart"
	"picnshop/pkg/csv_helper"
	"picnshop/pkg/pagination"
)

type CategoryService struct {
	r Repository
}

func NewCategoryService(r Repository) *CategoryService {
	r.Migration()
	r.InsertSampleData()
	return &CategoryService{
		r: r,
	}
}

func (c *CategoryService) Create(category *Category) error {
	existCity := c.r.GetByName(category.Name)
	if len(existCity) > 0 {
		return ErrCategoryExistWithName
	}

	err := c.r.Create(category)
	if err != nil {
		return err
	}

	return nil
}

func (c *CategoryService) BulkCreate(fileHeader *multipart.FileHeader) (int, error) {
	categories := make([]*Category, 0)
	bulkCategory, err := csv_helper.ReadCsv(fileHeader)
	if err != nil {
		return 0, err
	}
	for _, categoryVariables := range bulkCategory {
		categories = append(categories, NewCategory(categoryVariables[0], categoryVariables[1]))
	}
	c.r.BulkCreate(categories)
	return len(categories), nil
}

func (c *CategoryService) GetAll(page *pagination.Pages) *pagination.Pages {
	categories, count := c.r.GetAll(page.Page, page.PageSize)
	page.Items = categories
	page.TotalCount = count
	return page

}
