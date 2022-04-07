package product

import (
	"picnshop/pkg/pagination"
)

type Service struct {
	productRepository Repository
}

func NewService(cartRepository Repository) *Service {
	cartRepository.Migration()
	return &Service{
		productRepository: cartRepository,
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
