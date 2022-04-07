package cart

import (
	"errors"
	"picnshop/internal/domain/product"
)

type Service struct {
	cartRepository     Repository
	cartItemRepository ItemRepository
	productRepository  product.Repository
}

func NewService(cartRepository Repository, itemRepository ItemRepository, productRepository product.Repository) *Service {
	cartRepository.Migration()
	itemRepository.Migration()
	productRepository.Migration()
	return &Service{
		cartRepository:     cartRepository,
		cartItemRepository: itemRepository,
		productRepository:  productRepository,
	}

}

func (c *Service) AddItem(userID uint, sku string, count int) error {
	currentProduct, err := c.productRepository.FindBySKU(sku)
	if err != nil {
		return err
	}
	_, err = c.cartItemRepository.FindByID(currentProduct.ID)
	if err == nil {
		return errors.New("item already exists in cart")
	}
	if currentProduct.StockCount < count {
		return product.ErrProductStockIsNotEnough
	}
	currentCart, err := c.cartRepository.FindOrByUserID(userID)
	if err != nil {
		return err
	}
	err = c.cartItemRepository.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))
	if err != nil {
		return err
	}
	return nil
}

func (c *Service) GetCartItems(userId uint) ([]Item, error) {
	currentCart, err := c.cartRepository.FindOrByUserID(userId)
	if err != nil {
		return nil, err
	}
	items, err := c.cartItemRepository.GetItems(currentCart.ID)
	if err != nil {
		return nil, err
	}
	return items, nil

}
