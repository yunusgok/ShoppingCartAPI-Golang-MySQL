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

func NewService(
	cartRepository Repository, itemRepository ItemRepository, productRepository product.Repository) *Service {
	cartRepository.Migration()
	itemRepository.Migration()
	return &Service{
		cartRepository:     cartRepository,
		cartItemRepository: itemRepository,
		productRepository:  productRepository,
	}

}

// AddItem adds the product with given amount to user's cart
func (c *Service) AddItem(userID uint, sku string, count int) error {
	currentProduct, err := c.productRepository.FindBySKU(sku)
	if err != nil {
		return err
	}
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userID)
	if err != nil {
		return err
	}
	_, err = c.cartItemRepository.FindByID(currentProduct.ID, currentCart.ID)
	if err == nil {
		return ErrItemAlreadyExistInCart
	}
	if currentProduct.StockCount < count {
		return product.ErrProductStockIsNotEnough
	}
	if count <= 0 {
		return ErrCountInvalid
	}
	err = c.cartItemRepository.Create(NewCartItem(currentProduct.ID, currentCart.ID, count))

	return err
}

// UpdateItem updates the amount of product inside user's cart
func (c *Service) UpdateItem(userID uint, sku string, count int) error {
	currentProduct, err := c.productRepository.FindBySKU(sku)
	if err != nil {
		return err
	}
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userID)
	if err != nil {
		return err
	}
	currentItem, err := c.cartItemRepository.FindByID(currentProduct.ID, currentCart.ID)
	if err != nil {
		return errors.New("item not exists in cart")
	}
	if currentProduct.StockCount+currentItem.Count < count {
		return product.ErrProductStockIsNotEnough
	}
	currentItem.Count = count
	err = c.cartItemRepository.Update(*currentItem)

	return err
}

// GetCartItems returns the items inside user's cart
func (c *Service) GetCartItems(userId uint) ([]Item, error) {
	currentCart, err := c.cartRepository.FindOrCreateByUserID(userId)
	if err != nil {
		return nil, err
	}
	items, err := c.cartItemRepository.GetItems(currentCart.ID)
	if err != nil {
		return nil, err
	}
	return items, nil

}
