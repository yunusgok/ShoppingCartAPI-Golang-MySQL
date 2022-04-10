package product

import (
	"picnshop/pkg/pagination"
)

type Service struct {
	productRepository Repository
}

func NewService(productRepository Repository) *Service {
	productRepository.Migration()
	return &Service{
		productRepository: productRepository,
	}

}

func (c *Service) GetAll(page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.GetAll(page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page

}
func (c *Service) CreateProduct(name string, desc string, count int, price float32, cid uint) error {
	newProduct := NewProduct(name, desc, count, price, cid)
	err := c.productRepository.Create(newProduct)
	return err
}

func (c *Service) DeleteProduct(sku string) error {
	err := c.productRepository.Delete(sku)
	return err
}

func (c *Service) UpdateProduct(product *Product) error {
	err := c.productRepository.Update(*product)
	return err
}

// SearchProduct finds Products that matches their sku number or names with given str field
func (c *Service) SearchProduct(text string, page *pagination.Pages) *pagination.Pages {
	products, count := c.productRepository.SearchByString(text, page.Page, page.PageSize)
	page.Items = products
	page.TotalCount = count
	return page
}
